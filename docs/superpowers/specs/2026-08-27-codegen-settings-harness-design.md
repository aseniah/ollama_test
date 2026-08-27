# codegen: settings file, LM Studio harness, harness-based results layout

**Date:** 2026-08-27
**Status:** Approved for planning
**Scope:** `codegen/` only. `shell/` is untouched.

## Context

The codegen benchmark configures models by commenting/uncommenting the inline
`MODELS` list in `codegen/benchmark.py`. Adding a new runtime (LM Studio) or a
new model means editing Python. Results are written flat under
`results/v{NNN}/{model}/`, with Apple and Anthropic model dirs sitting as peers
of the Ollama model dirs, so there is no way to run the *same* model under two
runtimes and compare.

This change:

1. Moves model + runtime config into `codegen/settings.toml`.
2. Adds LM Studio as an optional local harness alongside Ollama.
3. Restructures results to `results/v{NNN}/{harness}/{model}/` so Ollama,
   LM Studio, Apple, and Anthropic results are separated by runtime and the
   same model can be compared across harnesses.
4. Surfaces token-rate data (already recorded per-record) in the end-of-run
   summary and the findings report.

`PROMPT_VERSION` stays at `2` — tests and prompt variant C are unchanged, so
records remain comparable. New JSONL fields are additive and backward-compatible.

## Decisions

| Question | Decision |
|---|---|
| Scope | `codegen/` only |
| Config format | TOML with inline tables, parsed via stdlib `tomllib` |
| Model selection | `enabled = true/false` per entry; a run uses all enabled entries |
| Harness selection | `--harness {ollama,lmstudio}` CLI flag, default from settings; one local harness per invocation |
| Apple | stays a separate `--apple` toggle, can be added to any run |
| Anthropic | folded into the restructure — `run_claude_test.py` reads models from settings, writes under `anthropic/` |
| Existing results | left in place; new layout going forward; one-time `migrate_results.py` |
| Version bump | none |

## Non-goals (YAGNI)

- `shell/` changes.
- `settings.local.toml` override layer.
- Model profiles / named groups.
- Auto-starting the LM Studio server.
- Mixing two local harnesses in one invocation.
- Rewriting `apfel_backend.py` at the repo root (shared with `shell/`).

## Architecture

### New files (all in `codegen/`)

| File | Purpose |
|---|---|
| `settings.toml` | Committed config data. Replaces the inline `MODELS` list. |
| `settings.py` | Loader + validation. Stdlib only. |
| `backends.py` | `Backend` protocol + `OllamaBackend`, `LMStudioBackend`, `AppleBackend`. |
| `migrate_results.py` | One-time, manually run, migrates existing `results/vNNN/*` into the harness layout. |

### `settings.toml`

```toml
[harness]
default = "ollama"
ollama   = { base_url = "http://localhost:11434" }
lmstudio = { base_url = "http://localhost:1234" }
apple    = { base_url = "http://localhost:11435", autostart = true }

[defaults]
infer_timeout = 120
exec_timeout  = 60
languages = ["python", "typescript", "go", "csharp"]

models = [
  { name = "qwen3.8:27b", think = false, enabled = true },
  { name = "qwen3.8:27b", think = true,  enabled = true, infer_timeout = 300 },
  { name = "gemma4:31b",  think = false, enabled = false },
]

anthropic_models = [
  { alias = "haiku",  model_id = "claude-haiku-4-5-20251001", enabled = true },
  { alias = "sonnet", model_id = "claude-sonnet-5",           enabled = false },
]
```

Per-model optional keys: `infer_timeout`, `exec_timeout` (override `[defaults]`),
`lmstudio_model` (the model key to send to LM Studio when it differs from the
Ollama tag; defaults to `name`).

`model_safe` — the results directory segment — always derives from the canonical
`name` (`:` and `/` replaced with `_`, plus `_think` / `_nothink` suffix), so the
same model under `ollama/` and `lmstudio/` produces matching directory names.

### `settings.py`

```python
def load_settings(path: Path = Path("settings.toml")) -> Settings
```

- Parses with `tomllib`.
- Validates: `[harness].default` is a known harness; each referenced harness has
  a `base_url`; at least one enabled entry in `models`; `languages` non-empty.
- Applies `[defaults]` to each model entry that does not override.
- Accessors:
  - `Settings.local_models() -> list[ModelConfig]` — enabled `models`, defaults applied.
  - `Settings.anthropic_models() -> list[AnthropicModel]` — enabled `anthropic_models`.
  - `Settings.harness_base_url(name: str) -> str`
  - `Settings.harness_apple() -> AppleHarnessConfig` (base_url, autostart)
  - `Settings.default_harness() -> str`
  - `Settings.languages() -> list[str]`

`ModelConfig` keeps its current `TypedDict` shape (`name`, `options`,
`infer_timeout`, `exec_timeout`) plus an internal `lmstudio_model`. `options` is
built from `think` (`{"think": bool}`) to preserve the existing payload shape and
the `_think` / `_nothink` results-dir suffix logic.

