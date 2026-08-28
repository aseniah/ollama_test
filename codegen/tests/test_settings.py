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
models = [
  { name = "m:1", think = false, enabled = true },
  { name = "m:2", think = true,  enabled = true, infer_timeout = 300 },
  { name = "m:3", think = false, enabled = false },
  { name = "m:4", think = false, enabled = true, lmstudio_model = "m-4-lms" },
]

anthropic_models = [
  { alias = "haiku",  model_id = "claude-haiku-x", enabled = true },
  { alias = "sonnet", model_id = "claude-sonnet-x", enabled = false },
]

[harness]
default = "ollama"
ollama   = { base_url = "http://localhost:11434" }
lmstudio = { base_url = "http://localhost:1234" }
apple    = { base_url = "http://localhost:11435", autostart = true }

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

    def test_apple_autostart(self) -> None:
        self.assertTrue(self.s.apple_autostart())

    def test_languages(self) -> None:
        self.assertEqual(self.s.languages(), ["python", "go"])

    def test_local_models_enabled_only(self) -> None:
        names = [m["name"] for m in self.s.local_models()]
        self.assertEqual(names, ["m:1", "m:2", "m:4"])

    def test_local_models_defaults_applied(self) -> None:
        m1 = self.s.local_models()[0]
        self.assertEqual(m1["infer_timeout"], 120)
        self.assertEqual(m1["exec_timeout"], 60)
        self.assertEqual(m1["options"], {"think": False})

    def test_local_models_per_model_override(self) -> None:
        m2 = self.s.local_models()[1]
        self.assertEqual(m2["infer_timeout"], 300)
        self.assertEqual(m2["options"], {"think": True})

    def test_lmstudio_model_defaults_to_name(self) -> None:
        m1 = self.s.local_models()[0]
        self.assertEqual(m1["lmstudio_model"], "m:1")

    def test_lmstudio_model_override(self) -> None:
        m4 = self.s.local_models()[2]
        self.assertEqual(m4["lmstudio_model"], "m-4-lms")

    def test_anthropic_models_enabled_only(self) -> None:
        aliases = [a["alias"] for a in self.s.anthropic_models()]
        self.assertEqual(aliases, ["haiku"])

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
        bad = MINIMAL.replace('default = "ollama"', 'default = "bogus"')
        with self.assertRaises(settings.SettingsError):
            settings.load_settings(_write(bad))

    def test_harness_without_base_url(self) -> None:
        bad = MINIMAL.replace('ollama   = { base_url = "http://localhost:11434" }', 'ollama   = { }')
        with self.assertRaises(settings.SettingsError):
            settings.load_settings(_write(bad))

    def test_no_enabled_models(self) -> None:
        bad = MINIMAL.replace("enabled = true", "enabled = false")
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
