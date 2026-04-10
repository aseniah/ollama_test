# iTerm2 Ollama Backend Benchmark — Findings

**Last updated:** 2026-03-30
**Goal:** Identify the best local model backend for the iTerm2 AI plugin for everyday shell command generation. Evaluation criteria: response cleanliness (no markdown/fences), macOS BSD correctness, speed (tok/s), and reliability. Ollama models are the primary focus; Apple's on-device model (via apfel) is also evaluated as an alternative backend. If a model cannot support a prompt or command category due to content guardrails, context limits, or other constraints, that is recorded as a disqualifying data point.

---

## Machines Tested

| | Secondary | Work | Primary |
|---|---|---|---|
| Machine | MacBook Air M3 | MacBook Pro M3 Pro | MacBook Pro M3 Max |
| Unified memory | 16GB | 36GB | 48GB |
| GPU inference | 100% via Ollama | 100% via Ollama | 100% via Ollama |
| Dates tested | 2026-03-26 | 2026-03-30 | 2026-03-28 to 2026-03-30 |

---

## Models Evaluated

### qwen3:4b — Eliminated early

Removed before full benchmarking. The default `qwen3:4b` tag on Ollama now points to thinking-only weights (`qwen3:4b-thinking-2507-q4_K_M`). Thinking mode cannot be disabled via the `think: false` API parameter — the model ignores it and runs full chain-of-thought regardless, producing 1-3 minute response times and verbose, unusable output for iTerm2 purposes.

The fix would be pulling `qwen3:4b-instruct-2507-q4_K_M` specifically, but given qwen3.5:4b works correctly and is a newer model, it's not worth pursuing.

### qwen3.5:4b — Benchmarked, not recommended

- `think: false` partially works but leaks `</think>` tokens into output on some responses
- M3 Air: ~1.3-1.7s responses, ~17-18 tok/s
- M3 Max: ~636ms responses, ~51 tok/s — roughly 2x slower than qwen2.5-coder:7b and qwen3-coder:30b on the same machine
- Accuracy is inconsistent across variants; BSD prompting helps but the model is less stable than the coder models
- Dropped from final benchmarking rounds — not a contender given the speed and consistency gap

### qwen2.5-coder:7b — Recommended for 16GB and 36GB

- No thinking mode, straightforward inference
- M3 Air: ~1.1s responses, ~20-21 tok/s
- M3 Max: ~304ms responses, ~77 tok/s
- Formatting reliable with prompt v2+
- Had a persistent recurring bug: `grep -r "TODO" *.js` instead of `grep -r "TODO" --include="*.js" .` — fixed with the v8 system prompt by adding an explicit `--include` instruction
- All known BSD flag issues resolved with the v8 prompt

### qwen3-coder:30b-64k — Recommended for 48GB

- MoE architecture: 30B total parameters, ~3.3B active per token at inference
- Not viable on 16GB (requires ~18-20GB to hold all expert weights resident in memory)
- M3 Max: ~313ms responses, ~86 tok/s — essentially tied with qwen2.5-coder:7b on wall time, higher tok/s
- All known BSD flag issues resolved with the v8 system prompt
- Consistently appends `-exec ls -l {} \;` to find executable commands (correct BSD, just verbose — not worth fixing)

### mistral:latest — Tested (not recommended)

- ~1.2s responses on M3 Air, ~21 tok/s
- Wraps responses in single backticks on most runs (8 of 15 prompts), making it unreliable for direct terminal pasting
- Commands are plausible but inconsistent across runs

### llama3.2:latest — Tested (not recommended)

