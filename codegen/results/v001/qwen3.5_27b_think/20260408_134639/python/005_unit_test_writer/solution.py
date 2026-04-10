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
    # Test cases defined as tuples: (price, quantity, expected_result, description)
    # Using price = 10.0 for easy calculation verification
    test_cases = [
        (10.0, 9, 90.0, "Quantity 9 (No discount)"),
        (10.0, 10, 90.0, "Quantity 10 (10% off boundary)"),
        (10.0, 49, 441.0, "Quantity 49 (10% off boundary)"),
        (10.0, 50, 375.0, "Quantity 50 (25% off boundary)"),
    ]

    for price, quantity, expected, description in test_cases:
        result = calculate_discount(price, quantity)
        
        # Use a small epsilon for floating point comparison safety
        if abs(result - expected) < 1e-9:
            print(f"PASS: {description}")
        else:
            # Format floats to 2 decimal places for clean output
            print(f"FAIL: {description} (expected: {expected:.2f}, got: {result:.2f})")


if __name__ == "__main__":
    run_tests()