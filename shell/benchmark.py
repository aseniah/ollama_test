#!/usr/bin/env python3
"""
iTerm2 Ollama Backend Benchmark

Tests local Ollama models for suitability as an iTerm2 AI plugin backend.
Results are appended to results.jsonl — run as many times as you like.

Usage:
  python3 benchmark.py              # single run  → results/v009.jsonl
  python3 benchmark.py 3            # 3 runs, results aggregated at end
  python3 benchmark.py --apple      # include Apple on-device model
  python3 benchmark.py 3 --apple    # 3 runs with Apple model
"""

import argparse
import datetime
import json
import sys
import time
import urllib.error
import urllib.request
from pathlib import Path
from typing import Any, NotRequired, TypedDict, cast

sys.path.insert(0, str(Path(__file__).parent.parent))
import apfel_backend

OLLAMA_BASE = "http://localhost:11434"
RESULTS_DIR = Path("results")

PROMPT_VERSION = 14

class ModelConfig(TypedDict):
    name: str
    options: dict[str, Any]
    backend: NotRequired[str]


class PromptVariant(TypedDict):
    id: str
    desc: str
    system_prompt: str


class TestPrompt(TypedDict):
    label: str
    task: str


MODELS: list[ModelConfig] = [
    {"name": "qwen3.5:27b",         "options": {"think": False}},  # 2x slower, "think" leakage
    {"name": "qwen3.5:27b-nvfp4",   "options": {"think": False}},  # MLX support should be much faster
    {"name": "qwen3.5:4b",          "options": {"think": False}},  # 2x slower, "think" leakage
    {"name": "qwen3.5:4b-nvfp4",    "options": {"think": False}},  # MLX support should be much faster
    {"name": "qwen2.5-coder:7b",    "options": {}},
    {"name": "qwen3-coder:30b",     "options": {}},
    # {"name": "qwen3-coder:30b-64k", "options": {}},
    # {"name": "qwen3-coder:30b-32k", "options": {}},
    # {"name": "mistral:latest",      "options": {}},
    # {"name": "llama3.2:latest",     "options": {}},
]

_PROMPT_SUFFIX = (
    "Your entire response must be the raw command only. "
    "Do not use backticks, code fences, or any markdown. "
    "Do not explain or add commentary. Just the command, nothing else."
)

PROMPT_VARIANTS: list[PromptVariant] = [
    {
        "id": "A",
        "desc": "v7C verbatim (confirmation)",
        "system_prompt": (
            "Output exactly one shell command for zsh on macOS. "
            "macOS uses BSD userland, not GNU coreutils. "
            "GNU-only flags that do NOT work on macOS: "
            "-executable for find (use -perm +111), "
            "--sort for ps (pipe to sort -k instead), "
            "--max-depth for du (use -d instead), "
            "symbolic perm notation like /u=x (use octal instead). "
            "To sort processes by memory on macOS, use: ps aux | sort -k4nr | head -n N. "
            "To filter grep by file type recursively, use --include='*.ext' with a path "
            "(e.g. grep -r 'x' --include='*.js' .), not shell globs (grep -r 'x' *.js is wrong). "
            + _PROMPT_SUFFIX
        ),
    },
    {
        "id": "B",
        "desc": "compressed",
        "system_prompt": (
            "Output exactly one shell command for zsh on macOS. "
            "macOS is BSD, not GNU — these flags do not exist on macOS: "
            "-executable (use -perm +111), "
            "ps --sort (use ps aux | sort -k4nr | head -N), "
            "--max-depth (use -d N), "
            "symbolic perms like /u=x (use octal), "
            "grep *.ext globs (use grep -r 'x' --include='*.ext' .). "
            + _PROMPT_SUFFIX
        ),
    },
    {
        "id": "C",
        "desc": "iTerm2 candidate (no grep)",
        "system_prompt": (
            "Output exactly one shell command for zsh on macOS. "
            "macOS uses BSD userland, not GNU coreutils. "
            "GNU-only flags that do NOT work on macOS: "
            "-executable for find (use -perm +111), "
            "--sort for ps (use ps aux | sort -k4nr | head -n N instead), "
            "--max-depth for du (use -d instead), "
            "symbolic perm notation like /u=x (use octal instead). "
            + _PROMPT_SUFFIX
        ),
    },
]

