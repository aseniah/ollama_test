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
    # Test cases defined as: (price, quantity, expected_result, description)
    test_cases = [
        (10.0, 9, 90.0, "Boundary: Just below 10 (No discount)"),
        (10.0, 10, 90.0, "Boundary: Exactly 10 (10% discount)"),
        (10.0, 49, 441.0, "Boundary: Just below 50 (10% discount)"),
        (10.0, 50, 375.0, "Boundary: Exactly 50 (25% discount)"),
        (10.0, 5, 50.0, "Standard: Small quantity"),
        (10.0, 100, 750.0, "Standard: Large quantity"),
    ]

    for price, qty, expected, desc in test_cases:
        actual = calculate_discount(price, qty)
        
        # Using round() or abs() difference to avoid floating point precision issues
        if abs(actual - expected) < 1e-9:
            print(f"PASS: {desc}")
        else:
            print(f"FAIL: {desc} (expected: {expected}, got: {actual})")

if __name__ == "__main__":
    run_tests()