#!/usr/bin/env python3
"""Identify the machine a benchmark run executes on.

Benchmark results are only comparable within one machine — token rates and
latency are hardware-bound. This module detects the host's hardware, derives a
stable fingerprint, and maps it to a short human-set slug via a committed
registry at ``machines/<slug>.toml``. Results then land under
``results/v{NNN}/<slug>/<harness>/<model>/``.

CLI:
  python3 machine.py                 # print detected specs + resolved slug
  python3 machine.py --json          # same, machine-readable
  python3 machine.py --register foo  # register this machine as slug "foo"
"""

from __future__ import annotations

import argparse
import datetime
import hashlib
import json
import platform
import re
import socket
import subprocess
import sys
import tomllib
from pathlib import Path
from typing import Any, TypedDict, cast

MACHINES_DIR = Path("machines")


class MachineError(Exception):
    """Raised when the running machine cannot be resolved to a slug."""


class MachineSpec(TypedDict):
    chip: str
    model_identifier: str
    model_name: str
    cpu_cores: int
    cpu_cores_performance: int
    cpu_cores_efficiency: int
    memory_gb: int
    arch: str
    os: str
    os_build: str
    python: str
    hostname: str
    ollama_version: str | None
    lms_version: str | None


class MachineSnapshot(TypedDict):
    """A MachineSpec plus the slug/fingerprint/time it was captured."""

    slug: str
    fingerprint: str
    captured_at: str
    specs: MachineSpec


# ---------------------------------------------------------------------------
# Detection
# ---------------------------------------------------------------------------

def _run(cmd: list[str]) -> str:
    try:
        out = subprocess.run(cmd, capture_output=True, text=True, timeout=10, check=False)
    except (OSError, subprocess.SubprocessError):
        return ""
    return f"{out.stdout}\n{out.stderr}"


def _parse_processors(value: str) -> tuple[int, int, int]:
    """``"proc 16:12:4:0"`` -> (total=16, performance=12, efficiency=4)."""
    nums = [int(n) for n in re.findall(r"\d+", value)]
    if len(nums) >= 3:
        return nums[0], nums[1], nums[2]
    if len(nums) == 1:
        return nums[0], 0, 0
    return 0, 0, 0


def _semver(text: str) -> str | None:
    m = re.search(r"\d+\.\d+\.\d+", text)
    return m.group(0) if m else None


def _mac_hardware() -> dict[str, Any]:
    raw = _run(["system_profiler", "SPHardwareDataType", "-json"])
    try:
        parsed: Any = json.loads(raw)
        first: Any = parsed["SPHardwareDataType"][0]
    except (json.JSONDecodeError, KeyError, IndexError, TypeError):
        return {}
    if not isinstance(first, dict):
        return {}
    return cast("dict[str, Any]", first)


def detect() -> MachineSpec:
    """Gather this host's hardware and toolchain identity (best-effort)."""
    hw = _mac_hardware()
    total, perf, eff = _parse_processors(str(hw.get("number_processors", "")))
    mem_match = re.search(r"\d+", str(hw.get("physical_memory", "")))

    return MachineSpec(
        chip=str(hw.get("chip_type") or hw.get("cpu_type") or "unknown"),
        model_identifier=str(hw.get("machine_model", "unknown")),
        model_name=str(hw.get("machine_name", "unknown")),
        cpu_cores=total,
        cpu_cores_performance=perf,
        cpu_cores_efficiency=eff,
        memory_gb=int(mem_match.group(0)) if mem_match else 0,
        arch=platform.machine(),
        os=platform.mac_ver()[0] or platform.release(),
        os_build=_run(["sw_vers", "-buildVersion"]).strip(),
        python=platform.python_version(),
        hostname=socket.gethostname(),
        ollama_version=_semver(_run(["ollama", "--version"])),
        lms_version=_semver(_run(["lms", "version"])) or _semver(_run(["lms", "--version"])),
    )


def fingerprint(spec: MachineSpec) -> str:
    """Stable ID from hardware class only — survives OS/toolchain updates.

    Deliberately excludes the serial number and hostname: two identical Macs
    share a fingerprint (same benchmark class), and the slug disambiguates.
    """
    basis = "|".join([
        spec["model_identifier"],
        spec["chip"],
        str(spec["cpu_cores"]),
        str(spec["memory_gb"]),
        spec["arch"],
    ])
    return hashlib.sha256(basis.encode()).hexdigest()[:12]


def suggest_slug(spec: MachineSpec) -> str:
    """``"Apple M3 Max"`` + 48 GB -> ``"m3-max-48gb"``."""
    chip = spec["chip"].lower()
    chip = re.sub(r"\bapple\s+", "", chip)
    chip = re.sub(r"[^a-z0-9]+", "-", chip).strip("-")
    gb = spec["memory_gb"]
    slug = f"{chip}-{gb}gb" if gb else chip
    return slug or "unknown-machine"


# ---------------------------------------------------------------------------
# Registry
# ---------------------------------------------------------------------------

class RegistryEntry(TypedDict):
    slug: str
    fingerprint: str
    label: str
    registered_at: str
    specs: dict[str, Any]


def _registry_path(slug: str, machines_dir: Path) -> Path:
    return machines_dir / f"{slug}.toml"


