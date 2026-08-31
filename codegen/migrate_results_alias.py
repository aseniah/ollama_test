"""One-time migration: relabel result records to the alias / model_id split.

Records used to carry `model` = the raw id sent to the harness. They now carry
`model` = a short alias and `model_id` = the raw id. This backfills existing
v{NNN} results and renames the LM Studio dirs (which were named from the raw id)
to match the alias. Ollama and anthropic dir names already match their alias.

Dry-run by default. Pass --apply to rewrite records and `git mv` dirs. Run once,
by hand, on a clean tree.
"""

from __future__ import annotations

import argparse
import json
import sys
from pathlib import Path

# raw model_id -> alias. Ollama ids equal their alias and are left alone; only
# ids that appear under a different alias need listing here.
ALIAS: dict[str, str] = {
    "qwen/qwen3.8-27b@4bit": "qwen3.8-27b-mlx-4bit",
    "qwen/qwen3.8-27b@8bit": "qwen3.8-27b-mlx-8bit",
    "qwen/qwen3.8-27b@q4_k_m": "qwen3.8-27b-gguf-q4km",
    "qwen/qwen3-coder-30b": "qwen3-coder-30b-mlx",
    "claude-haiku-4-5-20251001": "haiku",
    "claude-opus-4-6": "opus",
    "claude-sonnet-4-6": "sonnet",
}


def _relabel(jsonl: Path, apply: bool) -> int:
    lines = jsonl.read_text().splitlines()
    out: list[str] = []
    changed = 0
    for line in lines:
        if not line.strip():
            continue
        rec = json.loads(line)
        raw = str(rec.get("model", ""))
        if "model_id" not in rec:
            rec["model_id"] = raw
            rec["model"] = ALIAS.get(raw, raw)
            changed += 1
        out.append(json.dumps(rec))
    if changed and apply:
        jsonl.write_text("\n".join(out) + "\n")
    return changed


def _safe(model: str) -> str:
    return model.replace(":", "_").replace("/", "_").replace("@", "-")


def _dir_renames(version_dir: Path) -> list[tuple[Path, Path]]:
    """LM Studio dirs were named from the raw id — rename them to the alias.

    The dir name is `{safe(raw)}{suffix}` where suffix is _nothink / _think /
    (_hacked). Swap the raw stub for the alias stub, keeping the suffix.
    """
    renames: list[tuple[Path, Path]] = []
    for lms in version_dir.glob("*/lmstudio"):
        for d in sorted(lms.iterdir()):
            if not d.is_dir():
                continue
            jsonl = d / "results.jsonl"
            if not jsonl.is_file():
                continue
            first = json.loads(jsonl.read_text().splitlines()[0])
            raw = str(first.get("model_id") or first.get("model"))
            alias = ALIAS.get(raw)
            if not alias:
                continue
            old_stub, new_stub = _safe(raw), _safe(alias)
            if new_stub in d.name:
                continue  # already migrated
            suffix = d.name.split(old_stub, 1)[1] if old_stub in d.name else ""
            renames.append((d, d.with_name(new_stub + suffix)))
    return renames


def migrate(version_dir: Path, apply: bool) -> None:
    base = version_dir.parent
    total = 0
    for jsonl in sorted(version_dir.glob("*/*/*/results.jsonl")):
        n = _relabel(jsonl, apply)
        if n:
            total += n
            print(f"  {'relabelled' if apply else 'would relabel'} {n:>3} records  {jsonl.parent.relative_to(base)}")

    for src, dest in _dir_renames(version_dir):
        print(f"  {'MOVE' if apply else 'DRY '}  {src.relative_to(base)}  ->  {dest.relative_to(base)}")
        if apply:
            # plain rename — git detects the move for tracked files on commit;
            # some of these dirs are entirely untracked (never-committed runs).
            src.rename(dest)

    if not apply:
        print(f"\n  {total} records would be relabelled. Re-run with --apply.")


def main() -> None:
    parser = argparse.ArgumentParser(description="Migrate results to the alias/model_id split")
    parser.add_argument("version", nargs="?", default=None, help="e.g. v002 (default: all vNNN dirs)")
    parser.add_argument("--apply", action="store_true", help="perform the changes (default: dry run)")
    args = parser.parse_args()

    results = Path(__file__).parent / "results"
    if args.version:
        version_dirs = [results / args.version]
    else:
        version_dirs = sorted(d for d in results.iterdir() if d.is_dir() and d.name.startswith("v"))

    for vd in version_dirs:
        if not vd.is_dir():
            print(f"skip: {vd} not found", file=sys.stderr)
            continue
        print(f"\n=== {vd.name} ===")
        migrate(vd, args.apply)


if __name__ == "__main__":
    main()
