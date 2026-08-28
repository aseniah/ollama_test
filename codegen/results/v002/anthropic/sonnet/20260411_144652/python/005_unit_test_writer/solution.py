def calculate_discount(price: float, quantity: int) -> float:
    if quantity < 10:
        return price * quantity
    elif quantity < 50:
        return price * quantity * 0.9
    else:
        return price * quantity * 0.75


def run_test(description: str, expected: float, actual: float) -> None:
    if abs(expected - actual) < 1e-9:
        print(f"PASS: {description}")
    else:
        print(f"FAIL: {description} (expected: {expected}, got: {actual})")


run_test("quantity 9 - no discount", 10.0 * 9, calculate_discount(10.0, 9))
run_test("quantity 10 - 10% discount", 10.0 * 10 * 0.9, calculate_discount(10.0, 10))
run_test("quantity 49 - 10% discount", 10.0 * 49 * 0.9, calculate_discount(10.0, 49))
run_test("quantity 50 - 25% discount", 10.0 * 50 * 0.75, calculate_discount(10.0, 50))
run_test("quantity 1 - no discount", 5.0 * 1, calculate_discount(5.0, 1))
run_test("quantity 100 - 25% discount", 20.0 * 100 * 0.75, calculate_discount(20.0, 100))
