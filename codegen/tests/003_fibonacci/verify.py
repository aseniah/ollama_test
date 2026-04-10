EXPECTED = ["1", "1", "2", "3", "5", "8", "13", "21", "34", "55", "89"]


def verify(
    stdout: str,
    stderr: str,
    exit_code: int | None,
    language: str,
    code: str,
) -> dict[str, object]:
    checks: dict[str, bool] = {
        "correct_sequence": False,
        "no_extra_output": False,
    }

    lines = [line.strip() for line in stdout.strip().splitlines() if line.strip()]

    checks["correct_sequence"] = lines[:len(EXPECTED)] == EXPECTED
    checks["no_extra_output"] = lines == EXPECTED

    passed = exit_code == 0 and all(checks.values())
    return {"checks": checks, "passed": passed}
