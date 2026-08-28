# codegen Settings File + LM Studio Harness Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Move codegen model/runtime config into `codegen/settings.toml`, add LM Studio as an optional local harness alongside Ollama, and restructure results into `results/v{NNN}/{harness}/{model}/`.

**Architecture:** A new `settings.py` loads and validates `settings.toml` (stdlib `tomllib`). A new `backends.py` defines a `Backend` protocol with `OllamaBackend`, `LMStudioBackend`, and an `AppleBackend` adapter over the repo-root `apfel_backend`. `benchmark.py` and `run_claude_test.py` consume settings, dispatch through a backend, and write results under a per-harness directory segment. A standalone `migrate_results.py` moves existing result dirs into the new layout.

**Tech Stack:** Python 3.14 (stdlib only — `tomllib`, `urllib`, `argparse`, `unittest`). No new dependencies. Lint: `ruff` (E/F/W/ANN, E501 ignored) + `pyright` with `reportUnknown*` = error — all new code must be fully type-annotated and pyright-clean.

**Spec:** `docs/superpowers/specs/2026-08-27-codegen-settings-harness-design.md`

## Global Constraints

- Scope is `codegen/` only. Do not modify `shell/` or the repo-root `apfel_backend.py`.
- Stdlib only. No `pip install`, no additions to `pyproject.toml` dependencies.
- Python floor 3.11 (`tomllib`), though the dev machine runs 3.14.
- `PROMPT_VERSION` stays `2`. Do not bump it.
- All new `.py` files must pass `bash lint.sh` from `codegen/` (ruff + pyright, both clean).
- New JSONL record fields are additive only. Never remove or rename an existing field.
- Unit tests live in `codegen/tests/` as `test_*.py`, run with `python3 -m unittest discover -s tests -t . -p 'test_*.py'` from `codegen/`. `load_tests()` in `benchmark.py` ignores non-directories, so these files do not interfere with benchmark test discovery.
- Results directories are git-tracked; moves use `git mv`.
- Harness names are exactly: `ollama`, `lmstudio`, `apple`, `anthropic`.

---

## File Structure

| File | Responsibility |
|---|---|
| `codegen/settings.toml` | **Create.** Committed config data — harnesses, defaults, local model list, anthropic model list. Replaces the inline `MODELS` list. |
| `codegen/settings.py` | **Create.** `load_settings()` + `Settings` accessor object + validation. Stdlib `tomllib`. |
| `codegen/backends.py` | **Create.** `GenResult`, `Backend` protocol, `OllamaBackend`, `LMStudioBackend`, `AppleBackend`, `build_local_backend()`. |
| `codegen/migrate_results.py` | **Create.** One-time CLI to move `results/vNNN/*` into `results/vNNN/{harness}/*`. |
| `codegen/tests/test_settings.py` | **Create.** Unit tests for `settings.py`. |
| `codegen/tests/test_backends.py` | **Create.** Unit tests for `backends.py` against recorded HTTP fixtures. |
| `codegen/tests/test_migrate_results.py` | **Create.** Unit tests for `migrate_results.py` classification + moves on a temp tree. |
| `codegen/benchmark.py` | **Modify.** `main()` + `run_one()` use settings + backend; results path gains harness segment; record gains `harness`; `--harness` flag; summary shows tok/s. |
| `codegen/run_claude_test.py` | **Modify.** `MODELS` read from settings; results path gains `anthropic/` segment; record gains `harness`. |
| `codegen/CLAUDE.md` | **Modify.** "Structure", "Running", "Analyzing Results" sections — new paths, `--harness`, `settings.toml`, `harness` field, updated jq globs. |
| `codegen/results/findings_instructions.md` | **Modify.** Harness subdirs, new Harness Comparison section, mandatory tok/s. |
| `codegen/README.md` | **Modify.** `settings.toml` overview, `--harness`, LM Studio setup. |

---

## Task 1: `settings.toml` + `settings.py` loader

**Files:**
- Create: `codegen/settings.toml`
- Create: `codegen/settings.py`
- Create: `codegen/tests/test_settings.py`

**Interfaces:**
- Consumes: nothing (first task).
- Produces:
  - `settings.py`:
    - `class SettingsError(Exception)`
    - `class ModelEntry(TypedDict)`: `name: str`, `options: dict[str, Any]`, `infer_timeout: int`, `exec_timeout: int`, `lmstudio_model: str`
    - `class AnthropicEntry(TypedDict)`: `alias: str`, `model_id: str`
    - `class Settings` with methods:
      - `default_harness() -> str`
      - `harness_base_url(name: str) -> str`
      - `apple_autostart() -> bool`
      - `languages() -> list[str]`
      - `local_models() -> list[ModelEntry]` (enabled only, `[defaults]` applied, `lmstudio_model` defaults to `name`)
      - `anthropic_models() -> list[AnthropicEntry]` (enabled only)
      - `anthropic_default_alias() -> str` (first enabled alias, or `"haiku"` if none)
    - `def load_settings(path: Path = Path("settings.toml")) -> Settings`

- [ ] **Step 1: Write `codegen/settings.toml`**

```toml
# codegen benchmark configuration.
# Edit this file to change which models run and which local harness they run against.

[harness]
default = "ollama"
ollama   = { base_url = "http://localhost:11434" }
lmstudio = { base_url = "http://localhost:1234" }
apple    = { base_url = "http://localhost:11435", autostart = true }

[defaults]
infer_timeout = 120   # seconds; per-model override allowed below
exec_timeout  = 60    # seconds; per-model override allowed below
languages = ["python", "typescript", "go", "csharp"]

# Local models. Run against whichever harness --harness selects (default: ollama).
# `enabled = false` keeps a model in the file but out of the run.
# `lmstudio_model` overrides the identifier sent to LM Studio when it differs from `name`.
models = [
  { name = "qwen3.8:27b",                  think = false, enabled = true },
  { name = "qwen3.8:27b",                  think = true,  enabled = true,  infer_timeout = 300 },
  { name = "qwen3-coder:30b",              think = false, enabled = false },
  { name = "qwen2.5-coder:7b",             think = false, enabled = false },
  { name = "gemma4:26b",                   think = false, enabled = false },
  { name = "gemma4:26b",                   think = true,  enabled = false, infer_timeout = 300 },
  { name = "gemma4:31b",                   think = false, enabled = false },
  { name = "gemma4:31b",                   think = true,  enabled = false, infer_timeout = 300 },
  { name = "qwen3.6:35b",                  think = false, enabled = false },
  { name = "qwen3.6:35b",                  think = true,  enabled = false, infer_timeout = 300 },
  { name = "qwen3.6:35b-a3b-coding-nvfp4", think = false, enabled = false },
  { name = "qwen3.6:35b-a3b-coding-nvfp4", think = true,  enabled = false, infer_timeout = 300 },
  { name = "qwen3.5:35b-a3b-coding-nvfp4", think = false, enabled = false },
  { name = "qwen3.5:35b-a3b-coding-nvfp4", think = true,  enabled = false, infer_timeout = 300 },
  { name = "qwen3.5:27b",                  think = false, enabled = false },
  { name = "qwen3.5:27b",                  think = true,  enabled = false, infer_timeout = 300 },
  { name = "qwen3.5:27b-nvfp4",            think = false, enabled = false },
  { name = "qwen3.5:27b-nvfp4",            think = true,  enabled = false, infer_timeout = 300 },
  { name = "qwen3.5:4b",                   think = false, enabled = false },
  { name = "qwen3.5:4b",                   think = true,  enabled = false, infer_timeout = 300 },
  { name = "qwen3.5:4b-nvfp4",             think = false, enabled = false },
  { name = "qwen3.5:4b-nvfp4",             think = true,  enabled = false, infer_timeout = 300 },
]

# Anthropic models for run_claude_test.py (harness = "anthropic").
anthropic_models = [
  { alias = "haiku",  model_id = "claude-haiku-4-5-20251001", enabled = true },
  { alias = "sonnet", model_id = "claude-sonnet-5",           enabled = false },
  { alias = "opus",   model_id = "claude-opus-5",             enabled = false },
]
```

