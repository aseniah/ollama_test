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


def test_calculate_discount():
    test_cases = [
        (10.0, 9, 90.0, "boundary: quantity 9 (no discount)"),
        (10.0, 10, 90.0, "boundary: quantity 10 (10% off)"),
        (10.0, 49, 441.0, "boundary: quantity 49 (10% off)"),
        (10.0, 50, 375.0, "boundary: quantity 50 (25% off)"),
        (5.0, 5, 25.0, "quantity 5 (no discount)"),
        (20.0, 100, 1500.0, "quantity 100 (25% off)"),
    ]

    passed = 0
    failed = 0

    for price, quantity, expected, description in test_cases:
        result = calculate_discount(price, quantity)
        if result == expected:
            print(f"PASS: {description}")
            passed += 1
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {result})")
            failed += 1

    print(f"\nTotal: {passed} passed, {failed} failed")


if __name__ == "__main__":
    test_calculate_discount()
