#!/usr/bin/env python3
"""
Code Generation Benchmark

Tests local Ollama models for code generation quality.
Results are appended to results/vNNN/{model}/results.jsonl — run as many times as you like.

Usage:
  python3 benchmark.py              # single run
  python3 benchmark.py 3            # 3 runs, results aggregated at end
  python3 benchmark.py --apple      # include Apple on-device model
  python3 benchmark.py 3 --apple    # 3 runs with Apple model
"""

import argparse
import datetime
import importlib.util
import json
import os
import re
import shlex
import shutil
import subprocess
import sys
import tempfile
import time
import types
import urllib.error
import urllib.request
from pathlib import Path
from typing import Any, NotRequired, TypedDict, cast

sys.path.insert(0, str(Path(__file__).parent.parent))
import apfel_backend

# ---------------------------------------------------------------------------
# Terminal colors
# ---------------------------------------------------------------------------

_COLOR = sys.stdout.isatty() and not os.environ.get("NO_COLOR")

_RESET   = "\033[0m"
_WHITE   = "\033[97m"
_BG_GRAY = "\033[100m"
_GREEN   = "\033[32m"
_RED     = "\033[31m"
_YELLOW  = "\033[33m"
_BLUE    = "\033[34m"
_CYAN    = "\033[36m"
_MAGENTA = "\033[35m"
_PURPLE  = "\033[38;5;141m"


def _c(text: str, code: str) -> str:
    return f"{code}{text}{_RESET}" if _COLOR else text


OLLAMA_BASE = "http://localhost:11434"
RESULTS_DIR = Path("results")

PROMPT_VERSION = 2

EXEC_TIMEOUT = 60    # seconds; can be overridden per model via exec_timeout in ModelConfig
INFER_TIMEOUT = 120  # seconds; can be overridden per model via infer_timeout in ModelConfig

# Populated by check_runtimes() with extra env vars needed to run C# scripts.
_csharp_env: dict[str, str] | None = None


# ---------------------------------------------------------------------------
# Type definitions
# ---------------------------------------------------------------------------

class ModelConfig(TypedDict):
    name: str
    options: dict[str, Any]
    backend: NotRequired[str]
    exec_timeout: NotRequired[int]   # seconds; overrides EXEC_TIMEOUT for code execution
    infer_timeout: NotRequired[int]  # seconds; overrides INFER_TIMEOUT for API inference call


class PromptVariant(TypedDict):
    id: str        # "A" or "B"
    language: str  # "python", "typescript", "go", "csharp"
    system_prompt: str  # fully resolved (no {placeholders})


class TestCase(TypedDict):
    id: str                          # "001_csv_to_json"
    prompt: str                      # raw content of prompt.md
    input_dir: NotRequired[Path]     # path to test's input/ dir
    verify_mod: types.ModuleType     # imported verify.py module
    test_args: NotRequired[list[str]]  # extra CLI args passed to the generated program


# ---------------------------------------------------------------------------
# Models
# ---------------------------------------------------------------------------

MODELS: list[ModelConfig] = [
    {"name": "gemma4:26b",                   "options": {"think": False}},
    {"name": "gemma4:26b",                   "options": {"think": True},  "infer_timeout": 600},
    {"name": "gemma4:31b",                   "options": {"think": False}},
    {"name": "gemma4:31b",                   "options": {"think": True},  "infer_timeout": 600},
    {"name": "qwen3.5:35b-a3b-coding-nvfp4", "options": {"think": False}},
    {"name": "qwen3.5:35b-a3b-coding-nvfp4", "options": {"think": True},  "infer_timeout": 600},
    {"name": "qwen3.5:27b",                  "options": {"think": False}},
    {"name": "qwen3.5:27b",                  "options": {"think": True},  "infer_timeout": 600},
    {"name": "qwen3.5:27b-nvfp4",            "options": {"think": False}},
    {"name": "qwen3.5:27b-nvfp4",            "options": {"think": True},  "infer_timeout": 600},
    {"name": "qwen3.5:4b",                   "options": {"think": False}},
    {"name": "qwen3.5:4b",                   "options": {"think": True},  "infer_timeout": 600},
    {"name": "qwen3.5:4b-nvfp4",             "options": {"think": False}},
    {"name": "qwen3.5:4b-nvfp4",             "options": {"think": True},  "infer_timeout": 600},
    {"name": "qwen2.5-coder:7b",             "options": {}},
    {"name": "qwen3-coder:30b",              "options": {}},
]