TEST_PROMPTS: list[TestPrompt] = [
    {"label": "git log",            "task": "show the last 10 git commits with hash and message, one per line"},
    {"label": "top 5 by memory",    "task": "show the top 5 processes by memory usage"},
    {"label": "tar with exclusions","task": "create a tar.gz of the current directory named backup.tar.gz, excluding .git and node_modules"},
    {"label": "recursive grep",     "task": "search recursively for the string TODO in all .js files in the current directory"},
    {"label": "dir sizes",          "task": "show disk usage of each subdirectory in the current directory, sorted by size"},
    {"label": "find executables",   "task": "find all executable files in this directory and its subdirectories"},
    {"label": "sed in-place",       "task": "replace all occurrences of 'foo' with 'bar' in place in a file named file.txt"},
    {"label": "stat file size",     "task": "print the size in bytes of a file named data.txt"},
]


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
            return json.loads(resp.read())
    except urllib.error.HTTPError as e:
        raise RuntimeError(_http_error_message(e)) from e


def _http_error_message(e: urllib.error.HTTPError) -> str:
    """Extract a meaningful message from an HTTPError response body."""
    try:
        body = cast(dict[str, Any], json.loads(e.read().decode(errors="replace")))
        # Ollama: {"error": "..."}, apfel/OpenAI: {"error": {"message": "..."}}
        err = body.get("error", "")
        if isinstance(err, dict):
            err_dict = cast(dict[str, Any], err)
            return f"HTTP {e.code}: {str(err_dict.get('message') or err_dict)}"
        return f"HTTP {e.code}: {str(err) or e.reason}"
    except Exception:
        return f"HTTP {e.code}: {e.reason}"


def preload(model: str, variant: PromptVariant, options: dict[str, Any]) -> None:
    print(f"  preloading {model}...", flush=True)
    ollama_post("/api/chat", {"model": model, "messages": [], "keep_alive": -1}, timeout=60)
    print(f"  warming up {model}...", flush=True)
    ollama_post("/api/chat", {
        "model": model,
        "messages": [
            {"role": "system", "content": variant["system_prompt"]},
            {"role": "user",   "content": "It must do this: print the current date"},
        ],
        "stream": False,
        **options,
    })


def unload(model: str) -> None:
    print(f"  unloading {model}...", flush=True)
    ollama_post("/api/chat", {"model": model, "messages": [], "keep_alive": 0}, timeout=15)


# ---------------------------------------------------------------------------
# Inference
# ---------------------------------------------------------------------------

def is_clean(text: str) -> bool:
    """Single shell command: no fences, no backtick wrapping, no prose sentences."""
    stripped = text.strip()
    if stripped.startswith("`") and stripped.endswith("`"):
        return False
    lines = [ln for ln in stripped.splitlines() if ln.strip()]
    if not lines:
        return False
    for line in lines:
        if line.strip().startswith("```"):
            return False
        shell_tokens = ("|", "$", "&&", "||", ";", "/", "\\", "-", ">", "<", "(", ")")
        s = line.strip()
        if (
            s[0].isupper()
            and s[-1] in ".!?"
            and " " in s
            and not any(t in s for t in shell_tokens)
        ):
            return False
    return True