- ~600ms responses on M3 Air, ~44 tok/s — fastest model tested on 16GB
- 13/15 clean output — better than mistral but still has formatting failures
- Accuracy is unreliable: `find . -name "*.js" | grep -r TODO` (invalid — can't pipe filenames to `grep -r`), inconsistent git log format strings, one tar run hardcoding `~/Documents/`

---

## Benchmark Results

### M3 Air (16GB) — prompt v1/v2/v3, 6 runs total

| | qwen3.5:4b | qwen2.5-coder:7b | mistral:latest | llama3.2:latest |
|---|---|---|---|---|
| Disk size | 3.4GB | 4.7GB | ~4.1GB | ~2.0GB |
| Avg response (warmed) | ~1.7s | ~1.1s | ~1.2s | ~0.6s |
| tok/s | ~17-18 | ~20-21 | ~21 | ~44 |
| Clean output | 45/45 | 43/45 | 5/15 | 13/15 |
| Accuracy | Inconsistent | Reliable | Inconsistent | Unreliable |
| Recommendation | Viable | **Best on 16GB** | Not recommended | Not recommended |

### M3 Max (48GB) — prompt v8, 3 runs, 24 prompts per model

| | qwen2.5-coder:7b | qwen3-coder:30b-64k |
|---|---|---|
| Avg response (warmed) | ~304ms | ~313ms |
| tok/s | ~77 | ~86 |
| Clean output | 24/24 | 24/24 |
| Accuracy | Strong | **Strongest** |
| Recommendation | Strong | **Best on 48GB** |

### Speed uplift: M3 Air → M3 Max

| Model | 16GB tok/s | 48GB tok/s | Uplift |
|---|---|---|---|
| qwen3.5:4b | ~18 | ~51 | +183% |
| qwen2.5-coder:7b | ~21 | ~77 | +267% |
| qwen3-coder:30b-64k | not viable | ~86 | — |

All models crossed the "feels instant" threshold on the M3 Max. The MoE architecture of qwen3-coder is particularly well-suited to this machine — it matches qwen2.5-coder's response time despite operating on 30B parameter weights.

---

## System Prompt Evolution

### v1 — iTerm2 default

```
Return a command suitable for copy/pasting into zsh on Darwin. Do NOT include
commentary NOR Markdown triple-backtick code blocks as your whole response will
be copied into my terminal automatically.

It must do this: {task}
```

### v2 — Basic rewrite

Resolved formatting failures (qwen2.5-coder wrapped output in markdown fences on 2 of 5 prompts with v1). Switched to explicit positive framing.

```
Output exactly one shell command for zsh on macOS. Your entire response must be
the raw command only. Do not use backticks, code fences, or any markdown. Do not
explain or add commentary. Just the command, nothing else.

It must do this: {task}
```

### v3 — New test prompts (same system prompt as v2)

Replaced the 5 test prompts with tasks that have tighter canonical answers, making accuracy easier to evaluate by inspection.

### v4 — BSD/GNU distinction added

Both models returned GNU-only flags (`-executable`, `/u=x`) when prompted with "macOS". Adding an explicit BSD/GNU line helped but wasn't enough for the most stubborn cases.

### v5 — Explicit flag examples

Added concrete examples of GNU flags that fail on macOS. The abstract "avoid GNU-only flags" instruction was insufficient — models have a strong training prior for flags like `-executable`. Naming the flag explicitly helped.

### v6–v8 — Variant testing

Ran A/B/C/D variant tests across prompt versions 5–8 to find the minimum effective prompt. Key findings:

- **Principle-based language alone doesn't work** for flags with strong training priors (`-executable` in particular). Models need flag names and alternatives.
- **The grep `--include` instruction is load-bearing** for qwen2.5-coder:7b. Without it, the model consistently returns `grep -r "TODO" *.js` (globs current directory only) regardless of how the BSD line is worded.
- **Removing the grep instruction destabilizes qwen3-coder** — it reverts to broken find-pipe-tar patterns.
- **Compressed prompts perform comparably** to verbose ones, with minor noise (qwen2.5-coder adds `| grep -v 'COMMAND'` to ps output). The verbose form is more stable.

### v8 — Final recommended prompt

The result of all variant testing. Fixes every known failure mode for both qwen2.5-coder:7b and qwen3-coder:30b across 3 runs × 8 test prompts.

**Benchmark version:**
```
Output exactly one shell command for zsh on macOS.
macOS uses BSD userland, not GNU coreutils.
GNU-only flags that do NOT work on macOS: -executable for find (use -perm +111),
--sort for ps (pipe to sort -k instead), --max-depth for du (use -d instead),
symbolic perm notation like /u=x (use octal instead).
To sort processes by memory on macOS, use: ps aux | sort -k4nr | head -n N.
To filter grep by file type recursively, use --include='*.ext' with a path
(e.g. grep -r 'x' --include='*.js' .), not shell globs (grep -r 'x' *.js is wrong).
Your entire response must be the raw command only. Do not use backticks, code fences,
or any markdown. Do not explain or add commentary. Just the command, nothing else.
```

**iTerm2 production prompt:**
```
Output exactly one shell command for \(shell) on macOS.
macOS uses BSD userland, not GNU coreutils.
GNU-only flags that do NOT work on macOS: -executable for find (use -perm +111),
--sort for ps (pipe to sort -k instead), --max-depth for du (use -d instead),
symbolic perm notation like /u=x (use octal instead).
To sort processes by memory on macOS, use: ps aux | sort -k4nr | head -n N.
To filter grep by file type recursively, use --include='*.ext' with a path
(e.g. grep -r 'x' --include='*.js' .), not shell globs (grep -r 'x' *.js is wrong).
Your entire response must be the raw command only. Do not use backticks, code fences,
or any markdown. Do not explain or add commentary. Just the command, nothing else.

It must do this: \(ai.prompt)
```

`\(shell)` expands to `zsh` at runtime. `\(uname)` was considered but expands to `Darwin` rather than `macOS` — hardcoding `macOS` in the opening line is clearer for the model.

---

## BSD Flag Findings

Both models defaulted to GNU-style syntax for `find . -type f` with execute permission checks:

| Prompt | Model | Response | Result |
|---|---|---|---|
| v2 | qwen3-coder:30b | `find . -type f -executable` | Fails — GNU-only flag |
| v2 | qwen2.5-coder:7b | `find . -type f -perm /u=x` | Fails — GNU symbolic notation |
| v8 | qwen2.5-coder:7b | `find . -perm +111` | Correct |
| v8 | qwen3-coder:30b | `find . -type f -perm +111 -exec ls -l {} \;` | Correct (verbose) |
| Canonical | — | `find . -type f -perm +111` | Correct for macOS BSD find |

`-perm /mode` is specified in POSIX.1-2008 but Apple's BSD `find` does not support it — produces `illegal mode string`. `-perm +111` is the correct form on macOS.

`ps --sort` (GNU procps) was also a persistent failure for qwen3-coder until the explicit ps idiom was added in v7/v8. `sed -i ''` (BSD in-place edit) and `stat -f %z` (BSD file size) were handled correctly by both models without explicit instruction.

---

## Test Prompt Set (v4–v8)

1. show the last 10 git commits with hash and message, one per line
2. show the top 5 processes by memory usage
3. create a tar.gz of the current directory named backup.tar.gz, excluding .git and node_modules
4. search recursively for the string TODO in all .js files in the current directory
5. show disk usage of each subdirectory in the current directory, sorted by size
6. find all executable files in this directory and its subdirectories
7. replace all occurrences of 'foo' with 'bar' in place in a file named file.txt
8. print the size in bytes of a file named data.txt

Prompts 7 and 8 (sed in-place, stat file size) were added in v7 to test two classic BSD/GNU divergence points. Both were handled correctly by qwen2.5-coder and qwen3-coder without explicit instruction, confirming the v8 system prompt generalises beyond the flags it names.

---

## Setup Notes

### Warmup is required

The first inference request after model load has a significant spike (16s observed for qwen3.5 in early runs). The benchmark sends a throwaway prompt after preloading to settle the model before timed tests begin. Any production integration should do the same — or accept that the very first iTerm2 request after a cold start will be slow.

### Preloading

Use the Ollama `/api/chat` endpoint with an empty messages array to load the model into GPU memory without running inference:

```json
POST /api/chat
{ "model": "qwen2.5-coder:7b", "messages": [], "keep_alive": -1 }
```

### Unloading

```json
POST /api/chat
{ "model": "qwen2.5-coder:7b", "messages": [], "keep_alive": 0 }
```

---

## Results Storage

Results are stored in `results/v{NNN}.jsonl` — one JSON record per line, versioned by prompt version. Key fields:

```json
{
  "run_id": "20260326_143022",
  "timestamp": "2026-03-26T14:30:22+00:00",
  "prompt_v": 8,
  "prompt_variant": "A",
  "model": "qwen2.5-coder:7b",
  "model_options": {},
  "prompt_label": "recursive grep",
  "task": "search recursively for the string TODO in all .js files in the current directory",
  "response": "grep -r 'TODO' --include='*.js' .",
  "ms": 287,
  "eval_count": 12,
  "tok_per_sec": 77.4,
  "clean": true
}
```

Filter by `prompt_v` to compare only runs using the same prompt version. Filter by `prompt_variant` to isolate a specific system prompt. Earlier results (prompt versions 1–5) are archived in `results/v001-005.jsonl`.

---

## Clean Detector Bug (fixed)

The original `is_clean()` only caught triple-fence markdown (` ``` `). Both mistral and llama3.2 frequently wrap responses in single backticks (e.g. `` `command` ``), which passed the check incorrectly. Fixed in `benchmark.py` and all historical records were recomputed.

---

## iTerm2 Configuration

### AI Features Available

iTerm2 has more AI capabilities than just command generation:

| Feature | How to invoke | What it does |
|---|---|---|
| **Command Generator** | `Cmd+Y` | Natural language → shell command. The feature tested in this benchmark. |
| **Codecierge** | `View > Toolbelt > Codecierge` | Describe a multi-step task; watches your terminal as you work and updates guidance in real time. |
| **AI Chat** | `Shift+Ctrl+Cmd+Y` | Full chat window, optionally linked to your terminal session. |
| **Explain Output** | `Edit > Explain Output with AI` | Select terminal output; opens a linked chat that annotates it with explanations. |
| **Engage AI** | `Edit > Engage Artificial Intelligence` | Passes the text at your cursor to the model as a prompt. |

When AI Chat is linked to a session, it can be granted permissions: read terminal state, view command history, read man pages, run commands, send keystrokes, write to the filesystem, and write to clipboard. The write-to-filesystem permission enables script generation directly from the chat.

### Model: context size variants

The qwen3-coder:30b weights support a native 262k context window. Context-size variants are created via Modelfile with a `num_ctx` override — the weights are identical, only the KV cache allocation differs.

| Tag | num_ctx | KV cache (est.) | Use case |
|---|---|---|---|
| `qwen3-coder:30b-32k` | 32,768 | ~135MB | iTerm2 Cmd+Y, Codecierge |
| `qwen3-coder:30b-64k` | 65,536 | ~270MB | General iTerm2 default |
| `qwen3-coder:30b-128k` | 131,072 | ~540MB | AI Chat with file contents |

KV cache differences are negligible on 48GB. The `num_ctx` value only matters when the actual prompt approaches the limit — for Cmd+Y (50-100 tokens) any variant behaves identically.

### Recommended iTerm2 settings

| Machine | Model | Context limit | Response limit |
|---|---|---|---|
| 48GB M3 Max | `qwen3-coder:30b-128k` | 65,536 | 16,384 |
| 36GB M3 Pro | `qwen2.5-coder:7b` | 16,384 | 8,192 |
| 16GB M3 Air | `qwen2.5-coder:7b` | 16,384 | 8,192 |

The context limit is a cap on what iTerm2 *sends*, not what gets pre-allocated — so Cmd+Y (which sends ~100 tokens) is not slowed by a high context limit.

---

## Conclusion

**Recommendation depends on machine:**

- **48GB (M3 Max):** `qwen3-coder:30b-128k` with the v8 iTerm2 prompt. Fastest response times, highest tok/s, strongest accuracy. The MoE architecture delivers 30B-class knowledge at dense-3B inference cost.

- **36GB (M3 Pro):** `qwen2.5-coder:7b` — qwen3-coder:30b pushes memory on a developer workload. qwen2.5-coder:7b is lighter, fast, and fully accurate with the v8 prompt.

- **16GB (M3 Air):** `qwen2.5-coder:7b` — same recommendation as 36GB, slower but still usable (~1.1s responses).

The v8 system prompt resolved all known failure modes: BSD find flags, GNU ps options, recursive grep file filtering, sed in-place syntax, and stat file size. Both recommended models produce correct macOS commands consistently across all 8 test prompts.

Whether local Ollama models are "good enough" for iTerm2 depends on use case. For well-known commands (git, find, grep, ps, tar), both models are reliable with the v8 prompt. For complex or unusual requests, accuracy degrades and you may get plausible-looking but subtly wrong commands. A user who can spot a bad command will benefit; one who pastes blindly faces some risk.