# ---------------------------------------------------------------------------
# Language config
# ---------------------------------------------------------------------------

LANGUAGES = ["python", "typescript", "go", "csharp"]

LANG_NAME: dict[str, str] = {
    "python": "Python",
    "typescript": "TypeScript",
    "go": "Go",
    "csharp": "C#",
}
LANG_EXT: dict[str, str] = {
    "python": ".py",
    "typescript": ".ts",
    "go": ".go",
    "csharp": ".csx",
}
LANG_RUN: dict[str, str] = {
    "python": "python3",
    "typescript": "tsx",
    "go": "go run",
    "csharp": "dotnet script",
}
LANG_NOTE: dict[str, str] = {
    "python": "",
    "typescript": "Use Node.js-compatible TypeScript. The code will be run with tsx. Use only Node.js built-in modules (fs, path, etc.) — no npm packages are available. Read command-line arguments with process.argv (e.g. process.argv[2]), not readline or stdin.",
    "go": "Write a complete Go program with package main and a main() function.",
    "csharp": "Write a .csx script (dotnet-script). No class or Main method needed. Top-level statements only. Command-line arguments are available as `Args` (an IList<string>), not `args`. No NuGet packages are available — use only built-in .NET assemblies. For CSV: read with `File.ReadAllLines()` and parse with `string.Split(',')`. For JSON: use `using System.Text.Json;` and `using System.Text.Json.Nodes;`.",
}
LANG_COLOR: dict[str, str] = {
    "python":     _YELLOW,
    "typescript": _BLUE,
    "go":         _CYAN,
    "csharp":     _MAGENTA,
}

# Extensions used for language source files — excluded from data-file inlining
# in build_messages so {source_code} substitutions aren't double-appended.
_LANG_EXTENSIONS: frozenset[str] = frozenset({".py", ".ts", ".go", ".cs", ".csx"})


# ---------------------------------------------------------------------------
# Prompt variants
# ---------------------------------------------------------------------------

# Prompt history (for reference when reviewing older results):
# v1-v2 "A" — strict: "You are a {lang_name} developer. Return ONLY the raw {lang_name} code.
#               No markdown fences, no backticks, no explanation, no commentary.
#               The output must be valid {lang_name} source code that can be run directly. {lang_note}"
# v1-v2 "B" — natural: "You are a helpful {lang_name} developer. {lang_note}"
_VARIANT_TEMPLATES = {
    # v3+ "C" — natural + explicit "write a program" + stdout hint
    "C": (
        "You are a helpful {lang_name} developer.\n"
        "Write a {lang_name} program that solves the task described by the user.\n"
        "Write all output to stdout.\n"
        "{lang_note}"
    ).strip(),
}

PROMPT_VARIANTS: list[PromptVariant] = [
    PromptVariant(
        id=vid,
        language=lang,
        system_prompt=tmpl.replace("{lang_name}", LANG_NAME[lang])
                          .replace("{lang_note}", LANG_NOTE[lang])
                          .strip(),
    )
    for vid, tmpl in _VARIANT_TEMPLATES.items()
    for lang in LANGUAGES
]


# ---------------------------------------------------------------------------
# Test args & discovery
# ---------------------------------------------------------------------------

TEST_ARGS: dict[str, list[str]] = {
    "003_fibonacci":     ["100"],
    "008_prime_numbers": ["50"],
}


def load_verify(test_dir: Path) -> types.ModuleType:
    spec = importlib.util.spec_from_file_location("verify", test_dir / "grading" / "verify.py")
    mod = importlib.util.module_from_spec(spec)  # type: ignore[arg-type]
    spec.loader.exec_module(mod)  # type: ignore[union-attr]
    return mod


