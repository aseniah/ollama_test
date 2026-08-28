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
        r = b.generate([{"role": "user", "content": "hi"}], {"think": False}, 120, "m:1")
        self.assertEqual(r["response"], "print('x')")  # stripped
        self.assertEqual(r["eval_count"], 40)
        self.assertEqual(r["tok_per_sec"], 20.0)
        self.assertGreaterEqual(r["ms"], 0)
        self.assertEqual(http.last_url, "http://x:11434/api/chat")
        assert http.last_payload is not None
        self.assertEqual(http.last_payload["model"], "m:1")
        self.assertFalse(http.last_payload["stream"])
        self.assertFalse(http.last_payload["think"])


class LMStudioBackendTests(unittest.TestCase):
    def test_generate_prefers_v0_stats(self) -> None:
        http = _FakeHTTP(LMSTUDIO_V0_BODY)
        b = backends.LMStudioBackend("http://x:1234", _post=http)
        r = b.generate([{"role": "user", "content": "hi"}], {"think": True}, 120, "m:1")
        self.assertEqual(r["response"], "print('x')")
        self.assertEqual(r["tok_per_sec"], 33.3)
        self.assertEqual(r["eval_count"], 50)
        self.assertEqual(http.last_url, "http://x:1234/api/v0/chat/completions")
        assert http.last_payload is not None
        self.assertEqual(http.last_payload["model"], "m:1")
        self.assertNotIn("think", http.last_payload)  # not an OpenAI-compat param


class FactoryTests(unittest.TestCase):
    def _settings(self) -> settings.Settings:
        return settings.Settings({
            "harness": {
                "default": "ollama",
                "ollama": {"base_url": "http://o:11434"},
                "lmstudio": {"base_url": "http://l:1234"},
                "apple": {"base_url": "http://a:11435"},
            },
            "defaults": {"languages": ["python"], "infer_timeout": 120, "exec_timeout": 60},
            "models": [{"name": "m", "enabled": True}],
        })

    def test_build_local_backend_names(self) -> None:
        s = self._settings()
        self.assertIsInstance(backends.build_local_backend("ollama", s), backends.OllamaBackend)
        self.assertIsInstance(backends.build_local_backend("lmstudio", s), backends.LMStudioBackend)
        with self.assertRaises(backends.BackendError):
            backends.build_local_backend("apple", s)


if __name__ == "__main__":
    unittest.main()
