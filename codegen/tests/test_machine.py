import tempfile
import unittest
from pathlib import Path
from typing import Any

import machine

SPEC: machine.MachineSpec = {
    "chip": "Apple M3 Max",
    "model_identifier": "Mac15,9",
    "model_name": "MacBook Pro",
    "cpu_cores": 16,
    "cpu_cores_performance": 12,
    "cpu_cores_efficiency": 4,
    "memory_gb": 48,
    "arch": "arm64",
    "os": "26.6.2",
    "os_build": "25G83",
    "python": "3.14.7",
    "hostname": "host.local",
    "ollama_version": "0.33.1",
    "lms_version": None,
}


class ParsingTests(unittest.TestCase):
    def test_parse_processors(self) -> None:
        self.assertEqual(machine._parse_processors("proc 16:12:4:0"), (16, 12, 4))
        self.assertEqual(machine._parse_processors("proc 8"), (8, 0, 0))
        self.assertEqual(machine._parse_processors(""), (0, 0, 0))

    def test_semver(self) -> None:
        self.assertEqual(machine._semver("client version is 0.33.1"), "0.33.1")
        self.assertIsNone(machine._semver("CLI commit: 07b7252"))

    def test_suggest_slug(self) -> None:
        self.assertEqual(machine.suggest_slug(SPEC), "m3-max-48gb")
        no_mem = dict(SPEC)
        no_mem["memory_gb"] = 0
        self.assertEqual(machine.suggest_slug(no_mem), "m3-max")  # type: ignore[arg-type]


class FingerprintTests(unittest.TestCase):
    def test_stable_and_12_hex(self) -> None:
        fp = machine.fingerprint(SPEC)
        self.assertEqual(len(fp), 12)
        self.assertEqual(fp, machine.fingerprint(dict(SPEC)))  # type: ignore[arg-type]

    def test_os_and_python_do_not_change_fingerprint(self) -> None:
        drifted = dict(SPEC)
        drifted["os"] = "27.0"
        drifted["python"] = "3.15.0"
        drifted["ollama_version"] = "9.9.9"
        self.assertEqual(machine.fingerprint(drifted), machine.fingerprint(SPEC))  # type: ignore[arg-type]

    def test_hardware_change_changes_fingerprint(self) -> None:
        bigger = dict(SPEC)
        bigger["memory_gb"] = 128
        self.assertNotEqual(machine.fingerprint(bigger), machine.fingerprint(SPEC))  # type: ignore[arg-type]


class RegistryTests(unittest.TestCase):
    def setUp(self) -> None:
        self.dir = Path(tempfile.mkdtemp())

    def test_register_then_load(self) -> None:
        path = machine.register("m3-max-48gb", SPEC, self.dir)
        self.assertTrue(path.exists())
        entries = machine.load_registry(self.dir)
        self.assertEqual(len(entries), 1)
        self.assertEqual(entries[0]["slug"], "m3-max-48gb")
        self.assertEqual(entries[0]["fingerprint"], machine.fingerprint(SPEC))
        self.assertEqual(entries[0]["specs"]["memory_gb"], 48)

    def test_register_rejects_bad_slug(self) -> None:
        for bad in ("Has Space", "UPPER", "-lead", "trailing-"):
            with self.assertRaises(machine.MachineError):
                machine.register(bad, SPEC, self.dir)

    def test_load_registry_missing_dir(self) -> None:
        self.assertEqual(machine.load_registry(self.dir / "nope"), [])

    def test_none_values_omitted_from_toml(self) -> None:
        text = machine.register("m3-max-48gb", SPEC, self.dir).read_text()
        self.assertNotIn("lms_version", text)
        self.assertIn("ollama_version", text)


class ResolveTests(unittest.TestCase):
    def setUp(self) -> None:
        self.dir = Path(tempfile.mkdtemp())
        self._real_detect = machine.detect
        machine.detect = lambda: SPEC  # type: ignore[assignment]

    def tearDown(self) -> None:
        machine.detect = self._real_detect  # type: ignore[assignment]

    def test_known_fingerprint_returns_slug(self) -> None:
        machine.register("box-a", SPEC, self.dir)
        self.assertEqual(machine.resolve(machines_dir=self.dir, interactive=False), "box-a")

    def test_unknown_non_interactive_raises(self) -> None:
        with self.assertRaises(machine.MachineError):
            machine.resolve(machines_dir=self.dir, interactive=False)

    def test_unknown_with_register_slug_writes_entry(self) -> None:
        slug = machine.resolve(register_slug="box-b", machines_dir=self.dir, interactive=False)
        self.assertEqual(slug, "box-b")
        self.assertTrue((self.dir / "box-b.toml").exists())

    def test_register_slug_conflicting_with_known_raises(self) -> None:
        machine.register("box-a", SPEC, self.dir)
        with self.assertRaises(machine.MachineError):
            machine.resolve(register_slug="box-c", machines_dir=self.dir, interactive=False)

    def test_snapshot_shape(self) -> None:
        snap: dict[str, Any] = dict(machine.snapshot("box-a", SPEC))
        self.assertEqual(snap["slug"], "box-a")
        self.assertEqual(snap["fingerprint"], machine.fingerprint(SPEC))
        self.assertEqual(snap["specs"]["chip"], "Apple M3 Max")


if __name__ == "__main__":
    unittest.main()