def load_tests(tests_dir: Path) -> list[TestCase]:
    cases: list[TestCase] = []
    for test_dir in sorted(tests_dir.iterdir()):
        if not test_dir.is_dir() or test_dir.name.startswith('_'):
            continue
        test_id = test_dir.name
        case: TestCase = {
            "id":         test_id,
            "prompt":     (test_dir / "test" / "prompt.md").read_text(),
            "verify_mod": load_verify(test_dir),
        }
        input_dir = test_dir / "test" / "input"
        if input_dir.is_dir():
            case["input_dir"] = input_dir
        if test_id in TEST_ARGS:
            case["test_args"] = TEST_ARGS[test_id]
        cases.append(case)
    return cases


_test_cases: list[TestCase] = []  # populated in main()


# ---------------------------------------------------------------------------
# API
# ---------------------------------------------------------------------------

def ollama_post(path: str, payload: dict[str, Any], timeout: int = 120) -> dict[str, Any]:
    data = json.dumps(payload).encode()
    req = urllib.request.Request(
        f"{OLLAMA_BASE}{path}",
        data=data,
        headers={"Content-Type": "application/json"},
        method="POST",
    )
    try:
        with urllib.request.urlopen(req, timeout=timeout) as resp:
            return cast(dict[str, Any], json.loads(resp.read()))
    except urllib.error.HTTPError as e:
        raise RuntimeError(_http_error_message(e)) from e


def _http_error_message(e: urllib.error.HTTPError) -> str:
    try:
        body = cast(dict[str, Any], json.loads(e.read().decode(errors="replace")))
        err = body.get("error", "")
        if isinstance(err, dict):
            err_dict = cast(dict[str, Any], err)
            return f"HTTP {e.code}: {str(err_dict.get('message') or err_dict)}"
        return f"HTTP {e.code}: {str(err) or e.reason}"
    except Exception:
        return f"HTTP {e.code}: {e.reason}"


def preload(model: str, messages: list[dict[str, str]], options: dict[str, Any], infer_timeout: int = INFER_TIMEOUT) -> None:
    print(_c(f"  preloading {model}...", _CYAN), flush=True)
    ollama_post("/api/chat", {"model": model, "messages": [], "keep_alive": -1}, timeout=60)
    print(_c(f"  warming up {model}...", _CYAN), flush=True)
    warmup_options: dict[str, Any] = {**options, "think": False} if options.get("think") is True else options
    ollama_post("/api/chat", {
        "model": model,
        "messages": messages,
        "stream": False,
        **warmup_options,
    }, timeout=infer_timeout)


def unload(model: str) -> None:
    print(_c(f"  unloading {model}...", _CYAN), flush=True)
    ollama_post("/api/chat", {"model": model, "messages": [], "keep_alive": 0}, timeout=15)


# ---------------------------------------------------------------------------
# Runtime detection
# ---------------------------------------------------------------------------

def _run_check(cmd: str, env: dict[str, str] | None = None) -> bool:
    try:
        result = subprocess.run(
            cmd, shell=True, capture_output=True, timeout=10, env=env
        )
        return result.returncode == 0
    except Exception:
        return False


def _detect_csharp_env() -> None:
    """Detect env overrides needed for dotnet-script if it can't find the runtime on its own.

    Two common issues on Homebrew setups:
    - ~/.dotnet/tools is not in PATH, so `dotnet script` can't find dotnet-script
    - DOTNET_ROOT isn't set, so dotnet-script can't find the runtime
    """
    global _csharp_env
    import os

    if _run_check("dotnet script --version"):
        return  # already works

    # Build a candidate env: add ~/.dotnet/tools to PATH and detect DOTNET_ROOT from `dotnet --info`.
    tools_dir = str(Path.home() / ".dotnet" / "tools")
    current_path = os.environ.get("PATH", "")
    patched_path = f"{tools_dir}:{current_path}" if tools_dir not in current_path else current_path

    candidate_env: dict[str, str] = {**os.environ, "PATH": patched_path}

    try:
        result = subprocess.run(
            "dotnet --info", shell=True, capture_output=True, text=True, timeout=10
        )
        for line in result.stdout.splitlines():
            if "Base Path:" in line:
                # Base Path: /opt/homebrew/Cellar/dotnet/X.Y.Z/libexec/sdk/X.Y.Z/
                # DOTNET_ROOT is two levels up (past "sdk/X.Y.Z")
                sdk_path = line.split("Base Path:")[-1].strip()
                dotnet_root = str(Path(sdk_path).parent.parent)
                candidate_env["DOTNET_ROOT"] = dotnet_root
                break
    except Exception:
        pass

    if _run_check("dotnet script --version", candidate_env):
        _csharp_env = candidate_env
        print(
            "  dotnet-script: PATH and DOTNET_ROOT patched for Homebrew install",
            flush=True,
        )
    else:
        print("  WARNING: dotnet-script runtime not found; C# runs may fail.", file=sys.stderr)