- [ ] **Step 2: Write the failing tests — `codegen/tests/test_settings.py`**

```python
import tempfile
import unittest
from pathlib import Path

import settings


def _write(text: str) -> Path:
    d = Path(tempfile.mkdtemp())
    p = d / "settings.toml"
    p.write_text(text)
    return p


MINIMAL = """
[harness]
default = "ollama"
ollama   = { base_url = "http://localhost:11434" }
lmstudio = { base_url = "http://localhost:1234" }
apple    = { base_url = "http://localhost:11435", autostart = true }

[defaults]
infer_timeout = 120
exec_timeout  = 60
languages = ["python", "go"]

models = [
  { name = "m:1", think = false, enabled = true },
  { name = "m:2", think = true,  enabled = true, infer_timeout = 300 },
  { name = "m:3", think = false, enabled = false },
  { name = "m:4", think = false, enabled = true, lmstudio_model = "m-4-lms" },
]

anthropic_models = [
  { alias = "haiku",  model_id = "claude-haiku-x", enabled = true },
  { alias = "sonnet", model_id = "claude-sonnet-x", enabled = false },
]
"""


class LoadTests(unittest.TestCase):
    def setUp(self) -> None:
        self.s = settings.load_settings(_write(MINIMAL))

    def test_default_harness(self) -> None:
        self.assertEqual(self.s.default_harness(), "ollama")

    def test_harness_base_url(self) -> None:
        self.assertEqual(self.s.harness_base_url("lmstudio"), "http://localhost:1234")

    def test_apple_autostart(self) -> None:
        self.assertTrue(self.s.apple_autostart())

    def test_languages(self) -> None:
        self.assertEqual(self.s.languages(), ["python", "go"])

    def test_local_models_enabled_only(self) -> None:
        names = [m["name"] for m in self.s.local_models()]
        self.assertEqual(names, ["m:1", "m:2", "m:4"])

    def test_local_models_defaults_applied(self) -> None:
        m1 = self.s.local_models()[0]
        self.assertEqual(m1["infer_timeout"], 120)
        self.assertEqual(m1["exec_timeout"], 60)
        self.assertEqual(m1["options"], {"think": False})

    def test_local_models_per_model_override(self) -> None:
        m2 = self.s.local_models()[1]
        self.assertEqual(m2["infer_timeout"], 300)
        self.assertEqual(m2["options"], {"think": True})

    def test_lmstudio_model_defaults_to_name(self) -> None:
        m1 = self.s.local_models()[0]
        self.assertEqual(m1["lmstudio_model"], "m:1")

    def test_lmstudio_model_override(self) -> None:
        m4 = self.s.local_models()[2]
        self.assertEqual(m4["lmstudio_model"], "m-4-lms")

    def test_anthropic_models_enabled_only(self) -> None:
        aliases = [a["alias"] for a in self.s.anthropic_models()]
        self.assertEqual(aliases, ["haiku"])

    def test_anthropic_default_alias(self) -> None:
        self.assertEqual(self.s.anthropic_default_alias(), "haiku")


class ValidationTests(unittest.TestCase):
    def test_missing_file(self) -> None:
        with self.assertRaises(settings.SettingsError):
            settings.load_settings(Path("/nonexistent/settings.toml"))

    def test_bad_toml(self) -> None:
        with self.assertRaises(settings.SettingsError):
            settings.load_settings(_write("this is = not valid toml ["))

    def test_unknown_default_harness(self) -> None:
        bad = MINIMAL.replace('default = "ollama"', 'default = "bogus"')
        with self.assertRaises(settings.SettingsError):
            settings.load_settings(_write(bad))

    def test_harness_without_base_url(self) -> None:
        bad = MINIMAL.replace('ollama   = { base_url = "http://localhost:11434" }', 'ollama   = { }')
        with self.assertRaises(settings.SettingsError):
            settings.load_settings(_write(bad))

    def test_no_enabled_models(self) -> None:
        bad = MINIMAL.replace("enabled = true", "enabled = false")
        with self.assertRaises(settings.SettingsError):
            settings.load_settings(_write(bad))

    def test_empty_languages(self) -> None:
        bad = MINIMAL.replace('languages = ["python", "go"]', "languages = []")
        with self.assertRaises(settings.SettingsError):
            settings.load_settings(_write(bad))

    def test_anthropic_default_alias_fallback(self) -> None:
        bad = MINIMAL.replace('{ alias = "haiku",  model_id = "claude-haiku-x", enabled = true }',
                              '{ alias = "haiku",  model_id = "claude-haiku-x", enabled = false }')
        s = settings.load_settings(_write(bad))
        self.assertEqual(s.anthropic_default_alias(), "haiku")


if __name__ == "__main__":
    unittest.main()
```

- [ ] **Step 3: Run tests to verify they fail**

Run: `cd codegen && python3 -m unittest discover -s tests -t . -p 'test_settings.py' -v`
Expected: FAIL — `ModuleNotFoundError: No module named 'settings'`

- [ ] **Step 4: Write `codegen/settings.py`**

