import unittest

import migrate_results as mr

ALIASES = {"haiku", "opus", "sonnet"}


class ClassifyTests(unittest.TestCase):
    def test_apple(self) -> None:
        self.assertEqual(mr.classify_dir("apple-foundationmodel", ALIASES), "apple")

    def test_anthropic_alias(self) -> None:
        self.assertEqual(mr.classify_dir("haiku", ALIASES), "anthropic")
        self.assertEqual(mr.classify_dir("sonnet", ALIASES), "anthropic")

    def test_local_model(self) -> None:
        self.assertEqual(mr.classify_dir("qwen3.5_27b_think", ALIASES), "ollama")
        self.assertEqual(mr.classify_dir("gemma4_31b_nothink", ALIASES), "ollama")

    def test_already_harness_dir_skipped(self) -> None:
        for h in ("ollama", "lmstudio", "apple", "anthropic"):
            self.assertIsNone(mr.classify_dir(h, ALIASES))


if __name__ == "__main__":
    unittest.main()
