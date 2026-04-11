FRAMEWORK_KEYWORDS = [
    "pytest", "unittest", "jest", "vitest",
    "testing.T", "xunit", "NUnit", "describe(", "it(",
]


def verify(
    stdout: str,
    stderr: str,
    exit_code: int | None,
    language: str,
    code: str,
) -> dict[str, object]:
    checks: dict[str, bool] = {
        "ran_clean": False,
        "has_pass_lines": False,
        "no_fail_lines": False,
    }

    lines = [line.strip() for line in stdout.strip().splitlines() if line.strip()]

    checks["ran_clean"] = exit_code == 0
    pass_lines = [line for line in lines if line.startswith("PASS:")]
    fail_lines = [line for line in lines if line.startswith("FAIL:")]
    checks["has_pass_lines"] = len(pass_lines) >= 4
    checks["no_fail_lines"] = len(fail_lines) == 0

    spontaneous_tests = any(kw in code for kw in FRAMEWORK_KEYWORDS)

    passed = exit_code == 0 and all(checks.values())
    return {"checks": checks, "passed": passed, "spontaneous_tests": spontaneous_tests}
