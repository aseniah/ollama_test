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
    # Test data: (price, quantity, expected_result, description)
    test_cases = [
        (10.0, 9, 90.0, "No discount for quantity < 10 (boundary 9)"),
        (10.0, 10, 90.0, "10% discount for quantity 10-49 (boundary 10)"),
        (10.0, 49, 441.0, "10% discount for quantity 10-49 (boundary 49)"),
        (10.0, 50, 375.0, "25% discount for quantity >= 50 (boundary 50)"),
        (20.0, 5, 100.0, "Standard small order calculation"),
        (20.0, 100, 1500.0, "Standard large order calculation"),
    ]

    for price, quantity, expected, description in test_cases:
        actual = calculate_discount(price, quantity)
        
        # Use round() or a small epsilon for float comparison to avoid precision issues
        if abs(actual - expected) < 1e-9:
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {actual})")

if __name__ == "__main__":
    run_tests()