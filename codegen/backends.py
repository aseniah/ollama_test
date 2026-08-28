"""Inference backends for the codegen benchmark: Ollama, LM Studio, Apple."""

from __future__ import annotations

import json
import shutil
import subprocess
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
    def warmup(
        self, messages: list[dict[str, str]], options: dict[str, Any], model: str, timeout: int
    ) -> None: ...
    def unload(self, model: str) -> None: ...
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


def _lms_path() -> str | None:
    lms = shutil.which("lms") or str(Path.home() / ".lmstudio" / "bin" / "lms")
    return lms if Path(lms).exists() else None


def _port_of(base_url: str) -> str | None:
    tail = base_url.rsplit(":", 1)[-1].split("/", 1)[0]
    return tail if tail.isdigit() else None


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

    def warmup(
        self, messages: list[dict[str, str]], options: dict[str, Any], model: str, timeout: int
    ) -> None:
        print(f"  ollama: loading {model}...", flush=True)
        # pin the model in memory for the whole run
        self._post(
            f"{self._base}/api/chat",
            {"model": model, "messages": [], "keep_alive": -1},
            60,
        )
        # one warm inference (skip thinking — no need to burn tokens on warmup)
        warm: dict[str, Any] = {**options, "think": False} if options.get("think") is True else dict(options)
        self._post(
            f"{self._base}/api/chat",
            {"model": model, "messages": messages, "stream": False, **warm},
            timeout,
        )

    def unload(self, model: str) -> None:
        try:
            self._post(
                f"{self._base}/api/chat",
                {"model": model, "messages": [], "keep_alive": 0},
                15,
            )
        except BackendError:
            pass

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

    def __init__(self, base_url: str, autostart: bool = False, _post: PostFn = _http_post) -> None:
        self._base = base_url.rstrip("/")
        self._autostart = autostart
        self._post = _post
        self._started = False

    def _models_url(self) -> str:
        return f"{self._base}/v1/models"

    def start(self) -> None:
        if _http_get_ok(self._models_url()):
            return
        if not self._autostart:
            raise BackendError(
                f"LM Studio not reachable at {self._base} — start the LM Studio server "
                "(or set autostart = true under [harness.lmstudio] in settings.toml)"
            )
        lms = _lms_path()
        if lms is None:
            raise BackendError(
                f"LM Studio not reachable at {self._base} and the `lms` CLI was not found "
                "for autostart — start the server manually"
            )
        cmd = [lms, "server", "start"]
        port = _port_of(self._base)
        if port:
            cmd += ["--port", port]
        print("  lmstudio: starting server via `lms server start`...", flush=True)
        subprocess.run(cmd, capture_output=True, timeout=60, check=False)
        for _ in range(20):
            time.sleep(0.5)
            if _http_get_ok(self._models_url(), timeout=2):
                self._started = True
                print("  lmstudio: server ready", flush=True)
                return
        raise BackendError(f"LM Studio server did not become ready at {self._base} within 10s")

    def stop(self) -> None:
        if not self._started:
            return
        lms = _lms_path()
        if lms is None:
            return
        print("  lmstudio: stopping server...", flush=True)
        try:
            subprocess.run([lms, "server", "stop"], capture_output=True, timeout=30, check=False)
        except (OSError, subprocess.SubprocessError):
            pass
        self._started = False

    def warmup(
        self, messages: list[dict[str, str]], options: dict[str, Any], model: str, timeout: int
    ) -> None:
        # Force the JIT load now so the first test's timing is not skewed.
        print(f"  lmstudio: loading {model}...", flush=True)
        passthrough = {k: v for k, v in options.items() if k != "think"}
        try:
            self._post(
                f"{self._base}/v1/chat/completions",
                {"model": model, "messages": messages, "stream": False, **passthrough},
                timeout,
            )
        except BackendError as e:
            print(f"    lmstudio warmup failed (continuing): {e}", flush=True)

    def unload(self, model: str) -> None:
        """Best-effort: `lms unload <model>`. Falls back to LM Studio's own
        JIT auto-evict if the CLI is missing or the id does not match."""
        lms = _lms_path()
        if lms is None:
            return
        try:
            subprocess.run(
                [lms, "unload", model], capture_output=True, timeout=30, check=False
            )
        except (OSError, subprocess.SubprocessError):
            pass

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

    def warmup(
        self, messages: list[dict[str, str]], options: dict[str, Any], model: str, timeout: int
    ) -> None:
        apfel_backend.warmup(messages)

    def unload(self, model: str) -> None:
        return None

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
        return LMStudioBackend(
            s.harness_base_url("lmstudio"), autostart=s.harness_autostart("lmstudio")
        )
    raise BackendError(f"{name!r} is not a local harness (expected 'ollama' or 'lmstudio')")
