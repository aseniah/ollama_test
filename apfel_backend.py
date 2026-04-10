#!/usr/bin/env python3
"""
Shared apfel backend for benchmarks.

Handles server lifecycle and inference for Apple's on-device model via
apfel's OpenAI-compatible server (https://github.com/Arthur-Ficial/apfel).

Port 11435 is used to avoid conflicting with Ollama on 11434.

Usage from a benchmark script:
    import apfel_backend

    proc = apfel_backend.ensure_running()
    try:
        apfel_backend.warmup(messages)
        result = apfel_backend.run_prompt(messages, run_id, timestamp)
        # result keys: run_id, timestamp, model, model_options, response,
        #              ms, eval_count, tok_per_sec
    finally:
        if proc is not None:
            apfel_backend.teardown(proc)
"""

import json
import subprocess
import time
import urllib.error
import urllib.request
from typing import Any, cast

APFEL_BASE = "http://localhost:11435"
MODEL_NAME = "apple-foundationmodel"

APPLE_MODEL_CONFIG: dict[str, Any] = {
    "name": MODEL_NAME,
    "options": {},
    "backend": "apfel",
}


# ---------------------------------------------------------------------------
# Internal
# ---------------------------------------------------------------------------

def _post(payload: dict[str, Any], timeout: int = 120) -> dict[str, Any]:
    data = json.dumps(payload).encode()
    req = urllib.request.Request(
        f"{APFEL_BASE}/v1/chat/completions",
        data=data,
        headers={"Content-Type": "application/json"},
        method="POST",
    )
    try:
        with urllib.request.urlopen(req, timeout=timeout) as resp:
            return json.loads(resp.read())
    except urllib.error.HTTPError as e:
        try:
            body = cast(dict[str, Any], json.loads(e.read().decode(errors="replace")))
            err = body.get("error", "")
            if isinstance(err, dict):
                err_dict = cast(dict[str, Any], err)
                msg = str(err_dict.get("message") or err_dict)
                err_type = str(err_dict.get("type") or "")
                detail = f"{msg} [{err_type}]" if err_type else msg
            else:
                detail = str(err) or e.reason
        except Exception:
            detail = e.reason
        raise RuntimeError(f"HTTP {e.code}: {detail}") from e


def _health_check(timeout: int = 3) -> bool:
    try:
        req = urllib.request.Request(f"{APFEL_BASE}/health", method="GET")
        with urllib.request.urlopen(req, timeout=timeout):
            return True
    except Exception:
        return False


# ---------------------------------------------------------------------------
# Server lifecycle
# ---------------------------------------------------------------------------

def ensure_running() -> "subprocess.Popen[bytes] | None":
    """
    Check if apfel is already running on port 11435. If so, return None
    (caller must not tear it down). If not, start it and return the Popen
    handle so the caller can tear it down when done.
    """
    if _health_check():
        print("  apfel: already running on port 11435", flush=True)
        return None

    print("  apfel: starting server on port 11435...", flush=True)
    proc: subprocess.Popen[bytes] = subprocess.Popen(
        ["apfel", "--serve", "--port", "11435", "--permissive"],
        stdout=subprocess.DEVNULL,
        stderr=subprocess.DEVNULL,
    )

    for _ in range(20):
        time.sleep(0.5)
        if _health_check(timeout=2):
            print("  apfel: server ready", flush=True)
            return proc

    proc.terminate()
    raise RuntimeError("apfel server did not become ready within 10 seconds")


def teardown(proc: "subprocess.Popen[bytes]") -> None:
    """Stop an apfel server process started by ensure_running."""
    print("  apfel: stopping server...", flush=True)
    proc.terminate()
    try:
        proc.wait(timeout=5)
    except subprocess.TimeoutExpired:
        proc.kill()


# ---------------------------------------------------------------------------
# Inference
# ---------------------------------------------------------------------------

def warmup(messages: list[dict[str, Any]]) -> None:
    """Send a single warm-up request before the test loop."""
    print(f"  warming up {MODEL_NAME}...", flush=True)
    _post(
        {"model": MODEL_NAME, "messages": messages, "stream": False},
        timeout=60,
    )


def run_prompt(
    messages: list[dict[str, Any]],
    run_id: str,
    timestamp: str,
) -> dict[str, Any]:
    """
    Run a single inference request against apfel. Returns timing and response
    fields only — the caller is responsible for adding prompt metadata
    (variant, label, task, etc.) and any benchmark-specific fields (clean, etc.).

    Returned keys: run_id, timestamp, model, model_options, response,
                   ms, eval_count, tok_per_sec
    """
    payload: dict[str, Any] = {
        "model": MODEL_NAME,
        "messages": messages,
        "stream": False,
    }

    start = time.monotonic()
    data = _post(payload)
    elapsed = time.monotonic() - start
    ms = int(elapsed * 1000)

    response = data["choices"][0]["message"]["content"].strip()
    completion_tokens = data.get("usage", {}).get("completion_tokens", 0)
    tok_per_sec = round(completion_tokens / elapsed, 1) if elapsed > 0 else 0

    return {
        "run_id":        run_id,
        "timestamp":     timestamp,
        "model":         MODEL_NAME,
        "model_options": {},
        "response":      response,
        "ms":            ms,
        "eval_count":    completion_tokens,
        "tok_per_sec":   tok_per_sec,
    }
