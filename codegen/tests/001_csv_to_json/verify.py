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
        "row_count": False,
        "age_is_int": False,
        "score_is_float": False,
        "correct_first_name": False,
    }

    try:
        data = json.loads(stdout.strip())
    except Exception:
        return {"checks": checks, "passed": False}

    checks["valid_json"] = isinstance(data, list)
    if not checks["valid_json"]:
        return {"checks": checks, "passed": False}

    checks["row_count"] = len(data) == 5
    if data:
        checks["age_is_int"] = all(isinstance(r.get("Age"), int) for r in data)
        checks["score_is_float"] = all(isinstance(r.get("Score"), (int, float)) for r in data)
        checks["correct_first_name"] = data[0].get("Name") == "Alice Johnson"

    passed = exit_code == 0 and all(checks.values())
    return {"checks": checks, "passed": passed}