```python
"""Load and validate codegen/settings.toml."""

from __future__ import annotations

import tomllib
from pathlib import Path
from typing import Any, TypedDict

_HARNESSES = ("ollama", "lmstudio", "apple")


class SettingsError(Exception):
    """Raised for any malformed or missing settings file."""


class ModelEntry(TypedDict):
    name: str
    options: dict[str, Any]
    infer_timeout: int
    exec_timeout: int
    lmstudio_model: str


class AnthropicEntry(TypedDict):
    alias: str
    model_id: str


class Settings:
    def __init__(self, raw: dict[str, Any]) -> None:
        self._raw = raw

    def default_harness(self) -> str:
        return str(self._raw["harness"]["default"])

    def harness_base_url(self, name: str) -> str:
        return str(self._raw["harness"][name]["base_url"])

    def apple_autostart(self) -> bool:
        return bool(self._raw["harness"]["apple"].get("autostart", True))

    def languages(self) -> list[str]:
        return [str(x) for x in self._raw["defaults"]["languages"]]

    def local_models(self) -> list[ModelEntry]:
        d = self._raw["defaults"]
        out: list[ModelEntry] = []
        for m in self._raw["models"]:
            if not m.get("enabled", False):
                continue
            out.append(ModelEntry(
                name=str(m["name"]),
                options={"think": bool(m.get("think", False))},
                infer_timeout=int(m.get("infer_timeout", d["infer_timeout"])),
                exec_timeout=int(m.get("exec_timeout", d["exec_timeout"])),
                lmstudio_model=str(m.get("lmstudio_model", m["name"])),
            ))
        return out

    def anthropic_models(self) -> list[AnthropicEntry]:
        return [
            AnthropicEntry(alias=str(a["alias"]), model_id=str(a["model_id"]))
            for a in self._raw.get("anthropic_models", [])
            if a.get("enabled", False)
        ]

    def anthropic_default_alias(self) -> str:
        enabled = self.anthropic_models()
        if enabled:
            return enabled[0]["alias"]
        return "haiku"


def _validate(raw: dict[str, Any]) -> None:
    if "harness" not in raw or "default" not in raw["harness"]:
        raise SettingsError("settings.toml: missing [harness].default")
    default = raw["harness"]["default"]
    if default not in _HARNESSES:
        raise SettingsError(f"settings.toml: [harness].default = {default!r} is not one of {_HARNESSES}")
    for h in _HARNESSES:
        if h not in raw["harness"] or "base_url" not in raw["harness"][h]:
            raise SettingsError(f"settings.toml: [harness].{h} needs a base_url")
    if "defaults" not in raw:
        raise SettingsError("settings.toml: missing [defaults]")
    langs = raw["defaults"].get("languages")
    if not isinstance(langs, list) or not langs:
        raise SettingsError("settings.toml: [defaults].languages must be a non-empty list")
    for key in ("infer_timeout", "exec_timeout"):
        if key not in raw["defaults"]:
            raise SettingsError(f"settings.toml: [defaults].{key} is required")
    models = raw.get("models")
    if not isinstance(models, list) or not models:
        raise SettingsError("settings.toml: `models` must be a non-empty list")
    if not any(m.get("enabled", False) for m in models):
        raise SettingsError("settings.toml: no enabled models — set enabled = true on at least one")


def load_settings(path: Path = Path("settings.toml")) -> Settings:
    try:
        raw = tomllib.loads(path.read_text())
    except FileNotFoundError as e:
        raise SettingsError(f"settings.toml not found at {path}") from e
    except tomllib.TOMLDecodeError as e:
        raise SettingsError(f"settings.toml is not valid TOML: {e}") from e
    _validate(raw)
    return Settings(raw)
```

- [ ] **Step 5: Run tests to verify they pass**

Run: `cd codegen && python3 -m unittest discover -s tests -t . -p 'test_settings.py' -v`
Expected: PASS (all tests green)

- [ ] **Step 6: Lint**

Run: `cd codegen && bash lint.sh`
Expected: ruff + pyright both clean. Fix any `reportUnknown*` by adding annotations.

- [ ] **Step 7: Commit**

```bash
git add codegen/settings.toml codegen/settings.py codegen/tests/test_settings.py
git commit -m "feat(codegen): add settings.toml config loader"
```

---

## Task 2: `backends.py`

**Files:**
- Create: `codegen/backends.py`
- Create: `codegen/tests/test_backends.py`

**Interfaces:**
- Consumes: `settings.Settings`, `settings.ModelEntry` (Task 1).
- Produces:
  - `class GenResult(TypedDict)`: `response: str`, `ms: int`, `eval_count: int`, `tok_per_sec: float`
  - `class BackendError(Exception)`
  - `class Backend(Protocol)`: attribute `name: str`; methods `start() -> None`, `stop() -> None`, `warmup(messages: list[dict[str, str]]) -> None`, `generate(messages: list[dict[str, str]], options: dict[str, Any], timeout: int, model: str) -> GenResult`
  - `class OllamaBackend`, `class LMStudioBackend`, `class AppleBackend` implementing `Backend`
  - `def build_local_backend(name: str, s: settings.Settings) -> Backend` — returns `OllamaBackend` or `LMStudioBackend`; raises `BackendError` for any other name

- [ ] **Step 1: Write the failing tests — `codegen/tests/test_backends.py`**

```python
import json
import unittest
from typing import Any

import backends


class _FakeHTTP:
    """Records the last request and returns a canned JSON body."""

    def __init__(self, body: dict[str, Any]) -> None:
        self.body = body
        self.last_url: str | None = None
        self.last_payload: dict[str, Any] | None = None

    def __call__(self, url: str, payload: dict[str, Any], timeout: int) -> dict[str, Any]:
        self.last_url = url
        self.last_payload = payload
        return self.body


OLLAMA_BODY = {
    "message": {"content": "  print('x')  "},
    "eval_count": 40,
    "eval_duration": 2_000_000_000,  # 2s in ns -> 40/2 = 20 tok/s
}

LMSTUDIO_V0_BODY = {
    "choices": [{"message": {"content": "print('x')"}}],
    "stats": {"tokens_per_second": 33.3, "predicted_tokens_count": 50},
    "usage": {"completion_tokens": 50},
}


class OllamaBackendTests(unittest.TestCase):
    def test_generate_shape_and_timing(self) -> None:
        http = _FakeHTTP(OLLAMA_BODY)
        b = backends.OllamaBackend("http://x:11434", _post=http)
        r = b.generate([{"role": "user", "content": "hi"}], {"think": False}, 120, "m:1")
        self.assertEqual(r["response"], "print('x')")  # stripped
        self.assertEqual(r["eval_count"], 40)
        self.assertEqual(r["tok_per_sec"], 20.0)
        self.assertGreaterEqual(r["ms"], 0)
        assert http.last_url == "http://x:11434/api/chat"
        assert http.last_payload is not None
        self.assertEqual(http.last_payload["model"], "m:1")
        self.assertFalse(http.last_payload["stream"])
        self.assertFalse(http.last_payload["think"])


class LMStudioBackendTests(unittest.TestCase):
    def test_generate_prefers_v0_stats(self) -> None:
        http = _FakeHTTP(LMSTUDIO_V0_BODY)
        b = backends.LMStudioBackend("http://x:1234", _post=http)
        r = b.generate([{"role": "user", "content": "hi"}], {"think": True}, 120, "m:1")
        self.assertEqual(r["response"], "print('x')")
        self.assertEqual(r["tok_per_sec"], 33.3)
        self.assertEqual(r["eval_count"], 50)
        assert http.last_url == "http://x:1234/api/v0/chat/completions"
        assert http.last_payload is not None
        self.assertEqual(http.last_payload["model"], "m:1")
        self.assertNotIn("think", http.last_payload)  # not an OpenAI-compat param


class FactoryTests(unittest.TestCase):
    def test_build_local_backend_names(self) -> None:
        import settings
        s = settings.Settings({
            "harness": {
                "default": "ollama",
                "ollama": {"base_url": "http://o:11434"},
                "lmstudio": {"base_url": "http://l:1234"},
                "apple": {"base_url": "http://a:11435"},
            },
            "defaults": {"languages": ["python"], "infer_timeout": 120, "exec_timeout": 60},
            "models": [{"name": "m", "enabled": True}],
        })
        self.assertIsInstance(backends.build_local_backend("ollama", s), backends.OllamaBackend)
        self.assertIsInstance(backends.build_local_backend("lmstudio", s), backends.LMStudioBackend)
        with self.assertRaises(backends.BackendError):
            backends.build_local_backend("apple", s)


if __name__ == "__main__":
    unittest.main()
```

- [ ] **Step 2: Run tests to verify they fail**

