EXPECTED = [
    "1", "2", "Fizz", "4", "Buzz",
    "Fizz", "7", "8", "Fizz", "Buzz",
    "11", "Fizz", "13", "14", "FizzBuzz",
]


def verify(
    stdout: str,
    stderr: str,
    exit_code: int | None,
    language: str,
    code: str,
) -> dict[str, object]:
    checks: dict[str, bool] = {
        "correct_sequence": False,
        "first_line_is_1": False,
        "last_line_is_fizzbuzz": False,
    }

    lines = [line.strip() for line in stdout.strip().splitlines() if line.strip()]

    checks["correct_sequence"] = lines == EXPECTED
    checks["first_line_is_1"] = bool(lines) and lines[0] == "1"
    checks["last_line_is_fizzbuzz"] = len(lines) >= 15 and lines[14] == "FizzBuzz"

    passed = exit_code == 0 and all(checks.values())
    return {"checks": checks, "passed": passed}