def run_prompt(model_cfg: ModelConfig, variant: PromptVariant, prompt: TestPrompt, run_id: str, timestamp: str) -> dict[str, Any]:
    if model_cfg.get("backend") == "apfel":
        messages = [
            {"role": "system", "content": variant["system_prompt"]},
            {"role": "user",   "content": f"It must do this: {prompt['task']}"},
        ]
        apfel_result = apfel_backend.run_prompt(messages, run_id, timestamp)
        return {
            **apfel_result,
            "prompt_v":       PROMPT_VERSION,
            "prompt_variant": variant["id"],
            "prompt_label":   prompt["label"],
            "task":           prompt["task"],
            "clean":          is_clean(apfel_result["response"]),
        }

    model = model_cfg["name"]
    options = model_cfg["options"]

    payload: dict[str, Any] = {
        "model": model,
        "messages": [
            {"role": "system", "content": variant["system_prompt"]},
            {"role": "user",   "content": f"It must do this: {prompt['task']}"},
        ],
        "stream": False,
        **options,
    }

    start = time.monotonic()
    data = ollama_post("/api/chat", payload)
    ms = int((time.monotonic() - start) * 1000)

    response = data["message"]["content"].strip()
    eval_count = data.get("eval_count", 0)
    eval_dur_ns = data.get("eval_duration", 1)
    tok_per_sec = round(eval_count / (eval_dur_ns / 1e9), 1) if eval_dur_ns else 0

    return {
        "run_id":         run_id,
        "timestamp":      timestamp,
        "prompt_v":       PROMPT_VERSION,
        "prompt_variant": variant["id"],
        "model":          model,
        "model_options":  options,
        "prompt_label":   prompt["label"],
        "task":           prompt["task"],
        "response":       response,
        "ms":             ms,
        "eval_count":     eval_count,
        "tok_per_sec":    tok_per_sec,
        "clean":          is_clean(response),
    }


# ---------------------------------------------------------------------------
# Output
# ---------------------------------------------------------------------------

def results_file() -> Path:
    RESULTS_DIR.mkdir(exist_ok=True)
    return RESULTS_DIR / f"v{PROMPT_VERSION:03d}.jsonl"


def append_results(records: list[dict[str, Any]]) -> None:
    with open(results_file(), "a") as f:
        for r in records:
            f.write(json.dumps(r) + "\n")


def print_tables(all_records: list[dict[str, Any]]) -> None:
    """One table per model, prompt variants as columns. Aggregates across runs."""
    variant_ids = [v["id"] for v in PROMPT_VARIANTS]
    label_w = 22
    col_w = 30
    total_w = label_w + col_w * len(variant_ids)

    for model_cfg in MODELS:
        model = model_cfg["name"]
        model_records = [r for r in all_records if r["model"] == model]
        by_variant = {v: [r for r in model_records if r["prompt_variant"] == v] for v in variant_ids}

        print(f"\n{'  ' + model + '  ':=^{total_w}}")
        header = f"\n  {'Task':<{label_w - 2}}"
        for v in variant_ids:
            variant = next(pv for pv in PROMPT_VARIANTS if pv["id"] == v)
            col_title = f"[{v}] {variant['desc']}"
            header += f"{col_title:<{col_w}}"
        print(header)
        print("-" * total_w)

        for prompt in TEST_PROMPTS:
            row = f"  {prompt['label']:<{label_w - 2}}"
            for v in variant_ids:
                records = [r for r in by_variant[v] if r["prompt_label"] == prompt["label"]]
                if records:
                    avg_ms = int(sum(r["ms"] for r in records) / len(records))
                    clean_n = sum(1 for r in records if r["clean"])
                    cell = f"{avg_ms}ms  {clean_n}/{len(records)} ✓"
                else:
                    cell = "—"
                row += f"{cell:<{col_w}}"
            print(row)

        print("-" * total_w)
        row = f"  {'TOTALS':<{label_w - 2}}"
        for v in variant_ids:
            records = by_variant[v]
            if records:
                avg_ms = int(sum(r["ms"] for r in records) / len(records))
                clean_n = sum(1 for r in records if r["clean"])
                cell = f"avg {avg_ms}ms  {clean_n}/{len(records)} ✓"
            else:
                cell = "—"
            row += f"{cell:<{col_w}}"
        print(row)
        print()


def print_responses(all_records: list[dict[str, Any]]) -> None:
    """Per prompt: last recorded response per model+variant (across all runs)."""
    variant_ids = [v["id"] for v in PROMPT_VARIANTS]
    if not all_records:
        return

    print(f"\n{'  RESPONSES (latest per model)  ':=^60}\n")
    for prompt in TEST_PROMPTS:
        print(f"[{prompt['label']}]  {prompt['task']}")
        for model_cfg in MODELS:
            model = model_cfg["name"]
            short = model[:20]
            for v in variant_ids:
                matches = [r for r in all_records if r["model"] == model and r["prompt_label"] == prompt["label"] and r["prompt_variant"] == v]
                resp = matches[-1]["response"] if matches else "—"
                lines = resp.splitlines()
                label = f"{short} [{v}]"
                print(f"  {label:<26} {lines[0]}")
                for line in lines[1:]:
                    print(f"  {'':<26} {line}")
        print()