def _run_install(cmd: str) -> bool:
    print(f"    Running: {cmd}", flush=True)
    result = subprocess.run(cmd, shell=True)
    return result.returncode == 0


def check_runtimes(languages: list[str]) -> None:
    print("Checking runtimes...", flush=True)
    for lang in languages:
        if lang == "python":
            if not _run_check("python3 --version"):
                print("ERROR: python3 not found (required)", file=sys.stderr)
                sys.exit(1)
        elif lang == "typescript":
            if _run_check("tsx --version"):
                pass
            elif _run_check("node --version"):
                print("  tsx not found but node present — installing tsx...")
                if not _confirm_install("npm install -g tsx"):
                    sys.exit(1)
            else:
                if _run_check("bash -c 'fnm --version'"):
                    cmd = "bash -c 'fnm install --lts && fnm use lts-latest && npm install -g tsx'"
                elif _run_check("bash -c 'source \"$NVM_DIR/nvm.sh\" 2>/dev/null && nvm --version'"):
                    cmd = "bash -c 'source \"$NVM_DIR/nvm.sh\" && nvm install --lts && nvm use --lts && npm install -g tsx'"
                else:
                    cmd = "brew install node && npm install -g tsx"
                print(f"  node not found — will install via: {cmd}")
                if not _confirm_install(cmd):
                    sys.exit(1)
        elif lang == "go":
            if not _run_check("go version"):
                print("  go not found — installing via brew...")
                if not _confirm_install("brew install go"):
                    sys.exit(1)
        elif lang == "csharp":
            if not _run_check("dotnet --version"):
                print("  dotnet not found — installing via brew...")
                if not _confirm_install("brew install dotnet"):
                    sys.exit(1)
            if not _run_check("dotnet tool list --global | grep -q dotnet-script"):
                print("  dotnet-script not found — installing...")
                if not _confirm_install("dotnet tool install -g dotnet-script"):
                    sys.exit(1)
            _detect_csharp_env()
    print(_c("  All runtimes OK.", _GREEN), flush=True)


def _confirm_install(cmd: str) -> bool:
    answer = input(f"  Install? [{cmd}] (y/N): ").strip().lower()
    if answer != "y":
        print("  Aborted.", file=sys.stderr)
        return False
    return _run_install(cmd)


# ---------------------------------------------------------------------------
# Prompt construction
# ---------------------------------------------------------------------------

def build_messages(
    variant: PromptVariant,
    test: TestCase,
    language: str,
) -> list[dict[str, str]]:
    task = test["prompt"]
    task = task.replace("{ext}", LANG_EXT[language])
    task = task.replace("{language}", LANG_NAME[language])
    if "{source_code}" in task and "input_dir" in test:
        ext = LANG_EXT[language]
        for suffix in [ext, ".cs"]:
            candidates = list(test["input_dir"].glob(f"*{suffix}"))
            if candidates:
                task = task.replace("{source_code}", candidates[0].read_text())
                break
    # Append data input files inline so models see actual content rather than
    # having to infer schema from the prompt description alone. Language source
    # files (.py, .ts, .go, .cs, .csx) are skipped — already inlined via
    # {source_code} for tests that use it (005, 006).
    if "input_dir" in test:
        for f in sorted(test["input_dir"].iterdir()):
            if f.is_file() and f.suffix not in _LANG_EXTENSIONS:
                task += f"\n\n--- input/{f.name} ---\n{f.read_text()}"
    return [
        {"role": "system", "content": variant["system_prompt"]},
        {"role": "user",   "content": task},
    ]


# ---------------------------------------------------------------------------
# Code extraction
# ---------------------------------------------------------------------------

_FENCE_RE = re.compile(
    r"```[ \t]*(?:[a-zA-Z0-9_+\-]*)[ \t]*\n(.*?)```",
    re.DOTALL,
)


