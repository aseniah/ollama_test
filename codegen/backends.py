"""Inference backends for the codegen benchmark: Ollama, LM Studio, Apple."""

from __future__ import annotations

import json
import sys
import time
import urllib.error
import urllib.request
from collections.abc import Callable
from pathlib import Path
from typing import Any, Protocol, cast

sys.path.insert(0, str(Path(__file__).parent.parent))
import apfel_backend  # noqa: E402  (repo-root shared module)

import settings  # noqa: E402

PostFn = Callable[[str, dict[str, Any], int], dict[str, Any]]


class GenResult(dict[str, Any]):
    """Return shape of Backend.generate: response, ms, eval_count, tok_per_sec."""


class BackendError(Exception):
    """Raised when a backend cannot be reached or returns an unusable response."""


class Backend(Protocol):
    name: str

    def start(self) -> None: ...
    def stop(self) -> None: ...
    def warmup(self, messages: list[dict[str, str]]) -> None: ...
    def generate(
        self, messages: list[dict[str, str]], options: dict[str, Any], timeout: int, model: str
    ) -> GenResult: ...


def _result(response: str, ms: int, eval_count: int, tok_per_sec: float) -> GenResult:
    return GenResult(response=response, ms=ms, eval_count=eval_count, tok_per_sec=tok_per_sec)


def _http_post(url: str, payload: dict[str, Any], timeout: int) -> dict[str, Any]:
    data = json.dumps(payload).encode()
    req = urllib.request.Request(
        url, data=data, headers={"Content-Type": "application/json"}, method="POST"
    )
    try:
        with urllib.request.urlopen(req, timeout=timeout) as resp:
            return cast("dict[str, Any]", json.loads(resp.read()))
    except urllib.error.HTTPError as e:
        body = e.read().decode(errors="replace")
        raise BackendError(f"HTTP {e.code} from {url}: {body[:300]}") from e
    except urllib.error.URLError as e:
        raise BackendError(f"cannot reach {url}: {e.reason}") from e


def _http_get_ok(url: str, timeout: int = 3) -> bool:
    try:
        with urllib.request.urlopen(urllib.request.Request(url, method="GET"), timeout=timeout):
            return True
    except Exception:
        return False


class OllamaBackend:
    name = "ollama"

    def __init__(self, base_url: str, _post: PostFn = _http_post) -> None:
        self._base = base_url.rstrip("/")
        self._post = _post

    def start(self) -> None:
        if not _http_get_ok(f"{self._base}/api/tags"):
            raise BackendError(
                f"Ollama not reachable at {self._base} — is `ollama serve` running?"
            )

    def stop(self) -> None:
        return None

    def warmup(self, messages: list[dict[str, str]]) -> None:
        return None

    def generate(
        self, messages: list[dict[str, str]], options: dict[str, Any], timeout: int, model: str
    ) -> GenResult:
        payload: dict[str, Any] = {
            "model": model, "messages": messages, "stream": False, **options,
        }
        start = time.monotonic()
        data = self._post(f"{self._base}/api/chat", payload, timeout)
        ms = int((time.monotonic() - start) * 1000)
        eval_count = int(data.get("eval_count", 0))
        eval_dur = int(data.get("eval_duration", 0))
        tok_per_sec = round(eval_count / (eval_dur / 1e9), 1) if eval_dur else 0.0
        content = str(cast("dict[str, Any]", data["message"])["content"]).strip()
        return _result(content, ms, eval_count, tok_per_sec)


class LMStudioBackend:
    name = "lmstudio"

    def __init__(self, base_url: str, _post: PostFn = _http_post) -> None:
        self._base = base_url.rstrip("/")
        self._post = _post

    def start(self) -> None:
        if not _http_get_ok(f"{self._base}/v1/models"):
            raise BackendError(
                f"LM Studio not reachable at {self._base} — "
                "start the LM Studio server and load a model"
            )

    def stop(self) -> None:
        return None

    def warmup(self, messages: list[dict[str, str]]) -> None:
        return None

    def generate(
        self, messages: list[dict[str, str]], options: dict[str, Any], timeout: int, model: str
    ) -> GenResult:
        # `think` is an Ollama-only option; drop it. Pass anything else through.
        passthrough = {k: v for k, v in options.items() if k != "think"}
        payload: dict[str, Any] = {
            "model": model, "messages": messages, "stream": False, **passthrough,
        }
        start = time.monotonic()
        try:
            data = self._post(f"{self._base}/api/v0/chat/completions", payload, timeout)
            stats = cast("dict[str, Any]", data.get("stats") or {})
        except BackendError:
            data = self._post(f"{self._base}/v1/chat/completions", payload, timeout)
            stats = {}
        ms = int((time.monotonic() - start) * 1000)
        choices = cast("list[dict[str, Any]]", data["choices"])
        content = str(cast("dict[str, Any]", choices[0]["message"])["content"]).strip()
        usage = cast("dict[str, Any]", data.get("usage") or {})
        eval_count = int(
            stats.get("predicted_tokens_count") or usage.get("completion_tokens") or 0
        )
        tps = stats.get("tokens_per_second")
        if tps:
            tok_per_sec = round(float(tps), 1)
        else:
            tok_per_sec = round(eval_count / (ms / 1000), 1) if ms else 0.0
        return _result(content, ms, eval_count, tok_per_sec)


class AppleBackend:
    name = "apple"

    def __init__(self, autostart: bool = True) -> None:
        self._autostart = autostart
        self._proc: Any = None

    def start(self) -> None:
        self._proc = apfel_backend.ensure_running()

    def stop(self) -> None:
        if self._proc is not None:
            apfel_backend.teardown(self._proc)
            self._proc = None

    def warmup(self, messages: list[dict[str, str]]) -> None:
        apfel_backend.warmup(messages)

    def generate(
        self, messages: list[dict[str, str]], options: dict[str, Any], timeout: int, model: str
    ) -> GenResult:
        r: dict[str, Any] = apfel_backend.run_prompt(messages, "", "")
        return _result(
            str(r["response"]),
            int(r.get("ms", 0)),
            int(r.get("eval_count", 0)),
            float(r.get("tok_per_sec", 0.0)),
        )


def build_local_backend(name: str, s: settings.Settings) -> Backend:
    if name == "ollama":
        return OllamaBackend(s.harness_base_url("ollama"))
    if name == "lmstudio":
        return LMStudioBackend(s.harness_base_url("lmstudio"))
    raise BackendError(f"{name!r} is not a local harness (expected 'ollama' or 'lmstudio')")