# ---------------------------------------------------------------------------
# Main
# ---------------------------------------------------------------------------

def main() -> None:
    parser = argparse.ArgumentParser(description="iTerm2 Ollama Backend Benchmark")
    parser.add_argument("runs", nargs="?", type=int, default=1, help="number of times to run (default: 1)")
    parser.add_argument("--apple", action="store_true", help="include Apple on-device model via apfel (port 11435)")
    args = parser.parse_args()

    apfel_proc = None
    if args.apple:
        MODELS.append(cast(ModelConfig, apfel_backend.APPLE_MODEL_CONFIG))
        try:
            apfel_proc = apfel_backend.ensure_running()
        except Exception as e:
            print(f"ERROR: could not start apfel: {e}", file=sys.stderr)
            sys.exit(1)

    all_records: list[dict[str, Any]] = []

    try:
        for model_cfg in MODELS:
            model = model_cfg["name"]
            is_apfel = model_cfg.get("backend") == "apfel"
            print(f"\n{'=' * 60}")
            print(f"Model: {model}  options={model_cfg['options']}")
            print(f"{'=' * 60}")

            if is_apfel:
                try:
                    apfel_backend.warmup([
                        {"role": "system", "content": PROMPT_VARIANTS[0]["system_prompt"]},
                        {"role": "user",   "content": "It must do this: print the current date"},
                    ])
                except Exception as e:
                    print(f"  ERROR warming up: {e}", file=sys.stderr)
                    continue
            else:
                try:
                    preload(model, PROMPT_VARIANTS[0], model_cfg["options"])
                except Exception as e:
                    print(f"  ERROR preloading: {e}", file=sys.stderr)
                    continue

            for run_num in range(1, args.runs + 1):
                if args.runs > 1:
                    print(f"\n  {'#' * 50}")
                    print(f"  # Run {run_num} of {args.runs}")
                    print(f"  {'#' * 50}")
                run_id = datetime.datetime.now().strftime("%Y%m%d_%H%M%S")
                timestamp = datetime.datetime.now(datetime.timezone.utc).isoformat()
                run_records: list[dict[str, Any]] = []

                for variant in PROMPT_VARIANTS:
                    print(f"\n  -- Variant {variant['id']}: {variant['desc']} --")
                    for i, prompt in enumerate(TEST_PROMPTS, 1):
                        print(f"  [{i}/{len(TEST_PROMPTS)}] {prompt['label']}...", flush=True)
                        record: dict[str, Any]
                        try:
                            record = run_prompt(model_cfg, variant, prompt, run_id, timestamp)
                        except Exception as e:
                            print(f"    ERROR: {e}", file=sys.stderr)
                            record = {
                                "run_id": run_id, "timestamp": timestamp,
                                "prompt_v": PROMPT_VERSION, "prompt_variant": variant["id"],
                                "model": model, "model_options": model_cfg["options"],
                                "prompt_label": prompt["label"], "task": prompt["task"],
                                "response": f"ERROR: {e}",
                                "ms": 0, "eval_count": 0, "tok_per_sec": 0, "clean": False,
                            }
                        run_records.append(record)
                        icon = "✓" if record["clean"] else "✗"
                        print(f"    {record['ms']}ms  {record['tok_per_sec']}tok/s  {icon}  {record['response'][:60]!r}")

                all_records.extend(run_records)
                append_results(run_records)
                print(f"\n  Appended {len(run_records)} records to {results_file()}")

            if not is_apfel:
                try:
                    unload(model)
                except Exception as e:
                    print(f"  WARNING: failed to unload: {e}", file=sys.stderr)

    finally:
        if apfel_proc is not None:
            apfel_backend.teardown(apfel_proc)

    print_tables(all_records)
    print_responses(all_records)


if __name__ == "__main__":
    main()