Run: `cd codegen && python3 -m unittest discover -s tests -t . -p 'test_backends.py' -v`
Expected: FAIL — `ModuleNotFoundError: No module named 'backends'`

- [ ] **Step 3: Write `codegen/backends.py`**

```python
"""Inference backends for the codegen benchmark: Ollama, LM Studio, Apple."""

from __future__ import annotations

import json
import sys
import time
import urllib.error
import urllib.request
from pathlib import Path
from typing import Any, Callable, Protocol, TypedDict, cast

sys.path.insert(0, str(Path(__file__).parent.parent))
import apfel_backend  # noqa: E402  (repo-root shared module)

import settings  # noqa: E402

PostFn = Callable[[str, dict[str, Any], int], dict[str, Any]]


class GenResult(TypedDict):
    response: str
    ms: int
    eval_count: int
    tok_per_sec: float


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


def _http_post(url: str, payload: dict[str, Any], timeout: int) -> dict[str, Any]:
    data = json.dumps(payload).encode()
    req = urllib.request.Request(url, data=data, headers={"Content-Type": "application/json"}, method="POST")
    try:
        with urllib.request.urlopen(req, timeout=timeout) as resp:
            return cast(dict[str, Any], json.loads(resp.read()))
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
            raise BackendError(f"Ollama not reachable at {self._base} — is `ollama serve` running?")

    def stop(self) -> None:
        return None

    def warmup(self, messages: list[dict[str, str]]) -> None:
        self._post(f"{self._base}/api/chat", {"model": "", "messages": [], "keep_alive": -1}, 60)

    def generate(
        self, messages: list[dict[str, str]], options: dict[str, Any], timeout: int, model: str
    ) -> GenResult:
        payload: dict[str, Any] = {"model": model, "messages": messages, "stream": False, **options}
        start = time.monotonic()
        data = self._post(f"{self._base}/api/chat", payload, timeout)
        ms = int((time.monotonic() - start) * 1000)
        eval_count = int(data.get("eval_count", 0))
        eval_dur = int(data.get("eval_duration", 0)) or 1
        tok_per_sec = round(eval_count / (eval_dur / 1e9), 1) if eval_dur else 0.0
        return GenResult(
            response=str(data["message"]["content"]).strip(),
            ms=ms,
            eval_count=eval_count,
            tok_per_sec=tok_per_sec,
        )


class LMStudioBackend:
    name = "lmstudio"

    def __init__(self, base_url: str, _post: PostFn = _http_post) -> None:
        self._base = base_url.rstrip("/")
        self._post = _post

    def start(self) -> None:
        if not _http_get_ok(f"{self._base}/v1/models"):
            raise BackendError(
                f"LM Studio not reachable at {self._base} — start the LM Studio server and load a model"
            )

    def stop(self) -> None:
        return None

    def warmup(self, messages: list[dict[str, str]]) -> None:
        return None

    def generate(
        self, messages: list[dict[str, str]], options: dict[str, Any], timeout: int, model: str
    ) -> GenResult:
        # `think` is Ollama-only; drop it. Pass anything else through unchanged.
        passthrough = {k: v for k, v in options.items() if k != "think"}
        payload: dict[str, Any] = {"model": model, "messages": messages, "stream": False, **passthrough}
        start = time.monotonic()
        try:
            data = self._post(f"{self._base}/api/v0/chat/completions", payload, timeout)
            stats = data.get("stats") or {}
        except BackendError:
            data = self._post(f"{self._base}/v1/chat/completions", payload, timeout)
            stats = {}
        ms = int((time.monotonic() - start) * 1000)
        content = str(data["choices"][0]["message"]["content"]).strip()
        usage = data.get("usage") or {}
        eval_count = int(stats.get("predicted_tokens_count") or usage.get("completion_tokens") or 0)
        if stats.get("tokens_per_second"):
            tok_per_sec = round(float(stats["tokens_per_second"]), 1)
        else:
            tok_per_sec = round(eval_count / (ms / 1000), 1) if ms else 0.0
        return GenResult(response=content, ms=ms, eval_count=eval_count, tok_per_sec=tok_per_sec)


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
        r = apfel_backend.run_prompt(messages, "", "")
        return GenResult(
            response=str(r["response"]),
            ms=int(r.get("ms", 0)),
            eval_count=int(r.get("eval_count", 0)),
            tok_per_sec=float(r.get("tok_per_sec", 0.0)),
        )


def build_local_backend(name: str, s: settings.Settings) -> Backend:
    if name == "ollama":
        return OllamaBackend(s.harness_base_url("ollama"))
    if name == "lmstudio":
        return LMStudioBackend(s.harness_base_url("lmstudio"))
    raise BackendError(f"{name!r} is not a local harness (expected 'ollama' or 'lmstudio')")
```

- [ ] **Step 4: Run tests to verify they pass**

Run: `cd codegen && python3 -m unittest discover -s tests -t . -p 'test_backends.py' -v`
Expected: PASS. (Note: `_FakeHTTP` bypasses `start()`, so no live server is needed.)

- [ ] **Step 5: Lint**

Run: `cd codegen && bash lint.sh`
Expected: clean. `apfel_backend` is at the repo root and pyright has `extraPaths = [".."]`; if pyright still cannot resolve members, add `# type: ignore[...]` narrowly on the `apfel_backend.*` calls inside `AppleBackend` only.

- [ ] **Step 6: Commit**

```bash
git add codegen/backends.py codegen/tests/test_backends.py
git commit -m "feat(codegen): add Ollama/LMStudio/Apple backend abstraction"
```

---

## Task 3: Wire `benchmark.py` to settings + backends

**Files:**
- Modify: `codegen/benchmark.py` — imports, `MODELS` removal, `main()` (687-816), `run_one()` (593-680), `results_file()` (531-534), `append_results()` (537-540), artifact path (759), `print_tables()` tok/s line.

**Interfaces:**
- Consumes: `settings.load_settings`, `settings.Settings`, `settings.ModelEntry` (Task 1); `backends.Backend`, `backends.GenResult`, `backends.build_local_backend`, `backends.AppleBackend`, `backends.BackendError` (Task 2).
- Produces: JSONL records with a new `"harness": str` field. Results at `results/v{PROMPT_VERSION:03d}/{harness}/{model_safe}/results.jsonl`.

- [ ] **Step 1: Replace the `MODELS` list and add imports**

In `codegen/benchmark.py`, delete the entire `MODELS: list[ModelConfig] = [ ... ]` block (currently lines ~100-121) and its section header comment. Keep the `ModelConfig` TypedDict — `run_one` still uses its shape, now built from settings.

Add near the top imports (after `import apfel_backend`):

```python
import backends
import settings as settings_mod
```

- [ ] **Step 2: Add `harness` param to `run_one` and record; dispatch through backend**

Change `run_one` signature (line 593) from:

```python
def run_one(
    model_cfg: ModelConfig,
    variant: PromptVariant,
    test: TestCase,
    run_id: str,
    timestamp: str,
    artifact_base: Path,
) -> dict[str, Any]:
```

to:

```python
def run_one(
    backend: backends.Backend,
    harness: str,
    model_cfg: ModelConfig,
    variant: PromptVariant,
    test: TestCase,
    run_id: str,
    timestamp: str,
    artifact_base: Path,
) -> dict[str, Any]:
```

