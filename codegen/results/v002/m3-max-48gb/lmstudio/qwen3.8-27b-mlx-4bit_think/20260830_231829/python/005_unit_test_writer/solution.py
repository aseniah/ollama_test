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


def run_tests():
    tolerance = 1e-9

    test_cases = [
        (10.0, 9,  10.0 * 9 * 1.0,  "quantity 9 (no discount, boundary)"),
        (10.0, 10, 10.0 * 10 * 0.9, "quantity 10 (10% discount, boundary)"),
        (10.0, 49, 10.0 * 49 * 0.9, "quantity 49 (10% discount, boundary)"),
        (10.0, 50, 10.0 * 50 * 0.75,"quantity 50 (25% discount, boundary)"),
    ]

    for price, quantity, expected, description in test_cases:
        got = calculate_discount(price, quantity)
        if abs(got - expected) < tolerance:
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {got}")


if __name__ == "__main__":
    run_tests()