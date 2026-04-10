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
    test_cases = [
        # (price, quantity, expected_result, description)
        (10.0, 9, 90.0, "quantity 9: no discount"),
        (10.0, 10, 90.0, "quantity 10: 10% off (boundary)"),
        (10.0, 49, 441.0, "quantity 49: 10% off (boundary)"),
        (10.0, 50, 375.0, "quantity 50: 25% off (boundary)"),
        (20.0, 60, 900.0, "quantity 60: 25% off"),
    ]

    for price, quantity, expected, description in test_cases:
        result = calculate_discount(price, quantity)
        # Use a small epsilon for floating point comparison safety, though exact matches expected here
        if abs(result - expected) < 1e-9:
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {result})")

if __name__ == "__main__":
    run_tests()