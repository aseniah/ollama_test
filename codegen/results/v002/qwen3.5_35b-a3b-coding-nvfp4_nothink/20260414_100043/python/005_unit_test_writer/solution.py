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
    """Run test cases for calculate_discount function."""
    test_cases = [
        # (price, quantity, expected_result, description)
        (100.0, 9, 900.0, "quantity < 10 (no discount)"),
        (100.0, 10, 900.0, "quantity = 10 (10% off)"),
        (100.0, 49, 4410.0, "quantity = 49 (10% off)"),
        (100.0, 50, 3750.0, "quantity = 50 (25% off)"),
    ]

    for price, quantity, expected, description in test_cases:
        result = calculate_discount(price, quantity)
        if abs(result - expected) < 1e-9:  # Floating point comparison
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {result})")


if __name__ == "__main__":
    run_tests()