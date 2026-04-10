#!/usr/bin/env python3
"""
Helper for the Claude API codegen benchmark.
Runs generated code for one test, verifies output, writes artifacts,
and prints the result record as a JSON line to stdout.

Usage:
    python3 run_claude_test.py <language> <test_id> <code_file> <run_id> \\
        [--model MODEL] [--gen-ms N] [--results-dir PATH]

Exit code: 0 = PASS, 1 = FAIL
"""

import argparse
import datetime
import importlib.util
import json
import os
import shutil
import subprocess
import sys
import tempfile
import time
import types
from pathlib import Path

PROMPT_VERSION = 2

# Maps friendly name → model ID.
# friendly name: used for results dir (results/{name}{PROMPT_VERSION:03d}/) and artifact paths
# model ID:      recorded in results.jsonl
# Comment out models you don't want to run.
MODELS: dict[str, str] = {
    # "opus":   "claude-opus-4-6",
    "sonnet": "claude-sonnet-4-6",
    # "haiku":  "claude-haiku-4-5-20251001",
}

LANG_EXT: dict[str, str] = {
    "python":     ".py",
    "typescript": ".ts",
    "go":         ".go",
    "csharp":     ".csx",
}
LANG_RUN: dict[str, list[str]] = {
    "python":     ["python3"],
    "typescript": ["tsx"],
    "go":         ["go", "run"],
    "csharp":     ["dotnet", "script"],
}
TEST_ARGS: dict[str, list[str]] = {
    "003_fibonacci":     ["100"],
    "008_prime_numbers": ["50"],
}

BASE_DIR = Path(__file__).parent


def _detect_csharp_env() -> dict[str, str] | None:
    """Return a patched env for dotnet-script on Homebrew setups, or None if not needed."""

    def run_check(cmd: str, env: dict[str, str] | None = None) -> bool:
        try:
            r = subprocess.run(cmd, shell=True, capture_output=True, timeout=10, env=env)
            return r.returncode == 0
        except Exception:
            return False

    if run_check("dotnet script --version"):
        return None  # already works fine

    tools_dir = str(Path.home() / ".dotnet" / "tools")
    current_path = os.environ.get("PATH", "")
    patched_path = (
        f"{tools_dir}:{current_path}" if tools_dir not in current_path else current_path
    )
    candidate_env: dict[str, str] = {**os.environ, "PATH": patched_path}

    try:
        r = subprocess.run(
            "dotnet --info", shell=True, capture_output=True, text=True, timeout=10
        )
        for line in r.stdout.splitlines():
            if "Base Path:" in line:
                sdk_path = line.split("Base Path:")[-1].strip()
                dotnet_root = str(Path(sdk_path).parent.parent)
                candidate_env["DOTNET_ROOT"] = dotnet_root
                break
    except Exception:
        pass

    if run_check("dotnet script --version", candidate_env):
        return candidate_env
    return None


def run_code(
    code: str,
    language: str,
    input_dir: Path | None,
    test_args: list[str],
    timeout: int = 60,
) -> dict[str, object]:
    ext = LANG_EXT[language]
    run_cmd = LANG_RUN[language]

    run_env: dict[str, str] | None = None
    if language == "csharp":
        run_env = _detect_csharp_env()

    with tempfile.TemporaryDirectory() as tmpdir:
        tmp = Path(tmpdir)
        sol = tmp / f"solution{ext}"
        sol.write_text(code)
        if input_dir and input_dir.is_dir():
            shutil.copytree(input_dir, tmp / "input")
        cmd = run_cmd + [f"solution{ext}"] + test_args
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
            return {
                "ran": False,
                "exit_code": None,
                "stdout": "",
                "stderr": "TIMEOUT",
                "run_ms": run_ms,
            }
        except Exception as e:
            return {
                "ran": False,
                "exit_code": None,
                "stdout": "",
                "stderr": str(e),
                "run_ms": 0,
            }


def load_verify(test_dir: Path) -> types.ModuleType:
    spec = importlib.util.spec_from_file_location("verify", test_dir / "verify.py")
    assert spec and spec.loader
    mod = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(mod)  # type: ignore[union-attr]
    return mod


