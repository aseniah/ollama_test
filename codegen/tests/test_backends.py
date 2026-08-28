import unittest
from typing import Any

import backends
import settings


class _FakeHTTP:
    """Records the last request and returns a canned JSON body."""

    def __init__(self, body: dict[str, Any]) -> None:
        self.body = body
        self.last_url: str | None = None
        self.last_payload: dict[str, Any] | None = None

    def __call__(self, url: str, payload: dict[str, Any], timeout: int) -> dict[str, Any]:
        self.last_url = url
        self.last_payload = payload
        return self.body


OLLAMA_BODY: dict[str, Any] = {
    "message": {"content": "  print('x')  "},
    "eval_count": 40,
    "eval_duration": 2_000_000_000,  # 2s in ns -> 40 / 2 = 20 tok/s
}

LMSTUDIO_V0_BODY: dict[str, Any] = {
    "choices": [{"message": {"content": "print('x')"}}],
    "stats": {"tokens_per_second": 33.3, "predicted_tokens_count": 50},
    "usage": {"completion_tokens": 50},
}


class OllamaBackendTests(unittest.TestCase):
    def test_generate_shape_and_timing(self) -> None:
        http = _FakeHTTP(OLLAMA_BODY)
        b = backends.OllamaBackend("http://x:11434", _post=http)
        r = b.generate([{"role": "user", "content": "hi"}], {"think": False}, 120, "m:1", 4096)
        self.assertEqual(r["response"], "print('x')")  # stripped
        self.assertEqual(r["eval_count"], 40)
        self.assertEqual(r["tok_per_sec"], 20.0)
        self.assertGreaterEqual(r["ms"], 0)
        self.assertEqual(http.last_url, "http://x:11434/api/chat")
        assert http.last_payload is not None
        self.assertEqual(http.last_payload["model"], "m:1")
        self.assertFalse(http.last_payload["stream"])
        self.assertFalse(http.last_payload["think"])
        self.assertEqual(http.last_payload["options"], {"num_predict": 4096})

    def test_generate_nests_sampling_under_options(self) -> None:
        http = _FakeHTTP(OLLAMA_BODY)
        b = backends.OllamaBackend("http://x:11434", _post=http)
        b.generate(
            [{"role": "user", "content": "hi"}],
            {"think": False, "temperature": 1, "top_k": 20}, 120, "m:1", 4096,
        )
        assert http.last_payload is not None
        self.assertEqual(
            http.last_payload["options"],
            {"temperature": 1, "top_k": 20, "num_predict": 4096},
        )
        self.assertNotIn("temperature", http.last_payload)  # not top-level


class LMStudioBackendTests(unittest.TestCase):
    def test_generate_prefers_v0_stats(self) -> None:
        http = _FakeHTTP(LMSTUDIO_V0_BODY)
        b = backends.LMStudioBackend("http://x:1234", _post=http)
        r = b.generate([{"role": "user", "content": "hi"}], {"think": True}, 120, "m:1", 4096)
        self.assertEqual(r["response"], "print('x')")
        self.assertEqual(r["tok_per_sec"], 33.3)
        self.assertEqual(r["eval_count"], 50)
        self.assertEqual(http.last_url, "http://x:1234/api/v0/chat/completions")
        assert http.last_payload is not None
        self.assertEqual(http.last_payload["model"], "m:1")
        self.assertNotIn("think", http.last_payload)  # not an OpenAI-compat param
        self.assertEqual(http.last_payload["max_tokens"], 4096)
        self.assertEqual(http.last_payload["chat_template_kwargs"], {"enable_thinking": True})
        self.assertEqual(http.last_payload["reasoning_effort"], "medium")

    def test_generate_maps_nothink_to_reasoning_off(self) -> None:
        http = _FakeHTTP({
            "choices": [{"message": {"content": "x"}}],
            "usage": {"completion_tokens": 10, "completion_tokens_details": {"reasoning_tokens": 3}},
        })
        b = backends.LMStudioBackend("http://x:1234", _post=http)
        r = b.generate([{"role": "user", "content": "hi"}], {"think": False}, 120, "m:1", 4096)
        assert http.last_payload is not None
        self.assertEqual(http.last_payload["chat_template_kwargs"], {"enable_thinking": False})
        self.assertEqual(http.last_payload["reasoning_effort"], "none")
        self.assertEqual(r["reasoning_tokens"], 3)


class FactoryTests(unittest.TestCase):
    def _settings(self) -> settings.Settings:
        return settings.Settings({
            "harness": {
                "default": "ollama",
                "ollama": {"base_url": "http://o:11434", "models": [{"name": "m", "enabled": True}]},
                "lmstudio": {"base_url": "http://l:1234", "models": []},
                "apple": {"base_url": "http://a:11435"},
            },
            "defaults": {"languages": ["python"], "infer_timeout": 120, "exec_timeout": 60},
        })

    def test_build_local_backend_names(self) -> None:
        s = self._settings()
        self.assertIsInstance(backends.build_local_backend("ollama", s), backends.OllamaBackend)
        self.assertIsInstance(backends.build_local_backend("lmstudio", s), backends.LMStudioBackend)
        with self.assertRaises(backends.BackendError):
            backends.build_local_backend("apple", s)


class WarmupTests(unittest.TestCase):
    def test_ollama_warmup_pins_and_warms(self) -> None:
        calls: list[dict[str, Any]] = []

        def rec(url: str, payload: dict[str, Any], timeout: int) -> dict[str, Any]:
            calls.append(payload)
            return {"message": {"content": "ok"}}

        backends.OllamaBackend("http://x:11434", _post=rec).warmup(
            [{"role": "user", "content": "hi"}], {"think": True}, "m:1", 120
        )
        self.assertEqual(calls[0]["keep_alive"], -1)          # pinned
        self.assertEqual(calls[1]["model"], "m:1")
        self.assertFalse(calls[1]["think"])                   # thinking dropped for warmup

    def test_lmstudio_warmup_loads_then_pings_capped(self) -> None:
        cmds: list[list[str]] = []

        class _R:
            returncode = 0
            stderr = ""

        def fake_run(cmd: list[str], **_: object) -> _R:
            cmds.append(cmd)
            return _R()

        orig = backends._run
        backends._run = fake_run  # type: ignore[assignment]
        try:
            http = _FakeHTTP({"choices": [{"message": {"content": "ok"}}], "usage": {}})
            backends.LMStudioBackend("http://x:1234", _post=http).warmup(
                [{"role": "user", "content": "hi"}], {"think": True}, "m:1", 120
            )
        finally:
            backends._run = orig  # type: ignore[assignment]

        self.assertTrue(any("load" in c and "m:1" in c for c in cmds))  # lms load m:1
        assert http.last_payload is not None
        self.assertEqual(http.last_payload["max_tokens"], 16)          # tiny warmup ping


class UnloadTests(unittest.TestCase):
    def test_ollama_unload_posts_keep_alive_zero(self) -> None:
        http = _FakeHTTP({})
        backends.OllamaBackend("http://x:11434", _post=http).unload("m:1")
        assert http.last_payload is not None
        self.assertEqual(http.last_payload["keep_alive"], 0)
        self.assertEqual(http.last_payload["model"], "m:1")

    def test_ollama_unload_swallows_backend_error(self) -> None:
        def boom(url: str, payload: dict[str, Any], timeout: int) -> dict[str, Any]:
            raise backends.BackendError("down")
        backends.OllamaBackend("http://x:11434", _post=boom).unload("m:1")  # no raise


if __name__ == "__main__":
    unittest.main()