def extract_code(response: str) -> tuple[str, bool]:
    """Return (code, code_extracted). Strips markdown fences if present."""
    match = _FENCE_RE.search(response)
    if match:
        return match.group(1).strip(), True
    return response.strip(), False


# ---------------------------------------------------------------------------
# Code execution
# ---------------------------------------------------------------------------

def run_code(
    code: str,
    language: str,
    input_dir: Path | None,
    test_args: list[str],
    timeout: int = EXEC_TIMEOUT,
) -> dict[str, Any]:
    ext = LANG_EXT[language]
    run_cmd = shlex.split(LANG_RUN[language])
    solution_name = f"solution{ext}"

    with tempfile.TemporaryDirectory() as tmpdir:
        tmp = Path(tmpdir)
        solution_path = tmp / solution_name
        try:
            solution_path.write_text(code)
        except Exception as e:
            return {"ran": False, "exit_code": None, "stdout": "", "stderr": str(e), "run_ms": 0}

        if input_dir is not None:
            shutil.copytree(input_dir, tmp / "input")

        cmd = run_cmd + [solution_name] + test_args
        run_env = _csharp_env if language == "csharp" and _csharp_env is not None else None
        start = time.monotonic()
        try:
            proc = subprocess.run(
                cmd,
                cwd=tmpdir,
                capture_output=True,
                text=True,
                timeout=timeout,
                env=run_env,
            )
            run_ms = int((time.monotonic() - start) * 1000)
            return {
                "ran": True,
                "exit_code": proc.returncode,
                "stdout": proc.stdout,
                "stderr": proc.stderr,
                "run_ms": run_ms,
            }
        except subprocess.TimeoutExpired:
            run_ms = int((time.monotonic() - start) * 1000)
            return {"ran": False, "exit_code": None, "stdout": "", "stderr": "TIMEOUT", "run_ms": run_ms}
        except Exception as e:
            run_ms = int((time.monotonic() - start) * 1000)
            return {"ran": False, "exit_code": None, "stdout": "", "stderr": str(e), "run_ms": run_ms}


# ---------------------------------------------------------------------------
# Artifact writing
# ---------------------------------------------------------------------------

def write_artifacts(
    artifact_dir: Path,
    code: str,
    language: str,
    exec_result: dict[str, Any],
    verify_result: dict[str, Any],
) -> None:
    artifact_dir.mkdir(parents=True, exist_ok=True)
    ext = LANG_EXT[language]
    (artifact_dir / f"solution{ext}").write_text(code)
    (artifact_dir / "stdout.txt").write_text(exec_result["stdout"])
    (artifact_dir / "stderr.txt").write_text(exec_result["stderr"])
    (artifact_dir / "result.json").write_text(json.dumps(verify_result, indent=2))


# ---------------------------------------------------------------------------
# Output
# ---------------------------------------------------------------------------

def results_file(model_safe: str) -> Path:
    path = RESULTS_DIR / f"v{PROMPT_VERSION:03d}" / model_safe / "results.jsonl"
    path.parent.mkdir(parents=True, exist_ok=True)
    return path


def append_results(records: list[dict[str, Any]], model_safe: str) -> None:
    with open(results_file(model_safe), "a") as f:
        for r in records:
            f.write(json.dumps(r) + "\n")


def print_tables(records: list[dict[str, Any]]) -> None:
    if not records:
        return

    models = list(dict.fromkeys(
        (r["model"], r["thinking"]) for r in records
    ))
    variants = list(dict.fromkeys(
        (r["prompt_variant"], r["language"]) for r in records
    ))
    tests = list(dict.fromkeys(r["test"] for r in records))

    col_headers = [f"{v}/{lang[:2]}" for v, lang in variants]
    col_w = max(len(h) for h in col_headers + ["test"]) + 2

    for model, thinking in models:
        think_label = f"  [think={thinking}]" if thinking is not None else ""
        print(f"\n{'=' * 60}")
        print(f"Model: {model}{think_label}")
        print(f"{'=' * 60}")
        header = f"{'test':<20}" + "".join(h.rjust(col_w) for h in col_headers)
        print(header)
        print("-" * len(header))
        for test in tests:
            row = f"{test:<20}"
            for vid, lang in variants:
                match = next(
                    (
                        r for r in records
                        if r["model"] == model
                        and r["thinking"] == thinking
                        and r["test"] == test
                        and r["prompt_variant"] == vid
                        and r["language"] == lang
                    ),
                    None,
                )
                if match is None:
                    cell = "-"
                else:
                    symbol = "✓" if match.get("passed") else "✗"
                    cell = f"{symbol} {match['ms']}ms"
                row += cell.rjust(col_w)
            print(row)


