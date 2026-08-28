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
    # Test cases with expected results
    test_cases = [
        # (price, quantity, expected_result, description)
        (10.0, 9, 90.0, "quantity 9 (below threshold, no discount)"),
        (10.0, 10, 90.0, "quantity 10 (threshold, 10% off)"),
        (10.0, 49, 441.0, "quantity 49 (upper threshold for 10% off)"),
        (10.0, 50, 375.0, "quantity 50 (threshold for 25% off)"),
        (20.0, 5, 100.0, "quantity 5 with price 20 (no discount)"),
        (20.0, 25, 450.0, "quantity 25 with price 20 (10% off)"),
        (20.0, 100, 1500.0, "quantity 100 with price 20 (25% off)"),
    ]
    
    for price, quantity, expected, description in test_cases:
        result = calculate_discount(price, quantity)
        # Use a small tolerance for floating point comparison
        if abs(result - expected) < 0.001:
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {result})")


if __name__ == "__main__":
    test_calculate_discount()