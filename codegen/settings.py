"""Load and validate codegen/settings.toml."""

from __future__ import annotations

import tomllib
from pathlib import Path
from typing import Any, TypedDict, cast

_HARNESSES = ("ollama", "lmstudio", "apple")
_LOCAL_HARNESSES = ("ollama", "lmstudio")

_Table = dict[str, Any]


class SettingsError(Exception):
    """Raised for any malformed or missing settings file."""


class ModelEntry(TypedDict):
    name: str
    options: dict[str, Any]
    infer_timeout: int
    exec_timeout: int
    max_tokens: int


class AnthropicEntry(TypedDict):
    alias: str
    model_id: str


def _rows(raw: _Table, key: str) -> list[_Table]:
    return cast("list[_Table]", raw.get(key, []))


class Settings:
    def __init__(self, raw: _Table) -> None:
        self._raw = raw

    def _harness(self, name: str) -> _Table:
        return cast(_Table, self._raw["harness"][name])

    def default_harness(self) -> str:
        return str(self._raw["harness"]["default"])

    def harness_base_url(self, name: str) -> str:
        return str(self._harness(name)["base_url"])

    def harness_autostart(self, name: str) -> bool:
        return bool(self._harness(name).get("autostart", False))

    def languages(self) -> list[str]:
        return [str(x) for x in cast("list[Any]", self._raw["defaults"]["languages"])]

    def local_models(self, harness: str) -> list[ModelEntry]:
        d = cast(_Table, self._raw["defaults"])
        out: list[ModelEntry] = []
        for m in _rows(self._harness(harness), "models"):
            if not m.get("enabled", False):
                continue
            out.append(ModelEntry(
                name=str(m["name"]),
                options={"think": bool(m.get("think", False))},
                infer_timeout=int(m.get("infer_timeout", d["infer_timeout"])),
                exec_timeout=int(m.get("exec_timeout", d["exec_timeout"])),
                max_tokens=int(m.get("max_tokens", d.get("max_tokens", 4096))),
            ))
        return out

    def anthropic_models(self) -> list[AnthropicEntry]:
        return [
            AnthropicEntry(alias=str(a["alias"]), model_id=str(a["model_id"]))
            for a in _rows(self._raw, "anthropic_models")
            if a.get("enabled", False)
        ]

    def anthropic_default_alias(self) -> str:
        enabled = self.anthropic_models()
        if enabled:
            return enabled[0]["alias"]
        return "haiku"


def _validate(raw: _Table) -> None:
    harness = cast(_Table, raw.get("harness", {}))
    if "default" not in harness:
        raise SettingsError("settings.toml: missing [harness].default")
    default = harness["default"]
    if default not in _LOCAL_HARNESSES:
        raise SettingsError(
            f"settings.toml: [harness].default = {default!r} is not one of {_LOCAL_HARNESSES}"
        )
    for h in _HARNESSES:
        entry = harness.get(h)
        if not isinstance(entry, dict) or "base_url" not in entry:
            raise SettingsError(f"settings.toml: [harness.{h}] needs a base_url")
    for h in _LOCAL_HARNESSES:
        models = cast(_Table, harness[h]).get("models")
        if not isinstance(models, list):
            raise SettingsError(f"settings.toml: [harness.{h}] needs a models list")
    default_models = cast("list[Any]", cast(_Table, harness[default])["models"])
    if not any(cast(_Table, m).get("enabled", False) for m in default_models):
        raise SettingsError(
            f"settings.toml: no enabled models for the default harness [{default}] — "
            "set enabled = true on at least one, or change [harness].default"
        )
    defaults = cast(_Table, raw.get("defaults", {}))
    if not defaults:
        raise SettingsError("settings.toml: missing [defaults]")
    langs = defaults.get("languages")
    if not isinstance(langs, list) or not langs:
        raise SettingsError("settings.toml: [defaults].languages must be a non-empty list")
    for key in ("infer_timeout", "exec_timeout"):
        if key not in defaults:
            raise SettingsError(f"settings.toml: [defaults].{key} is required")


def load_settings(path: Path = Path("settings.toml")) -> Settings:
    try:
        text = path.read_text()
    except FileNotFoundError as e:
        raise SettingsError(f"settings.toml not found at {path}") from e
    try:
        raw = tomllib.loads(text)
    except tomllib.TOMLDecodeError as e:
        raise SettingsError(f"settings.toml is not valid TOML: {e}") from e
    _validate(raw)
    return Settings(raw)