# ---------------------------------------------------------------------------
# Inference
# ---------------------------------------------------------------------------

def run_one(
    model_cfg: ModelConfig,
    variant: PromptVariant,
    test: TestCase,
    run_id: str,
    timestamp: str,
    artifact_base: Path,
) -> dict[str, Any]:
    language = variant["language"]
    messages = build_messages(variant, test, language)

    is_apfel = model_cfg.get("backend") == "apfel"
    model = model_cfg["name"]
    options = model_cfg["options"]

    if is_apfel:
        apfel_result = apfel_backend.run_prompt(messages, run_id, timestamp)
        response_raw = apfel_result["response"]
        ms = apfel_result.get("ms", 0)
        eval_count = apfel_result.get("eval_count", 0)
        tok_per_sec = apfel_result.get("tok_per_sec", 0.0)
    else:
        payload: dict[str, Any] = {
            "model": model,
            "messages": messages,
            "stream": False,
            **options,
        }
        infer_timeout = model_cfg.get("infer_timeout", INFER_TIMEOUT)
        start = time.monotonic()
        data = ollama_post("/api/chat", payload, timeout=infer_timeout)
        ms = int((time.monotonic() - start) * 1000)
        response_raw = data["message"]["content"].strip()
        eval_count = data.get("eval_count", 0)
        eval_dur_ns = data.get("eval_duration", 1)
        tok_per_sec = round(eval_count / (eval_dur_ns / 1e9), 1) if eval_dur_ns else 0.0

    code, code_extracted = extract_code(response_raw)

    input_dir = test.get("input_dir")
    test_args = test.get("test_args", [])
    exec_timeout = model_cfg.get("exec_timeout", EXEC_TIMEOUT)
    exec_result = run_code(code, language, input_dir, test_args, timeout=exec_timeout)

    verify_result: dict[str, Any]
    try:
        verify_result = test["verify_mod"].verify(
            exec_result["stdout"],
            exec_result["stderr"],
            exec_result["exit_code"],
            language,
            code,
        )
    except Exception as e:
        verify_result = {"checks": {"verify_error": True}, "passed": False}
        print(f"    verify() error: {e}", file=sys.stderr)

    artifact_dir = artifact_base / language / test["id"]
    write_artifacts(artifact_dir, code, language, exec_result, verify_result)

    record: dict[str, Any] = {
        "run_id":          run_id,
        "timestamp":       timestamp,
        "prompt_v":        PROMPT_VERSION,
        "test":            test["id"],
        "language":        language,
        "prompt_variant":  variant["id"],
        "model":           model,
        "model_options":   options,
        "thinking":        options.get("think", None),
        "response_raw":    response_raw,
        "code_extracted":  code_extracted,
        "code":            code,
        "ms":              ms,
        "eval_count":      eval_count,
        "tok_per_sec":     tok_per_sec,
        "ran":             exec_result["ran"],
        "exit_code":       exec_result["exit_code"],
        "run_ms":          exec_result["run_ms"],
        "stdout":          exec_result["stdout"],
        "stderr":          exec_result["stderr"],
        "checks":          verify_result.get("checks", {}),
        "passed":          verify_result.get("passed", False),
    }
    if "spontaneous_tests" in verify_result:
        record["spontaneous_tests"] = verify_result["spontaneous_tests"]

    return record


# ---------------------------------------------------------------------------
# Main
# ---------------------------------------------------------------------------