Replace the inference block (currently lines ~604-628, the `is_apfel` branch through `tok_per_sec = ...`) with:

```python
    model = model_cfg["name"]
    options = model_cfg["options"]
    infer_timeout = model_cfg.get("infer_timeout", INFER_TIMEOUT)
    result = backend.generate(messages, options, infer_timeout, model)
    response_raw = result["response"]
    ms = result["ms"]
    eval_count = result["eval_count"]
    tok_per_sec = result["tok_per_sec"]
```

In the `record` dict (line ~653), add after `"model_options": options,`:

```python
        "harness":         harness,
```

- [ ] **Step 3: Add harness segment to results + artifact paths**

Change `results_file` (line 531):

```python
def results_file(harness: str, model_safe: str) -> Path:
    path = RESULTS_DIR / f"v{PROMPT_VERSION:03d}" / harness / model_safe / "results.jsonl"
    path.parent.mkdir(parents=True, exist_ok=True)
    return path
```

Change `append_results` (line 537):

```python
def append_results(records: list[dict[str, Any]], harness: str, model_safe: str) -> None:
    with open(results_file(harness, model_safe), "a") as f:
        for r in records:
            f.write(json.dumps(r) + "\n")
```

- [ ] **Step 4: Rewrite `main()` to load settings, build backends, loop**

Replace `main()` (lines 687-817) with:

```python
def main() -> None:
    parser = argparse.ArgumentParser(description="Code Generation Benchmark")
    parser.add_argument("runs", nargs="?", type=int, default=1, help="number of times to run (default: 1)")
    parser.add_argument("--harness", choices=["ollama", "lmstudio"], default=None,
                        help="local inference harness (default: [harness].default in settings.toml)")
    parser.add_argument("--apple", action="store_true", help="also run the Apple on-device model via apfel")
    args = parser.parse_args()

    try:
        cfg = settings_mod.load_settings(Path("settings.toml"))
    except settings_mod.SettingsError as e:
        print(f"ERROR: {e}", file=sys.stderr)
        sys.exit(1)

    global LANGUAGES
    LANGUAGES = cfg.languages()

    global _test_cases
    _test_cases = load_tests(Path("tests"))
    if not _test_cases:
        print("ERROR: no tests found in tests/", file=sys.stderr)
        sys.exit(1)

    check_runtimes(LANGUAGES)

    harness_name = args.harness or cfg.default_harness()
    try:
        local_backend = backends.build_local_backend(harness_name, cfg)
        local_backend.start()
    except backends.BackendError as e:
        print(f"ERROR: {e}", file=sys.stderr)
        sys.exit(1)

    # (backend, harness, model_cfg) queue
    queue: list[tuple[backends.Backend, str, ModelConfig]] = []
    for m in cfg.local_models():
        model_name = m["lmstudio_model"] if harness_name == "lmstudio" else m["name"]
        queue.append((local_backend, harness_name, ModelConfig(
            name=model_name, options=m["options"],
            infer_timeout=m["infer_timeout"], exec_timeout=m["exec_timeout"],
        )))

    apple_backend_obj: backends.AppleBackend | None = None
    if args.apple:
        apple_backend_obj = backends.AppleBackend(autostart=cfg.apple_autostart())
        try:
            apple_backend_obj.start()
        except Exception as e:  # apfel_backend raises RuntimeError
            print(f"ERROR: could not start apfel: {e}", file=sys.stderr)
            sys.exit(1)
        queue.append((apple_backend_obj, "apple", ModelConfig(
            name=apfel_backend.MODEL_NAME, options={}, infer_timeout=INFER_TIMEOUT, exec_timeout=EXEC_TIMEOUT,
        )))

    warmup_messages = build_messages(
        PROMPT_VARIANTS[0],
        {"id": "warmup", "prompt": "Write a Python function that returns 'Hello, World!'", "verify_mod": None},  # type: ignore[typeddict-item]
        "python",
    )

    all_records: list[dict[str, Any]] = []
    try:
        for backend, harness, model_cfg in queue:
            model = model_cfg["name"]
            sep = _c('=' * 60, _CYAN)
            print(f"\n{sep}")
            print(f"Model: {model}  harness={harness}  options={model_cfg['options']}")
            print(sep)

            try:
                if harness in ("ollama",):
                    preload(model, warmup_messages, model_cfg["options"],
                            model_cfg.get("infer_timeout", INFER_TIMEOUT))
                else:
                    backend.warmup(warmup_messages)
            except Exception as e:
                print(f"  ERROR warming up: {e}", file=sys.stderr)
                continue

            think = model_cfg["options"].get("think")
            think_suffix = " think" if think is True else " nothink" if think is False else ""
            model_label = f"{model}{think_suffix}"
            model_safe = model.replace(":", "_").replace("/", "_")
            if think is True:
                model_safe += "_think"
            elif think is False:
                model_safe += "_nothink"

            for run_num in range(1, args.runs + 1):
                if args.runs > 1:
                    rsep = _c('#' * 50, _PURPLE)
                    print(f"\n  {rsep}")
                    print(f"  {_c(' ' * 20, _WHITE)}Run {run_num} of {args.runs}")
                    print(f"  {rsep}")
                run_ctx = f"  run {run_num}/{args.runs}" if args.runs > 1 else ""
                run_id = datetime.datetime.now().strftime("%Y%m%d_%H%M%S")
                timestamp = datetime.datetime.now(datetime.timezone.utc).isoformat()
                artifact_base = RESULTS_DIR / f"v{PROMPT_VERSION:03d}" / harness / model_safe / run_id
                run_records: list[dict[str, Any]] = []

                total = len(PROMPT_VARIANTS) * len(_test_cases)
                idx = 0
                for variant in PROMPT_VARIANTS:
                    for test in _test_cases:
                        idx += 1
                        lang = variant["language"]
                        lang_str = _c(LANG_NAME[lang], LANG_COLOR[lang])
                        count_str = _c(f"[{idx:02d}/{total}]", _WHITE)
                        model_str = _c(f" {model_label} ", _BG_GRAY)
                        print(f"  {count_str} {model_str}{run_ctx}  {lang_str}  {test['id']}", flush=True)
                        try:
                            record = run_one(backend, harness, model_cfg, variant, test, run_id, timestamp, artifact_base)
                        except Exception as e:
                            print(f"    ERROR: {e}", file=sys.stderr)
                            record = {
                                "run_id": run_id, "timestamp": timestamp, "prompt_v": PROMPT_VERSION,
                                "test": test["id"], "language": lang, "prompt_variant": variant["id"],
                                "model": model, "model_options": model_cfg["options"], "harness": harness,
                                "thinking": model_cfg["options"].get("think", None),
                                "response_raw": f"ERROR: {e}", "code_extracted": False, "code": "",
                                "ms": 0, "eval_count": 0, "tok_per_sec": 0, "ran": False,
                                "exit_code": None, "run_ms": 0, "stdout": "", "stderr": str(e),
                                "checks": {}, "passed": False,
                            }
                        status_str = _c("✅ PASS", _GREEN) if record.get("passed") else _c("❌ FAIL", _RED)
                        print(f"    {status_str}  exit={record.get('exit_code')}  {record['ms']}ms gen  "
                              f"{record.get('run_ms', 0)}ms run  {record.get('tok_per_sec', 0)} tok/s", flush=True)
                        run_records.append(record)

                all_records.extend(run_records)
                append_results(run_records, harness, model_safe)
                print(f"\n  Appended {len(run_records)} records to {results_file(harness, model_safe)}")
                passed_n = sum(1 for r in run_records if r.get("passed"))
                print(f"  Score: {_c(f'✅ {passed_n}', _GREEN)}  {_c(f'❌ {len(run_records) - passed_n}', _RED)}"
                      f"  ({passed_n}/{len(run_records)})")

            if harness == "ollama":
                try:
                    unload(model)
                except Exception as e:
                    print(_c(f"  WARNING: failed to unload: {e}", _YELLOW), file=sys.stderr)
    finally:
        local_backend.stop()
        if apple_backend_obj is not None:
            apple_backend_obj.stop()

    print_tables(all_records)
```

