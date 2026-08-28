import tempfile
import unittest
from pathlib import Path

import settings


def _write(text: str) -> Path:
    d = Path(tempfile.mkdtemp())
    p = d / "settings.toml"
    p.write_text(text)
    return p


MINIMAL = """
anthropic_models = [
  { alias = "haiku",  model_id = "claude-haiku-x", enabled = true },
  { alias = "sonnet", model_id = "claude-sonnet-x", enabled = false },
]

[harness]
default = "ollama"

[harness.ollama]
base_url = "http://localhost:11434"
models = [
  { name = "o:1", think = false, enabled = true },
  { name = "o:2", think = true,  enabled = true, infer_timeout = 300 },
  { name = "o:3", think = false, enabled = false },
]

[harness.lmstudio]
base_url = "http://localhost:1234"
models = [
  { name = "lms-1", think = false, enabled = true },
]

[harness.apple]
base_url = "http://localhost:11435"
autostart = true

[defaults]
infer_timeout = 120
exec_timeout  = 60
languages = ["python", "go"]
"""


class LoadTests(unittest.TestCase):
    def setUp(self) -> None:
        self.s = settings.load_settings(_write(MINIMAL))

    def test_default_harness(self) -> None:
        self.assertEqual(self.s.default_harness(), "ollama")

    def test_harness_base_url(self) -> None:
        self.assertEqual(self.s.harness_base_url("lmstudio"), "http://localhost:1234")

    def test_harness_autostart(self) -> None:
        self.assertTrue(self.s.harness_autostart("apple"))
        self.assertFalse(self.s.harness_autostart("ollama"))

    def test_languages(self) -> None:
        self.assertEqual(self.s.languages(), ["python", "go"])

    def test_local_models_per_harness(self) -> None:
        self.assertEqual([m["name"] for m in self.s.local_models("ollama")], ["o:1", "o:2"])
        self.assertEqual([m["name"] for m in self.s.local_models("lmstudio")], ["lms-1"])

    def test_local_models_defaults_applied(self) -> None:
        m1 = self.s.local_models("ollama")[0]
        self.assertEqual(m1["infer_timeout"], 120)
        self.assertEqual(m1["exec_timeout"], 60)
        self.assertEqual(m1["options"], {"think": False})

    def test_local_models_per_model_override(self) -> None:
        m2 = self.s.local_models("ollama")[1]
        self.assertEqual(m2["infer_timeout"], 300)
        self.assertEqual(m2["options"], {"think": True})

    def test_anthropic_models_enabled_only(self) -> None:
        self.assertEqual([a["alias"] for a in self.s.anthropic_models()], ["haiku"])

    def test_anthropic_default_alias(self) -> None:
        self.assertEqual(self.s.anthropic_default_alias(), "haiku")


class ValidationTests(unittest.TestCase):
    def test_missing_file(self) -> None:
        with self.assertRaises(settings.SettingsError):
            settings.load_settings(Path("/nonexistent/settings.toml"))

    def test_bad_toml(self) -> None:
        with self.assertRaises(settings.SettingsError):
            settings.load_settings(_write("this is = not valid toml ["))

    def test_unknown_default_harness(self) -> None:
        bad = MINIMAL.replace('default = "ollama"', 'default = "apple"')
        with self.assertRaises(settings.SettingsError):
            settings.load_settings(_write(bad))

    def test_harness_without_base_url(self) -> None:
        bad = MINIMAL.replace('base_url = "http://localhost:11434"\n', "")
        with self.assertRaises(settings.SettingsError):
            settings.load_settings(_write(bad))

    def test_no_enabled_models_for_default_harness(self) -> None:
        bad = MINIMAL.replace('{ name = "o:1", think = false, enabled = true }',
                              '{ name = "o:1", think = false, enabled = false }')
        bad = bad.replace('{ name = "o:2", think = true,  enabled = true, infer_timeout = 300 }',
                          '{ name = "o:2", think = true,  enabled = false, infer_timeout = 300 }')
        with self.assertRaises(settings.SettingsError):
            settings.load_settings(_write(bad))

    def test_empty_languages(self) -> None:
        bad = MINIMAL.replace('languages = ["python", "go"]', "languages = []")
        with self.assertRaises(settings.SettingsError):
            settings.load_settings(_write(bad))

    def test_anthropic_default_alias_fallback(self) -> None:
        bad = MINIMAL.replace(
            '{ alias = "haiku",  model_id = "claude-haiku-x", enabled = true }',
            '{ alias = "haiku",  model_id = "claude-haiku-x", enabled = false }',
        )
        s = settings.load_settings(_write(bad))
        self.assertEqual(s.anthropic_default_alias(), "haiku")


if __name__ == "__main__":
    unittest.main()
