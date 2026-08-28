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
    """Target harness dir for a legacy model dir, or None to skip (already migrated)."""
    if name in HARNESS_DIRS:
        return None
    if name == "apple-foundationmodel":
        return "apple"
    if name in anthropic_aliases or name in _HISTORICAL_ANTHROPIC:
        return "anthropic"
    return "ollama"


def _git(args: list[str]) -> None:
    subprocess.run(["git", *args], check=True)


def migrate(version_dir: Path, aliases: set[str], apply: bool) -> int:
    moves: list[tuple[Path, Path]] = []
    for child in sorted(version_dir.iterdir()):
        if not child.is_dir():
            continue
        target = classify_dir(child.name, aliases)
        if target is None:
            continue
        moves.append((child, version_dir / target / child.name))

    if not moves:
        print("  nothing to migrate")
        return 0

    base = version_dir.parent
    for src, dest in moves:
        print(f"  {'MOVE' if apply else 'DRY '}  {src.relative_to(base)}  ->  {dest.relative_to(base)}")
        if apply:
            dest.parent.mkdir(parents=True, exist_ok=True)
            _git(["mv", str(src), str(dest)])

    if not apply:
        print(f"\n  {len(moves)} dirs would move. Re-run with --apply.")
    return len(moves)


def main() -> None:
    parser = argparse.ArgumentParser(description="Migrate results into the harness layout")
    parser.add_argument("version", nargs="?", default=None, help="e.g. v002 (default: all vNNN dirs)")
    parser.add_argument("--apply", action="store_true", help="perform the moves (default: dry run)")
    args = parser.parse_args()

    base = Path(__file__).parent
    aliases = {a["alias"] for a in settings_mod.load_settings(base / "settings.toml").anthropic_models()}

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
        print(f"\n=== {vd.name} ===")
        migrate(vd, aliases, args.apply)


if __name__ == "__main__":
    main()