- [ ] **Step 5: Add tok/s to the per-model summary in `print_tables`**

In `print_tables` (lines 543-586), after the per-test loop that prints rows for a model, add an aggregate line. Immediately before the `for test in tests:` loop closes for each `(model, thinking)`, compute and print:

```python
        model_recs = [
            r for r in records
            if r["model"] == model and r["thinking"] == thinking
        ]
        gen_ms = [r["ms"] for r in model_recs if isinstance(r.get("ms"), (int, float)) and r["ms"]]
        toks = [r["tok_per_sec"] for r in model_recs if r.get("tok_per_sec")]
        if gen_ms:
            mean_s = sum(gen_ms) / len(gen_ms) / 1000
            mean_tok = sum(toks) / len(toks) if toks else 0.0
            print(f"{'':<20}mean ~{mean_s:.1f}s/task · {mean_tok:.0f} tok/s")
```

- [ ] **Step 6: Verify — settings load + dry structure**

Run: `cd codegen && python3 -c "import settings; s = settings.load_settings(); print([m['name'] for m in s.local_models()])"`
Expected: `['qwen3.8:27b', 'qwen3.8:27b']` (both enabled entries).

- [ ] **Step 7: Verify — one real Ollama run**

Precondition: `ollama serve` running, `qwen3.8:27b` pulled.
Run: `cd codegen && python3 benchmark.py 1`
Expected: run completes; new records land in `results/v002/ollama/qwen3.8_27b_nothink/results.jsonl` and `.../qwen3.8_27b_think/results.jsonl`; each record has `"harness": "ollama"`; artifacts under `results/v002/ollama/qwen3.8_27b_*/{run_id}/{lang}/{test}/`; the summary prints a `mean ~Xs/task · N tok/s` line per model.

Verify field: `jq -r '.harness' results/v002/ollama/qwen3.8_27b_think/results.jsonl | sort -u` → `ollama`

- [ ] **Step 8: Verify — LM Studio path errors cleanly when down**

Run (LM Studio NOT running): `cd codegen && python3 benchmark.py 1 --harness lmstudio`
Expected: exits 1 with `ERROR: LM Studio not reachable at http://localhost:1234 — start the LM Studio server and load a model`. No partial `results/v002/lmstudio/` tree created.

- [ ] **Step 9: Verify — Apple path still works**

Precondition: `apfel` installed, macOS 26+.
Run: `cd codegen && python3 benchmark.py 1 --apple`
Expected: apfel starts, `apple-foundationmodel` runs, results at `results/v002/apple/apple-foundationmodel/`, apfel stops at the end.

- [ ] **Step 10: Lint + commit**

```bash
cd codegen && bash lint.sh
git add codegen/benchmark.py
git commit -m "feat(codegen): drive benchmark.py from settings + backends, harness results layout"
```

---

## Task 4: Wire `run_claude_test.py` to settings + `anthropic/` path

**Files:**
- Modify: `codegen/run_claude_test.py` — `MODELS` dict (33-37), `--model` default (169), model resolution (176-182), `results_dir` (183-187), artifact path (215), record dict (223-246), `results_jsonl` path (250).

**Interfaces:**
- Consumes: `settings.load_settings`, `settings.Settings` (Task 1).
- Produces: JSONL records with `"harness": "anthropic"`. Results at `results/v{PROMPT_VERSION:03d}/anthropic/{friendly}/results.jsonl`.

- [ ] **Step 1: Replace the `MODELS` dict with a settings load**

Delete lines 29-37 (the comment block + `MODELS: dict[str, str] = {...}`). Replace with:

```python
import settings as settings_mod

BASE_DIR = Path(__file__).parent


def _anthropic_models() -> dict[str, str]:
    """alias -> model_id, from settings.toml [[anthropic_models]] (enabled only)."""
    s = settings_mod.load_settings(BASE_DIR / "settings.toml")
    return {a["alias"]: a["model_id"] for a in s.anthropic_models()}
```

(Note: `BASE_DIR` is currently defined at line 56 — move it up as shown and delete the later definition, or keep the later one and drop the duplicate here. One definition only.)

- [ ] **Step 2: Update `main()` model resolution and `--model` default**

Replace lines 169 and 176-182. New `--model` argument:

```python
    parser.add_argument("--model", default=None,
                        help="Anthropic alias (haiku/sonnet/opus) or full model id; "
                             "default: first enabled alias in settings.toml")
```

New resolution block (replacing the `_reverse` logic):

```python
    models = _anthropic_models()
    reverse = {v: k for k, v in models.items()}
    requested = args.model or (next(iter(models), "haiku"))
    if requested in models:
        friendly, model = requested, models[requested]
    elif requested in reverse:
        friendly, model = reverse[requested], requested
    else:
        friendly, model = requested, requested  # unknown — use as-is
```

- [ ] **Step 3: Add `anthropic/` to results + artifact paths**

Change `results_dir` (183-187):

```python
    results_dir = (
        Path(args.results_dir)
        if args.results_dir
        else BASE_DIR / "results" / f"v{PROMPT_VERSION:03d}" / "anthropic"
    )
```

The existing `results_dir / friendly / run_id / ...` (line 215) and `results_dir / friendly / "results.jsonl"` (line 250) then resolve to `.../anthropic/{friendly}/...` automatically — no further path edits.

- [ ] **Step 4: Add `harness` to the record**

In the `record` dict (line ~223), add after `"model_options":  {},`:

```python
        "harness":        "anthropic",
```

- [ ] **Step 5: Verify — resolution + default**

Run: `cd codegen && python3 -c "import run_claude_test as r; print(r._anthropic_models())"`
Expected: `{'haiku': 'claude-haiku-4-5-20251001'}` (only `haiku` enabled in shipped settings).

- [ ] **Step 6: Verify — one scored cell writes to `anthropic/`**

Precondition: a solution file, e.g. `echo 'print("hi")' > /tmp/t.py`.
Run: `cd codegen && python3 run_claude_test.py python 003_fibonacci /tmp/t.py 20260827_000000 --model haiku`
Expected: writes `results/v002/anthropic/haiku/results.jsonl` with a record whose `"harness"` is `"anthropic"`; artifacts under `results/v002/anthropic/haiku/20260827_000000/python/003_fibonacci/`. (The solution will FAIL verification — that is fine; we are checking paths and the record, not the score.)

Clean up the throwaway: `rm -r results/v002/anthropic/haiku/20260827_000000 && sed -i '' '/20260827_000000/d' results/v002/anthropic/haiku/results.jsonl` (or delete the file if it was newly created).