def main() -> None:
    parser = argparse.ArgumentParser(description="Claude API codegen benchmark helper")
    parser.add_argument("language", help="Language key (python, typescript, go, csharp)")
    parser.add_argument("test_id", help="Test directory name (e.g. 001_csv_to_json)")
    parser.add_argument("code_file", help="Path to the generated solution file")
    parser.add_argument("run_id", help="Run ID (e.g. 20260407_140352)")
    parser.add_argument("--model", default="sonnet", help="Friendly model name: opus, sonnet, haiku (default: sonnet)")
    parser.add_argument("--gen-ms", type=int, default=0, dest="gen_ms", help="Generation time in ms (default: 0)")
    parser.add_argument("--results-dir", dest="results_dir", default=None, help="Results root directory (default: results/v{PROMPT_VERSION:03d}/)")
    args = parser.parse_args()

    friendly = args.model
    model = MODELS.get(friendly, friendly)  # fall back to raw value if not a known alias
    results_dir = (
        Path(args.results_dir)
        if args.results_dir
        else BASE_DIR / "results" / f"v{PROMPT_VERSION:03d}"
    )

    language = args.language
    test_id = args.test_id
    run_id = args.run_id
    gen_ms = args.gen_ms

    code = Path(args.code_file).read_text()
    test_dir = BASE_DIR / "tests" / test_id

    input_dir = test_dir / "input" if (test_dir / "input").is_dir() else None
    test_args = TEST_ARGS.get(test_id, [])

    exec_result = run_code(code, language, input_dir, test_args)

    verify_mod = load_verify(test_dir)
    try:
        verify_result: dict[str, object] = verify_mod.verify(
            exec_result["stdout"],
            exec_result["stderr"],
            exec_result["exit_code"],
            language,
            code,
        )
    except Exception as e:
        verify_result = {"checks": {"verify_error": True}, "passed": False}
        print(f"verify() error: {e}", file=sys.stderr)

    ext = LANG_EXT[language]
    artifact_dir = results_dir / friendly / run_id / language / test_id
    artifact_dir.mkdir(parents=True, exist_ok=True)
    (artifact_dir / f"solution{ext}").write_text(code)
    (artifact_dir / "stdout.txt").write_text(str(exec_result["stdout"]))
    (artifact_dir / "stderr.txt").write_text(str(exec_result["stderr"]))
    (artifact_dir / "result.json").write_text(json.dumps(verify_result, indent=2))

    passed = bool(verify_result.get("passed", False))
    record: dict[str, object] = {
        "run_id":         run_id,
        "timestamp":      datetime.datetime.now(datetime.timezone.utc).isoformat(),
        "prompt_v":       PROMPT_VERSION,
        "test":           test_id,
        "language":       language,
        "prompt_variant": "C",
        "model":          model,
        "model_options":  {},
        "thinking":       None,
        "response_raw":   code,
        "code_extracted": False,
        "code":           code,
        "ms":             gen_ms,
        "eval_count":     0,
        "tok_per_sec":    0.0,
        "ran":            exec_result["ran"],
        "exit_code":      exec_result["exit_code"],
        "run_ms":         exec_result["run_ms"],
        "stdout":         exec_result["stdout"],
        "stderr":         exec_result["stderr"],
        "checks":         verify_result.get("checks", {}),
        "passed":         passed,
    }
    if "spontaneous_tests" in verify_result:
        record["spontaneous_tests"] = verify_result["spontaneous_tests"]

    results_jsonl = results_dir / friendly / "results.jsonl"
    results_jsonl.parent.mkdir(parents=True, exist_ok=True)
    with open(results_jsonl, "a") as f:
        f.write(json.dumps(record) + "\n")

    status = "PASS" if passed else "FAIL"
    print(
        f"  {status}  model={model}  lang={language}  test={test_id}"
        f"  exit={exec_result['exit_code']}  gen_ms={gen_ms}  run_ms={exec_result['run_ms']}",
    )
    sys.exit(0 if passed else 1)


if __name__ == "__main__":
    main()
