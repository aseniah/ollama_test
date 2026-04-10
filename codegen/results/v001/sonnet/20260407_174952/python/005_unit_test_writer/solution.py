def calculate_discount(price: float, quantity: int) -> float:
    """
    Calculate total price after bulk discount.
    quantity < 10: no discount
    quantity 10-49: 10% off
    quantity >= 50: 25% off
    Returns: price * quantity * discount_multiplier
    """
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


# Boundary: quantity 9 — no discount
run_test("quantity=9, no discount", 10.0 * 9, calculate_discount(10.0, 9))

# Boundary: quantity 10 — 10% off kicks in
run_test("quantity=10, 10% discount", 10.0 * 10 * 0.9, calculate_discount(10.0, 10))

# Boundary: quantity 49 — still 10% off
run_test("quantity=49, 10% discount", 10.0 * 49 * 0.9, calculate_discount(10.0, 49))

# Boundary: quantity 50 — 25% off kicks in
run_test("quantity=50, 25% discount", 10.0 * 50 * 0.75, calculate_discount(10.0, 50))

# Below lower boundary: quantity 1 — no discount
run_test("quantity=1, no discount", 5.0 * 1, calculate_discount(5.0, 1))

# Well above upper boundary: quantity 100 — 25% off
run_test("quantity=100, 25% discount", 2.0 * 100 * 0.75, calculate_discount(2.0, 100))