- [ ] **Step 7: Lint + commit**

```bash
cd codegen && bash lint.sh
git add codegen/run_claude_test.py
git commit -m "feat(codegen): run_claude_test.py reads models from settings, writes under anthropic/"
```

---

## Task 5: Documentation updates

**Files:**
- Modify: `codegen/CLAUDE.md`
- Modify: `codegen/results/findings_instructions.md`
- Modify: `codegen/README.md`

- [ ] **Step 1: `codegen/CLAUDE.md` — "Structure" section**

Update the results-path bullets and add the new files. Replace the current structure bullets with:

```markdown
- `settings.toml` — model list, harness config, timeouts, languages. Edit this to change a run.
- `settings.py` — loader/validator for `settings.toml`
- `backends.py` — `OllamaBackend`, `LMStudioBackend`, `AppleBackend`
- `benchmark.py` — main test runner (Ollama / LM Studio / apfel)
- `run_claude_test.py` — Claude API test runner (harness = `anthropic`)
- `migrate_results.py` — one-time migration of old `results/vNNN/{model}/` dirs into the harness layout
- `results/FINDINGS_v{NNN}.md` — human-readable analysis and conclusions
- `results/v{NNN}/{harness}/{model}/results.jsonl` — results per harness+model; `NNN` = `PROMPT_VERSION`
- `results/v{NNN}/{harness}/{model}/{timestamp}/{lang}/{test}/` — per-run artifacts
```

- [ ] **Step 2: `codegen/CLAUDE.md` — "Running" section**

Replace the running block with:

```markdown
Model list, harness, and timeouts live in `settings.toml`. Set `enabled = true`
on the models you want; pick a harness with `--harness` (default from
`[harness].default`).

    python3 benchmark.py                 # enabled models, default harness (ollama)
    python3 benchmark.py 3               # 3 runs
    python3 benchmark.py --harness lmstudio   # run the same models against LM Studio
    python3 benchmark.py --apple         # also run the Apple on-device model

Ollama or LM Studio must be running. With `--apple`, apfel is started
automatically on port 11435 and stopped when the run completes (unless it was
already running).
```

- [ ] **Step 3: `codegen/CLAUDE.md` — "Analyzing Results" section**

Every jq glob `results/v002/*/results.jsonl` becomes `results/v002/*/*/results.jsonl` (harness/model). Add harness-scoped examples and a tok/s query. Replace the code block after "Use targeted jq queries" with:

````markdown
```sh
# Pass rate per model across a version (all harnesses)
jq -s 'group_by(.model) | map({model: .[0].model, passed: map(select(.passed))|length, total: length})' results/v002/*/*/results.jsonl

# One harness only
jq -s 'group_by(.model) | map({model: .[0].model, passed: map(select(.passed))|length, total: length})' results/v002/ollama/*/results.jsonl

# Same model, ollama vs lmstudio
jq -s 'group_by(.harness) | map({harness: .[0].harness, passed: map(select(.passed))|length, total: length})' \
  results/v002/ollama/qwen3.8_27b_nothink/results.jsonl results/v002/lmstudio/qwen3.8_27b_nothink/results.jsonl

# Mean tok/s per model
jq -s 'group_by(.model + (.thinking|tostring)) | map({model: .[0].model, thinking: .[0].thinking, tok_per_sec: (map(.tok_per_sec) | add / length)})' results/v002/*/*/results.jsonl

# Anthropic results
jq 'select(.passed == false) | {test, language}' results/v002/anthropic/haiku/results.jsonl
```
````

Add a sentence noting the new `harness` field: "Every record carries a `harness` field (`ollama` / `lmstudio` / `apple` / `anthropic`)."

- [ ] **Step 4: `codegen/results/findings_instructions.md` — harness subdirs + tok/s**

In "Before You Start" step 1, change `ls results/v002/` guidance to `ls results/v002/*/` and add: "Note which harnesses ran (`ollama`, `lmstudio`, `apple`, `anthropic`) and whether any model ran under more than one."

In "Report Structure", add a new subsection after "Quantization":

```markdown
### Harness Comparison (Ollama vs LM Studio) *(if a model ran under both)*
For each model run under both harnesses: score parity (should be near-equal — a
large gap is a harness bug, not a model finding), speed delta, and tok/s delta.
Note that tok/s is sourced differently per harness (Ollama `eval_duration`,
LM Studio `stats.tokens_per_second` or wall-clock fallback, Apple wall-clock)
and is not perfectly comparable.
```

In "Speed vs. Quality", change "Table of all timed models: Model | Mode | Avg time/task | tok/s | Score | Pass Rate." to make tok/s mandatory and add a Harness column: "Model | Harness | Mode | Avg time/task | tok/s | Score | Pass Rate. tok/s is required for every row."

- [ ] **Step 5: `codegen/README.md` — settings + harness + LM Studio setup**

Add a "Configuration" section describing `settings.toml` (models with `enabled`, `[harness]`, `[defaults]`), document `--harness ollama|lmstudio`, and add an "LM Studio" subsection: "Start the LM Studio server (Developer tab → Start Server, default port 1234) and load the model you want to test. There is no autostart — the benchmark errors if the server is unreachable. LM Studio's model key may differ from the Ollama tag; set `lmstudio_model` on the model entry in `settings.toml` if so."

- [ ] **Step 6: Commit**

```bash
git add codegen/CLAUDE.md codegen/results/findings_instructions.md codegen/README.md
git commit -m "docs(codegen): settings.toml, --harness, per-harness results layout"
```

---

## Task 6: `migrate_results.py`

**Files:**
- Create: `codegen/migrate_results.py`
- Create: `codegen/tests/test_migrate_results.py`

**Interfaces:**
- Consumes: `settings.load_settings` (Task 1) for the anthropic alias list.
- Produces: a CLI. `classify_dir(name: str, anthropic_aliases: set[str]) -> str | None` returns the target harness (`ollama`/`lmstudio`/`apple`/`anthropic`) or `None` if the dir is already a harness dir and should be skipped.

- [ ] **Step 1: Write the failing tests — `codegen/tests/test_migrate_results.py`**

```python
import unittest

import migrate_results as mr

ALIASES = {"haiku", "opus", "sonnet"}


class ClassifyTests(unittest.TestCase):
    def test_apple(self) -> None:
        self.assertEqual(mr.classify_dir("apple-foundationmodel", ALIASES), "apple")

    def test_anthropic_alias(self) -> None:
        self.assertEqual(mr.classify_dir("haiku", ALIASES), "anthropic")
        self.assertEqual(mr.classify_dir("sonnet", ALIASES), "anthropic")

    def test_local_model(self) -> None:
        self.assertEqual(mr.classify_dir("qwen3.5_27b_think", ALIASES), "ollama")
        self.assertEqual(mr.classify_dir("gemma4_31b_nothink", ALIASES), "ollama")

    def test_already_harness_dir_skipped(self) -> None:
        for h in ("ollama", "lmstudio", "apple", "anthropic"):
            self.assertIsNone(mr.classify_dir(h, ALIASES))


if __name__ == "__main__":
    unittest.main()
```

- [ ] **Step 2: Run tests to verify they fail**

Run: `cd codegen && python3 -m unittest discover -s tests -t . -p 'test_migrate_results.py' -v`
Expected: FAIL — `ModuleNotFoundError: No module named 'migrate_results'`

