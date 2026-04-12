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
    # List of test cases: (price, quantity, expected_value, description)
    test_cases = [
        (10.0, 9, 90.0, "Quantity 9 (No Discount)"),
        (10.0, 10, 90.0, "Quantity 10 (Start of 10% Discount)"),
        (10.0, 49, 441.0, "Quantity 49 (End of 10% Discount)"),
        (10.0, 50, 375.0, "Quantity 50 (Start of 25% Discount)")
    ]

    for price, quantity, expected, description in test_cases:
        result = calculate_discount(price, quantity)
        
        # Use a small tolerance for floating point comparison
        if abs(result - expected) < 1e-9:
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {result})")


if __name__ == "__main__":
    run_tests()