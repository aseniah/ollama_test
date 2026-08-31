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


def run_test(description, price, quantity, expected):
    got = calculate_discount(price, quantity)
    # Use a small tolerance for floating point comparison
    if abs(got - expected) < 1e-9:
        print(f"PASS: {description}")
    else:
        print(f"FAIL: {description} (expected: {expected}, got: {got})")


if __name__ == "__main__":
    run_test("boundary: quantity 9 (no discount)", 10.0, 9, 90.0)
    run_test("boundary: quantity 10 (10% discount)", 10.0, 10, 90.0)
    run_test("boundary: quantity 49 (10% discount)", 10.0, 49, 441.0)
    run_test("boundary: quantity 50 (25% discount)", 10.0, 50, 375.0)
    run_test("quantity 1 (no discount)", 25.5, 1, 25.5)
    run_test("quantity 100 (25% discount)", 5.0, 100, 375.0)