- [ ] **Step 3: Write `codegen/migrate_results.py`**

```python
"""One-time migration: results/vNNN/{model}/ -> results/vNNN/{harness}/{model}/.

Dry-run by default. Pass --apply to perform `git mv`. Run once, by hand, on a
clean tree.
"""

from __future__ import annotations

import argparse
import subprocess
import sys
from pathlib import Path

import settings as settings_mod

HARNESS_DIRS = {"ollama", "lmstudio", "apple", "anthropic"}
_HISTORICAL_ANTHROPIC = {"haiku", "opus", "sonnet"}


def classify_dir(name: str, anthropic_aliases: set[str]) -> str | None:
    if name in HARNESS_DIRS:
        return None
    if name == "apple-foundationmodel":
        return "apple"
    if name in anthropic_aliases or name in _HISTORICAL_ANTHROPIC:
        return "anthropic"
    return "ollama"


def _git(args: list[str]) -> None:
    subprocess.run(["git", *args], check=True)


def migrate(version_dir: Path, aliases: set[str], apply: bool) -> None:
    moves: list[tuple[Path, Path]] = []
    for child in sorted(version_dir.iterdir()):
        if not child.is_dir():
            continue
        target_harness = classify_dir(child.name, aliases)
        if target_harness is None:
            continue
        dest = version_dir / target_harness / child.name
        moves.append((child, dest))

    if not moves:
        print("Nothing to migrate.")
        return

    for src, dest in moves:
        print(f"{'MOVE' if apply else 'DRY '} {src.relative_to(version_dir.parent)}  ->  {dest.relative_to(version_dir.parent)}")
        if apply:
            dest.parent.mkdir(parents=True, exist_ok=True)
            _git(["mv", str(src), str(dest)])

    if not apply:
        print(f"\n{len(moves)} dirs would move. Re-run with --apply.")


def main() -> None:
    parser = argparse.ArgumentParser(description="Migrate results into the harness layout")
    parser.add_argument("version", nargs="?", default=None, help="e.g. v002 (default: all vNNN dirs)")
    parser.add_argument("--apply", action="store_true", help="perform the moves (default: dry run)")
    args = parser.parse_args()

    base = Path(__file__).parent
    aliases = {a["alias"] for a in settings_mod.load_settings(base / "settings.toml").anthropic_models()}

    results = base / "results"
    version_dirs = (
        [results / args.version] if args.version
        else sorted(d for d in results.iterdir() if d.is_dir() and d.name.startswith("v"))
    )
    for vd in version_dirs:
        if not vd.is_dir():
            print(f"skip: {vd} not found", file=sys.stderr)
            continue
        print(f"\n=== {vd.name} ===")
        migrate(vd, aliases, args.apply)


if __name__ == "__main__":
    main()
```

- [ ] **Step 4: Run tests to verify they pass**

Run: `cd codegen && python3 -m unittest discover -s tests -t . -p 'test_migrate_results.py' -v`
Expected: PASS

- [ ] **Step 5: Full test sweep + lint**

Run: `cd codegen && python3 -m unittest discover -s tests -t . -p 'test_*.py' -v && bash lint.sh`
Expected: all tests pass, lint clean.

- [ ] **Step 6: Commit the script**

```bash
git add codegen/migrate_results.py codegen/tests/test_migrate_results.py
git commit -m "feat(codegen): add migrate_results.py for the harness layout"
```

- [ ] **Step 7: Dry-run the migration and review**

Run: `cd codegen && python3 migrate_results.py v002`
Expected: prints `DRY ... -> ...` lines. Confirm: `apple-foundationmodel` → `apple/`, `haiku`/`opus`/`sonnet` → `anthropic/`, every `qwen*`/`gemma*` dir → `ollama/`. Nothing already under a harness dir.

- [ ] **Step 8: Apply the migration on a clean tree**

Precondition: `git status` clean (all prior tasks committed).
Run: `cd codegen && python3 migrate_results.py v002 --apply`
Then: `git status` — should show only renames (`R`).
Then: `jq -s 'length' results/v002/ollama/*/results.jsonl | head` and a couple of the updated jq queries from `CLAUDE.md` to confirm they resolve.

- [ ] **Step 9: Commit the migration**

```bash
git add -A codegen/results/v002
git commit -m "chore(codegen): migrate v002 results into harness layout"
```

- [ ] **Step 10: Regenerate the findings doc header path references (optional, if stale)**

`FINDINGS_v002.md` prose does not hard-code result paths, so no change is needed. If a future `jq` snippet is added to it, use the `results/v002/{harness}/{model}/` form.

---

## Self-Review

**Spec coverage:**
- settings.toml + loader → Task 1 ✅
- LM Studio backend (`/api/v0` preferred, `/v1` fallback, clear error when down) → Task 2 (`LMStudioBackend`), Task 3 Step 8 ✅
- Ollama backend (server-side timing preserved) → Task 2 (`OllamaBackend`) ✅
- Apple adapter over root `apfel_backend`, no root changes → Task 2 (`AppleBackend`) ✅
- `--harness` flag, one local harness per run, `--apple` still separate → Task 3 Step 4 ✅
- `enabled` flag model selection → Task 1 (`local_models()` filters) ✅
- Results at `results/v{NNN}/{harness}/{model}/` + `harness` record field → Task 3 Steps 2-3, Task 4 Steps 3-4 ✅
- run_claude_test.py folded into `anthropic/`, model list from settings, `--model` default = first enabled alias → Task 4 ✅
- Summary shows tok/s → Task 3 Step 5 ✅
- Doc updates (CLAUDE.md analysis paths, findings_instructions harness comparison + mandatory tok/s, README LM Studio setup) → Task 5 ✅
- `migrate_results.py` dry-run default, `git mv`, classification rules → Task 6 ✅
- `PROMPT_VERSION` unchanged, additive JSONL fields → Global Constraints, respected throughout ✅
- Non-goals (no shell/, no apfel_backend.py rewrite, no profiles, no autostart LM Studio, no mixed local harnesses, no v003) → respected ✅

**Gaps / deviations from spec:**
- Spec mentioned `settings.py` accessor `harness_apple()` returning a config object; simplified to `apple_autostart()` + `harness_base_url("apple")` since only those two values are used.
- Spec's `migrate_results.py --stamp-harness` (inject `harness` into old JSONL records) is **not** implemented in this plan. Old records simply lack the field; jq queries tolerate that. If stamping is wanted, add a follow-up task. Flagged for the executor.
- Spec's `--force` guard on a dirty results tree is replaced by Task 6 Step 8's precondition ("clean tree") plus `git mv` failing loudly on conflicts. Simpler; same safety.
- `AppleBackend.generate` passes empty `run_id`/`timestamp` to `apfel_backend.run_prompt` because `run_one` builds the record itself; the apfel return values for those fields are discarded (as they were before this refactor).

**Placeholder scan:** none — all steps carry real code or exact commands.

**Type consistency:** `ModelEntry` (settings) vs `ModelConfig` (benchmark) are deliberately distinct — Task 3 Step 4 converts one to the other explicitly. `GenResult` shape is identical in `backends.py` and consumed unchanged in `run_one`. `classify_dir` signature matches between Task 6 Step 1 (test) and Step 3 (impl).
