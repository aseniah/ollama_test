import re


def verify(
    stdout: str,
    stderr: str,
    exit_code: int | None,
    language: str,
    code: str,
) -> dict[str, object]:
    checks: dict[str, bool] = {
        "valid_format": False,
        "top_word": False,
        "it_count": False,
        "of_count": False,
        "ties_sorted": False,
    }

    lines = [line.strip() for line in stdout.strip().splitlines() if line.strip()]
    if not lines:
        return {"checks": checks, "passed": False}

    checks["valid_format"] = all(re.match(r"^\w+: \d+$", line) for line in lines)
    checks["top_word"] = lines[0] == "the: 11"

    checks["it_count"] = "it: 10" in lines
    checks["of_count"] = "of: 10" in lines

    # Among words with count 10, order must be: it, of, was
    count10 = [line.split(":")[0] for line in lines if line.endswith(": 10")]
    if "it" in count10 and "of" in count10 and "was" in count10:
        checks["ties_sorted"] = (
            count10.index("it") < count10.index("of") < count10.index("was")
        )

    passed = exit_code == 0 and all(checks.values())
    return {"checks": checks, "passed": passed}
