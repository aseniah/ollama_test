import math

EXPECTED = ["2", "3", "5", "7", "11", "13", "17", "19", "23", "29", "31", "37", "41", "43", "47"]


def is_prime(n: int) -> bool:
    if n < 2:
        return False
    for i in range(2, int(math.sqrt(n)) + 1):
        if n % i == 0:
            return False
    return True


def verify(
    stdout: str,
    stderr: str,
    exit_code: int | None,
    language: str,
    code: str,
) -> dict[str, object]:
    checks: dict[str, bool] = {
        "correct_sequence": False,
        "one_excluded": False,
        "no_composite": False,
    }

    lines = [line.strip() for line in stdout.strip().splitlines() if line.strip()]

    checks["correct_sequence"] = lines == EXPECTED
    if lines:
        checks["one_excluded"] = "1" not in lines
        try:
            checks["no_composite"] = all(is_prime(int(line)) for line in lines if line.isdigit())
        except ValueError:
            checks["no_composite"] = False

    passed = exit_code == 0 and all(checks.values())
    return {"checks": checks, "passed": passed}
