import json


def verify(
    stdout: str,
    stderr: str,
    exit_code: int | None,
    language: str,
    code: str,
) -> dict[str, object]:
    checks: dict[str, bool] = {
        "valid_json": False,
        "correct_count": False,
        "correct_order": False,
        "bob_excluded": False,
        "carol_excluded": False,
    }

    try:
        data = json.loads(stdout.strip())
    except Exception:
        return {"checks": checks, "passed": False}

    checks["valid_json"] = isinstance(data, list)
    if not checks["valid_json"]:
        return {"checks": checks, "passed": False}

    names = [r.get("name") for r in data]
    checks["correct_count"] = len(data) == 4
    checks["correct_order"] = names == ["Alice", "David", "Eva", "Grace"]
    if data:
        checks["bob_excluded"] = "Bob" not in names
        checks["carol_excluded"] = "Carol" not in names

    passed = exit_code == 0 and all(checks.values())
    return {"checks": checks, "passed": passed}
