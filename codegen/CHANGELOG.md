# Changelog

## v2 (2026-04-09)

### Impact on results

The bug fix prompt rewrite (006) should meaningfully improve scores for that test — v1 runs were largely polluted by models returning the FizzBuzz sequence as plain text rather than code, which is a harness misunderstanding rather than a capability failure. Expect most models to do better here in v2.

The vacuous check fixes (001, 004, 008) will modestly lower scores for models that produce empty output due to compile failures. Previously those models received partial credit for checks that passed vacuously — now they won't. The effect is largest for languages with high compile failure rates (Go, C#).

The C# LANG_NOTE update gives models explicit guidance on CSV and JSON parsing without NuGet. This should reduce compile failures on tests that involve those formats, raising C# scores across the board.

### Test changes

- `006_bug_fix/prompt.md`: removed "output the corrected sequence" phrasing that caused models to return the FizzBuzz sequence as plain text instead of returning code
- `008_prime_numbers/verify.py`: `one_excluded` and `no_composite` now require non-empty output — previously both passed vacuously on empty stdout from compile errors
- `001_csv_to_json/verify.py`: `age_is_int` and `score_is_float` now require non-empty output — same vacuous pass pattern as above
- `004_json_filter/verify.py`: `bob_excluded` and `carol_excluded` now require non-empty output — same vacuous pass pattern as above

### Harness changes

- `benchmark.py` / `run_claude_test.py`: `verify_error` flag in exception handler corrected to `True` (was `False`)
- `benchmark.py` LANG_NOTE C#: added CSV parsing guidance (`File.ReadAllLines()` / `string.Split(',')`); "no NuGet packages" moved earlier in the note
- `run_claude_test.py`: removed `--gen-ms` argument; `ms` is now stored as `null` for all Claude API runs. Generation time is not captured — Claude API latency includes network and subagent overhead, making it incomparable to Ollama inference times. Pass/fail is the primary signal for Claude comparisons.
- `run_claude_test.py` / `README.md`: Claude API orchestration methodology updated. Subagents now only return code — the orchestrator extracts it and calls `run_claude_test.py`. Previously, subagents were instructed to run the verifier themselves, which allowed them to see failures and iterate (invalidating first-shot measurement). The new approach enforces a single generation attempt per test.

## v1 (2026-04-05)

Initial release.

- 8 tests across 4 languages: Python, TypeScript, Go, C#
- Prompt variant C: natural style with explicit stdout hint
- Ollama backend with preload/warmup/unload lifecycle
- Apple on-device model support via apfel backend
- Claude API runner (Sonnet, Haiku, Opus) via subagent tool use
- Artifact writing per run: solution, stdout, stderr, result.json
- Results appended to `results/v{NNN}/{model}/results.jsonl` (all backends)
