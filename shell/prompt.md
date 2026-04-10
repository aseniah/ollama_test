Write a shell script that benchmarks three local Ollama models for use as an iTerm2 AI plugin backend.

The models are: qwen3:4b, qwen3.5:4b, qwen2.5-coder:7b

For each model:
1. Pull it into memory by sending a warmup request first
2. Send 5 test prompts via the Ollama REST API (localhost:11434) that simulate iTerm2 command generation — things like finding files, bulk renaming, ffmpeg conversion, killing a process on a port, and a moderately complex pipeline
3. Unload the model from memory before moving to the next one

For each model + prompt, capture:
- The response text
- Response time in ms
- Whether the output is a clean command (no markdown fences, no explanation text)

Output a formatted comparison table at the end. Save full responses to a file for manual review.

Important:
- Do NOT execute any of the generated commands — capture responses only
- Test models sequentially, never concurrently
- Use the Ollama API's keep_alive: 0 to explicitly unload each model after its tests complete