def main() -> None:
    parser = argparse.ArgumentParser(description="Code Generation Benchmark")
    parser.add_argument("runs", nargs="?", type=int, default=1, help="number of times to run (default: 1)")
    parser.add_argument("--apple", action="store_true", help="include Apple on-device model via apfel (port 11435)")
    args = parser.parse_args()

    global _test_cases
    _test_cases = load_tests(Path("tests"))
    if not _test_cases:
        print("ERROR: no tests found in tests/", file=sys.stderr)
        sys.exit(1)

    check_runtimes(LANGUAGES)

    apfel_proc = None
    if args.apple:
        MODELS.append(cast(ModelConfig, apfel_backend.APPLE_MODEL_CONFIG))
        try:
            apfel_proc = apfel_backend.ensure_running()
        except Exception as e:
            print(f"ERROR: could not start apfel: {e}", file=sys.stderr)
            sys.exit(1)

    warmup_messages = build_messages(
        PROMPT_VARIANTS[0],
        {"id": "warmup", "prompt": "Write a Python function that returns 'Hello, World!'", "verify_mod": None},  # type: ignore[typeddict-item]
        "python",
    )

    all_records: list[dict[str, Any]] = []

    try:
        for model_cfg in MODELS:
            model = model_cfg["name"]
            is_apfel = model_cfg.get("backend") == "apfel"
            sep = _c('=' * 60, _CYAN)
            print(f"\n{sep}")
            print(f"Model: {model}  options={model_cfg['options']}")
            print(sep)

            if is_apfel:
                try:
                    apfel_backend.warmup(warmup_messages)
                except Exception as e:
                    print(f"  ERROR warming up: {e}", file=sys.stderr)
                    continue
            else:
                try:
                    preload(model, warmup_messages, model_cfg["options"], model_cfg.get("infer_timeout", INFER_TIMEOUT))
                except Exception as e:
                    print(f"  ERROR preloading: {e}", file=sys.stderr)
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
                    run_space = _c(' ' * 20, _WHITE)
                    print(f"  {run_space}Run {run_num} of {args.runs}")
                    print(f"  {rsep}")
                run_ctx = f"  run {run_num}/{args.runs}" if args.runs > 1 else ""
                run_id = datetime.datetime.now().strftime("%Y%m%d_%H%M%S")
                timestamp = datetime.datetime.now(datetime.timezone.utc).isoformat()
                artifact_base = RESULTS_DIR / f"v{PROMPT_VERSION:03d}" / model_safe / run_id
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
                        print(
                            f"  {count_str} {model_str}{run_ctx}  {lang_str}  {test['id']}",
                            flush=True,
                        )
                        record: dict[str, Any]
                        try:
                            record = run_one(model_cfg, variant, test, run_id, timestamp, artifact_base)
                        except Exception as e:
                            print(f"    ERROR: {e}", file=sys.stderr)
                            record = {
                                "run_id": run_id, "timestamp": timestamp,
                                "prompt_v": PROMPT_VERSION,
                                "test": test["id"], "language": lang,
                                "prompt_variant": variant["id"],
                                "model": model, "model_options": model_cfg["options"],
                                "thinking": model_cfg["options"].get("think", None),
                                "response_raw": f"ERROR: {e}", "code_extracted": False,
                                "code": "", "ms": 0, "eval_count": 0, "tok_per_sec": 0,
                                "ran": False, "exit_code": None, "run_ms": 0,
                                "stdout": "", "stderr": str(e),
                                "checks": {}, "passed": False,
                            }
                        if record.get("passed"):
                            status_str = _c("✅ PASS", _GREEN)
                        else:
                            status_str = _c("❌ FAIL", _RED)
                        print(f"    {status_str}  exit={record.get('exit_code')}  {record['ms']}ms gen  {record.get('run_ms', 0)}ms run", flush=True)
                        run_records.append(record)

                all_records.extend(run_records)
                append_results(run_records, model_safe)
                print(f"\n  Appended {len(run_records)} records to {results_file(model_safe)}")
                passed_n = sum(1 for r in run_records if r.get("passed"))
                failed_n = len(run_records) - passed_n
                print(f"  Score: {_c(f'✅ {passed_n}', _GREEN)}  {_c(f'❌ {failed_n}', _RED)}  ({passed_n}/{len(run_records)})")

            if not is_apfel:
                try:
                    unload(model)
                except Exception as e:
                    print(_c(f"  WARNING: failed to unload: {e}", _YELLOW), file=sys.stderr)

    finally:
        if apfel_proc is not None:
            apfel_backend.teardown(apfel_proc)

    print_tables(all_records)


if __name__ == "__main__":
    main()