def load_registry(machines_dir: Path = MACHINES_DIR) -> list[RegistryEntry]:
    if not machines_dir.is_dir():
        return []
    entries: list[RegistryEntry] = []
    for path in sorted(machines_dir.glob("*.toml")):
        try:
            raw = cast("dict[str, Any]", tomllib.loads(path.read_text()))
        except (OSError, tomllib.TOMLDecodeError):
            continue
        if "fingerprint" not in raw:
            continue
        specs = cast("dict[str, Any]", raw.get("specs") or {})
        entries.append(RegistryEntry(
            slug=str(raw.get("slug", path.stem)),
            fingerprint=str(raw["fingerprint"]),
            label=str(raw.get("label", "")),
            registered_at=str(raw.get("registered_at", "")),
            specs=specs,
        ))
    return entries


def _toml_value(value: str | int | None) -> str:
    if isinstance(value, bool):
        return "true" if value else "false"
    if isinstance(value, int):
        return str(value)
    return '"' + str(value).replace("\\", "\\\\").replace('"', '\\"') + '"'


def register(slug: str, spec: MachineSpec, machines_dir: Path = MACHINES_DIR) -> Path:
    """Write ``machines/<slug>.toml`` for this machine. Returns the path."""
    if not re.fullmatch(r"[a-z0-9]+(?:-[a-z0-9]+)*", slug):
        raise MachineError(
            f"invalid slug {slug!r}: use lowercase letters, digits and dashes"
        )
    machines_dir.mkdir(parents=True, exist_ok=True)
    path = _registry_path(slug, machines_dir)
    fp = fingerprint(spec)
    label = f"{spec['model_name']} ({spec['chip']}, {spec['memory_gb']} GB)"
    now = datetime.datetime.now(datetime.timezone.utc).replace(microsecond=0).isoformat()

    lines = [
        f"slug = {_toml_value(slug)}",
        f"fingerprint = {_toml_value(fp)}",
        f"label = {_toml_value(label)}",
        f"registered_at = {_toml_value(now)}",
        "",
        "[specs]",
    ]
    for key, value in cast("dict[str, str | int | None]", spec).items():
        if value is None:
            continue
        lines.append(f"{key} = {_toml_value(value)}")
    path.write_text("\n".join(lines) + "\n")
    return path


# ---------------------------------------------------------------------------
# Resolution
# ---------------------------------------------------------------------------

def resolve(
    *,
    register_slug: str | None = None,
    interactive: bool | None = None,
    machines_dir: Path = MACHINES_DIR,
) -> str:
    """Return the slug for this machine, registering it if new.

    - Known fingerprint -> its slug (must match ``register_slug`` if given).
    - Unknown + ``register_slug`` -> write a new registry entry, return it.
    - Unknown + interactive TTY -> prompt for a slug, register, return it.
    - Unknown + non-interactive -> raise MachineError.
    """
    spec = detect()
    fp = fingerprint(spec)
    registry = load_registry(machines_dir)
    by_fp = {e["fingerprint"]: e for e in registry}

    if fp in by_fp:
        known = by_fp[fp]["slug"]
        if register_slug and register_slug != known:
            raise MachineError(
                f"this machine is already registered as {known!r}; "
                f"remove machines/{known}.toml to rename it"
            )
        return known

    if register_slug:
        register(register_slug, spec, machines_dir)
        return register_slug

    if interactive is None:
        interactive = sys.stdin.isatty()
    if interactive:
        default = suggest_slug(spec)
        print(f"This machine ({spec['chip']}, {spec['memory_gb']} GB) is not registered.")
        reply = input(f"  slug for it [{default}]: ").strip() or default
        register(reply, spec, machines_dir)
        print(f"  wrote machines/{reply}.toml")
        return reply

    raise MachineError(
        f"unknown machine (fingerprint {fp}); pass --machine <slug> to register it"
    )


def snapshot(slug: str, spec: MachineSpec | None = None) -> MachineSnapshot:
    spec = spec or detect()
    return MachineSnapshot(
        slug=slug,
        fingerprint=fingerprint(spec),
        captured_at=datetime.datetime.now(datetime.timezone.utc)
        .replace(microsecond=0).isoformat(),
        specs=spec,
    )


# ---------------------------------------------------------------------------
# CLI
# ---------------------------------------------------------------------------

def main() -> None:
    parser = argparse.ArgumentParser(description="Identify the benchmark host machine")
    parser.add_argument("--register", metavar="SLUG", default=None,
                        help="register this machine under SLUG")
    parser.add_argument("--json", action="store_true", help="emit JSON")
    args = parser.parse_args()

    spec = detect()
    fp = fingerprint(spec)
    known = {e["fingerprint"]: e["slug"] for e in load_registry()}

    if args.register:
        path = register(args.register, spec)
        slug = args.register
        if not args.json:
            print(f"wrote {path}")
    else:
        slug = known.get(fp)

    if args.json:
        print(json.dumps({"slug": slug, "fingerprint": fp, "specs": spec}, indent=2))
        return

    print(f"fingerprint : {fp}")
    print(f"slug        : {slug or '(unregistered — run with --register SLUG)'}")
    print(f"chip        : {spec['chip']}")
    print(f"cores       : {spec['cpu_cores']} "
          f"({spec['cpu_cores_performance']}P + {spec['cpu_cores_efficiency']}E)")
    print(f"memory      : {spec['memory_gb']} GB")
    print(f"model       : {spec['model_identifier']} ({spec['model_name']})")
    print(f"os          : macOS {spec['os']} ({spec['os_build']})  {spec['arch']}")
    print(f"python      : {spec['python']}")
    print(f"ollama      : {spec['ollama_version'] or '-'}")
    print(f"lms         : {spec['lms_version'] or '-'}")
    if not slug:
        print(f"\nsuggested slug: {suggest_slug(spec)}")


if __name__ == "__main__":
    main()