### `backends.py`

```python
class GenResult(TypedDict):
    response: str
    ms: int
    eval_count: int
    tok_per_sec: float

class Backend(Protocol):
    name: str                                   # "ollama" | "lmstudio" | "apple"
    def start(self) -> None: ...                # health check or process spawn
    def stop(self) -> None: ...                 # teardown; no-op for ollama/lmstudio
    def warmup(self, messages: list[dict]) -> None: ...
    def generate(self, messages: list[dict], options: dict, timeout: int) -> GenResult: ...
```

- **`OllamaBackend(base_url)`** — `POST /api/chat`, `stream: false`, `**options`
  merged into the payload (unchanged from today). Timing: server-side
  `eval_count / (eval_duration / 1e9)`. `start()` pings `/api/tags`.
- **`LMStudioBackend(base_url, model_key_fn)`** — prefers
  `POST /api/v0/chat/completions` (LM Studio REST API, returns
  `stats.tokens_per_second`, `stats.tokens`); falls back to
  `POST /v1/chat/completions` with wall-clock timing and
  `usage.completion_tokens`. `options` mapping: `think` is dropped (not an
  OpenAI-compat param); other options passed through. `start()` pings
  `/v1/models` and raises a clear "LM Studio server not reachable at {base_url}"
  error if it is down (no autostart).
- **`AppleBackend(apfel_module, autostart)`** — thin adapter over the repo-root
  `apfel_backend`. `start()` → `ensure_running()` (stores the proc handle if it
  started one), `stop()` → `teardown(proc)` only if `start()` spawned it,
  `warmup()` → `apfel_backend.warmup()`, `generate()` → `apfel_backend.run_prompt()`
  reshaped into `GenResult`.

A `build_local_backend(name, settings) -> Backend` factory returns the Ollama or
LM Studio backend for the active `--harness`.

### `benchmark.py` changes

- **`main()`**
  - Add `--harness {ollama,lmstudio}`, default `settings.default_harness()`.
    Keep positional `runs` and `--apple`.
  - `settings = load_settings()`; `LANGUAGES = settings.languages()`.
  - `backend = build_local_backend(args.harness, settings)`; `backend.start()`.
  - `models = settings.local_models()`.
  - If `--apple`: `apple_backend = AppleBackend(...)`, `apple_backend.start()`,
    append the synthesized apple `ModelConfig` (`name = "apple-foundationmodel"`,
    `backend_name = "apple"`) to `models`.
  - `try/finally` calls `backend.stop()` and (if built) `apple_backend.stop()`.
- **Per-model loop** — resolve which backend object handles the entry
  (`apple` name → apple backend, else the local backend) and its harness name
  for the results path.
- **`run_one(backend, harness_name, model_cfg, variant, test, run_id, timestamp, artifact_base)`**
  - Replace the `is_apfel` branch (current `run_one`, the model-dispatch block)
    with `result = backend.generate(messages, model_cfg["options"], infer_timeout)`.
  - Record gains `"harness": harness_name`.
- **Results paths**
  - `save_record`: `RESULTS_DIR / f"v{PROMPT_VERSION:03d}" / harness_name / model_safe / "results.jsonl"`.
  - `artifact_base`: same prefix, then `run_id`.
- **`print_summary`** — keep the pass/fail x `ms` grid. After each model's grid,
  add one aggregate line: `mean {tok_per_sec} tok/s · mean gen ~{s}s` computed
  from that model's records in the run.

### `run_claude_test.py` changes

- Replace the module-level `MODELS` dict with a read of
  `settings.anthropic_models()` (build the same `alias -> model_id` mapping,
  plus the reverse lookup).
- `--model` default becomes the first enabled anthropic alias (fall back to
  `"haiku"` if the list is empty), instead of the current hard-coded `"sonnet"`.
- Results dir: `results/v{PROMPT_VERSION:03d}/anthropic/{friendly}/results.jsonl`.
  `--results-dir` override still bypasses this.
- Record gains `"harness": "anthropic"`.

### `migrate_results.py`

- CLI: `python3 migrate_results.py [VERSION] [--apply] [--stamp-harness]`.
  Default version is `v{PROMPT_VERSION:03d}`. Dry-run unless `--apply`.
- For each immediate child directory of `results/{version}/`:
  - Skip if the name is already `ollama` / `lmstudio` / `apple` / `anthropic`.
  - `apple-foundationmodel` → `apple/`.
  - Name matches an alias in `settings.anthropic_models()` (or the historical
    set `haiku`, `opus`, `sonnet`) → `anthropic/`.
  - Everything else → `ollama/` (all historical local runs were Ollama).
- Moves use `git mv` (results are tracked). Prints the planned moves in dry-run.
- `--stamp-harness`: after moving, inject `"harness": "<dir>"` into every JSONL
  record under that harness dir that lacks the field.
