"""One-time migration: results/vNNN/{harness}/ -> results/vNNN/{machine}/{harness}/.

Adds the machine axis to the results tree. Local-harness dirs (ollama, lmstudio,
apple) move under a machine slug; the API-backed anthropic dir moves under the
"api" pseudo-machine. Each moved results.jsonl is backfilled with "machine" and
"harness" fields on any record that lacks them (legacy records predate both).

Dry-run by default. Pass --apply to perform `git mv` + rewrites. Run once, by
hand, on a clean tree.
"""

from __future__ import annotations

import argparse
import json
import subprocess
import sys
from pathlib import Path

import machine as machine_mod

LOCAL_HARNESSES = {"ollama", "lmstudio", "apple"}
API_HARNESSES = {"anthropic"}
API_MACHINE = "api"


def _git(args: list[str]) -> None:
    subprocess.run(["git", *args], check=True)


def _plan(version_dir: Path, slug: str) -> list[tuple[Path, Path, str, str]]:
    """(src harness dir, dest dir, machine slug, harness) for each legacy harness dir.

    Anything that is not a legacy harness dir is assumed already migrated
    (a machine slug) and skipped.
    """
    moves: list[tuple[Path, Path, str, str]] = []
    for child in sorted(version_dir.iterdir()):
        if not child.is_dir():
            continue
        if child.name in LOCAL_HARNESSES:
            moves.append((child, version_dir / slug / child.name, slug, child.name))
        elif child.name in API_HARNESSES:
            moves.append((child, version_dir / API_MACHINE / child.name, API_MACHINE, child.name))
    return moves


def _backfill(results_jsonl: Path, slug: str, harness: str, apply: bool) -> int:
    if not results_jsonl.is_file():
        return 0
    lines = results_jsonl.read_text().splitlines()
    out: list[str] = []
    changed = 0
    for line in lines:
        if not line.strip():
            continue
        rec = json.loads(line)
        added = False
        if "machine" not in rec:
            rec["machine"] = slug
            added = True
        if "harness" not in rec:
            rec["harness"] = harness
            added = True
        changed += added
        out.append(json.dumps(rec))
    if changed and apply:
        results_jsonl.write_text("\n".join(out) + "\n")
    return changed


def migrate(version_dir: Path, slug: str, apply: bool) -> int:
    moves = _plan(version_dir, slug)
    if not moves:
        print("  nothing to migrate")
        return 0

    base = version_dir.parent
    for src, dest, target_slug, harness in moves:
        tag = "MOVE" if apply else "DRY "
        print(f"  {tag}  {src.relative_to(base)}  ->  {dest.relative_to(base)}")
        if apply:
            dest.parent.mkdir(parents=True, exist_ok=True)
            _git(["mv", str(src), str(dest)])
        landing = dest if apply else src
        for jsonl in sorted(landing.glob("*/results.jsonl")):
            n = _backfill(jsonl, target_slug, harness, apply)
            if n:
                verb = "backfilled" if apply else "would backfill"
                print(f"        {verb} machine={target_slug} harness={harness} on {n:>3} records  ({jsonl.parent.name})")

    if not apply:
        print(f"\n  {len(moves)} harness dirs would move. Re-run with --apply.")
    return len(moves)


def main() -> None:
    parser = argparse.ArgumentParser(description="Migrate results into the machine layout")
    parser.add_argument("version", nargs="?", default=None, help="e.g. v002 (default: all vNNN dirs)")
    parser.add_argument("--slug", default=None,
                        help="machine slug for local-harness results (default: resolve this host)")
    parser.add_argument("--apply", action="store_true", help="perform the moves (default: dry run)")
    args = parser.parse_args()

    base = Path(__file__).parent
    slug = args.slug or machine_mod.resolve(machines_dir=base / "machines", interactive=False)

    results = base / "results"
    if args.version:
        version_dirs = [results / args.version]
    else:
        version_dirs = sorted(
            d for d in results.iterdir() if d.is_dir() and d.name.startswith("v")
        )

    for vd in version_dirs:
        if not vd.is_dir():
            print(f"skip: {vd} not found", file=sys.stderr)
            continue
        print(f"\n=== {vd.name} (machine={slug}) ===")
        migrate(vd, slug, args.apply)


if __name__ == "__main__":
    main()
