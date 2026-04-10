import json
import re
from typing import Any, cast

FRAMEWORK_KEYWORDS = [
    "unittest", "assert", "describe", "test_",
    "it(", "xunit", "NUnit", "testing.T",
]


def verify(
    stdout: str,
    stderr: str,
    exit_code: int | None,
    language: str,
    code: str,
) -> dict[str, object]:
    checks: dict[str, bool] = {
        "valid_json": False,
        "four_records": False,
        "correct_first_names": False,
        "correct_last_names": False,
        "birthday_format": False,
        "age_john": False,
        "age_paul": False,
        "age_ringo": False,
        "age_george": False,
        "null_relatives_excluded": False,
        "relative_count_george": False,
    }

    spontaneous_tests = any(kw in code for kw in FRAMEWORK_KEYWORDS)

    try:
        parsed: Any = json.loads(stdout.strip())
    except Exception:
        return {"checks": checks, "passed": False, "spontaneous_tests": spontaneous_tests}

    if not isinstance(parsed, list):
        return {"checks": checks, "passed": False, "spontaneous_tests": spontaneous_tests}

    raw_list: list[Any] = cast(list[Any], parsed)
    data: list[dict[str, Any]] = [
        cast(dict[str, Any], item) for item in raw_list if isinstance(item, dict)
    ]
    checks["valid_json"] = True
    checks["four_records"] = len(data) == 4

    first_names = [r.get("FirstName") for r in data]
    last_names = [r.get("LastName") for r in data]
    checks["correct_first_names"] = all(
        n in first_names for n in ["John", "James", "Ringo", "George"]
    )
    checks["correct_last_names"] = all(
        n in last_names for n in ["Lennon", "McCartney", "Starr", "Harrison"]
    )

    birthday_pattern = re.compile(r"^\d{4}-\d{2}-\d{2}$")
    checks["birthday_format"] = all(
        birthday_pattern.match(str(r.get("Birthday", ""))) for r in data
    )

    def find_by_last(last: str) -> dict[str, Any] | None:
        for rec in data:
            if rec.get("LastName") == last:
                return rec
        return None

    john = find_by_last("Lennon")
    paul = find_by_last("McCartney")
    ringo = find_by_last("Starr")
    george = find_by_last("Harrison")

    checks["age_john"] = john is not None and john.get("Age") == 40
    checks["age_paul"] = paul is not None and paul.get("Age") == 83
    checks["age_ringo"] = ringo is not None and ringo.get("Age") == 84
    checks["age_george"] = george is not None and george.get("Age") == 58

    if john is not None:
        john_relatives: list[dict[str, Any]] = john.get("Relatives") or []
        john_relationships = [r.get("Relationship") for r in john_relatives]
        checks["null_relatives_excluded"] = (
            "Brother" not in john_relationships and "Sister" not in john_relationships
        )

    if george is not None:
        george_relatives: list[dict[str, Any]] = george.get("Relatives") or []
        checks["relative_count_george"] = len(george_relatives) == 4

    passed = exit_code == 0 and all(checks.values())
    return {"checks": checks, "passed": passed, "spontaneous_tests": spontaneous_tests}