- Not imported or called by `benchmark.py`. Run once, by hand, on a clean tree.

### Results interpretation updates

- **`codegen/CLAUDE.md` → "Analyzing Results"**
  - Every jq path `results/v002/*/results.jsonl` → `results/v002/*/*/results.jsonl`.
  - Add harness-scoped examples: `results/v002/ollama/*/results.jsonl`,
    `results/v002/lmstudio/*/results.jsonl`,
    `results/v002/anthropic/*/results.jsonl`.
  - Add a `tok_per_sec` aggregation query (mean/median per model).
  - Document the new `harness` field.
- **`codegen/CLAUDE.md` → "Structure"** — add `settings.toml`, `settings.py`,
  `backends.py`, `migrate_results.py`; update the results path to
  `results/v{NNN}/{harness}/{model}/…`.
- **`codegen/CLAUDE.md` → "Running"** — document `--harness` and `settings.toml`.
- **`codegen/results/findings_instructions.md`**
  - "Before You Start": list `results/v{NNN}/{harness}/` subdirs and note which
    harnesses ran.
  - New optional section **Harness Comparison (Ollama vs LM Studio)**: for any
    model run under both, compare score (expect near-parity — a large gap is a
    harness bug, not a model finding), speed, and tok/s. Note that tok/s is
    sourced differently per harness (Ollama `eval_duration`, LM Studio
    `stats.tokens_per_second` or wall-clock fallback, Apple wall-clock) and is
    not perfectly comparable.
  - "Speed vs. Quality": tok/s column becomes mandatory; add a Harness column.
- **`codegen/README.md`** — `settings.toml` overview, `--harness` flag, LM Studio
  setup (start server, load the model, default port 1234, no autostart).

## Data flow

```
settings.toml
   │  load_settings()
   ▼
Settings ──────────────┬───────────────────────────┐
   │                   │                           │
   │ local_models()    │ anthropic_models()        │ harness_base_url()
   ▼                   ▼                           ▼
benchmark.py       run_claude_test.py         backends.py
   │                   │                           │
   │ build_local_backend(--harness)                │
   ▼                                               │
Backend.generate() ◄───────────────────────────────┘
   │
   ▼
record {..., "harness": <name>}
   │
   ▼
results/v002/<harness>/<model_safe>/results.jsonl
   │
   ▼
CLAUDE.md jq queries  +  findings_instructions.md  →  FINDINGS_V002.md
```

## Error handling

- `settings.py`: raise `SettingsError` with a specific message on missing file,
  TOML parse error, unknown default harness, harness without `base_url`, zero
  enabled models, empty `languages`. `benchmark.py` and `run_claude_test.py`
  print it and `sys.exit(1)`.
- `LMStudioBackend.start()`: unreachable server → `RuntimeError` naming the
  base URL and telling the user to start LM Studio and load a model.
- `AppleBackend`: unchanged apfel failure behavior (already exits on start
  failure).
- `migrate_results.py`: refuse to run with a dirty `results/` tree unless
  `--apply` is combined with an explicit `--force`; always dry-run first.

## Testing

Unit (new `codegen/tests/` entries — mirror existing test style):

- `settings.py`: enabled filter, `[defaults]` application, per-model override,
  `lmstudio_model` fallback to `name`, each validation error path.
- `backends.py`: `GenResult` shape for each backend against recorded HTTP
  fixtures (Ollama `/api/chat`, LM Studio `/api/v0` and `/v1` fallback);
  `model_safe` derivation.
- `migrate_results.py`: classification of `apple-foundationmodel`, `haiku`,
  `qwen3.5_27b_think`, and an already-migrated `ollama` dir, on a temp tree.

End-to-end (manual, needs the runtimes):

- `python3 benchmark.py 1` → records in `results/v002/ollama/qwen3.8_27b_nothink/`.
- `python3 benchmark.py 1 --harness lmstudio` (LM Studio up) →
  `results/v002/lmstudio/qwen3.8_27b_nothink/`.
- `python3 benchmark.py 1 --apple` → `results/v002/apple/apple-foundationmodel/`,
  apfel process starts and stops as before.
- `python3 benchmark.py 1 --harness lmstudio` with LM Studio down → clean error,
  no partial results dir.
- One `run_claude_test.py` cell → `results/v002/anthropic/haiku/`.
- `bash lint.sh` clean.
- `python3 migrate_results.py --apply` on a git-clean tree → `git status` shows
  only renames; updated jq queries from CLAUDE.md return sane numbers.

## Rollout order

1. `settings.py` + `settings.toml` + tests (no behavior change yet).
2. `backends.py` + tests.
3. `benchmark.py` switched to settings + backends + harness path segment.
4. `run_claude_test.py` switched to settings + `anthropic/` path.
5. Doc updates (`CLAUDE.md`, `findings_instructions.md`, `README.md`).
6. `migrate_results.py` + tests; run it manually when ready.
