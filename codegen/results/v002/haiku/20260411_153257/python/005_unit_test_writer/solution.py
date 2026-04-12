def calculate_discount(price: float, quantity: int) -> float:
    if quantity < 10:
        return price * quantity
    elif quantity < 50:
        return price * quantity * 0.9
    else:
        return price * quantity * 0.75


def run_test(description: str, result: float, expected: float) -> None:
    if abs(result - expected) < 1e-9:
        print(f"PASS: {description}")
    else:
        print(f"FAIL: {description} (expected: {expected}, got: {result})")


tests = [
    ("quantity 9 - no discount",       calculate_discount(10.0, 9),  10.0 * 9),
    ("quantity 10 - 10% discount",     calculate_discount(10.0, 10), 10.0 * 10 * 0.9),
    ("quantity 49 - 10% discount",     calculate_discount(10.0, 49), 10.0 * 49 * 0.9),
    ("quantity 50 - 25% discount",     calculate_discount(10.0, 50), 10.0 * 50 * 0.75),
    ("quantity 1 - no discount",       calculate_discount(5.0, 1),   5.0 * 1),
    ("quantity 100 - 25% discount",    calculate_discount(20.0, 100), 20.0 * 100 * 0.75),
]

for description, result, expected in tests:
    run_test(description, result, expected)
