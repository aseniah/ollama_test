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
    # Test cases: (quantity, price, expected_result, description)
    test_cases = [
        (9, 100.0, 900.0, "No discount (quantity < 10)"),
        (10, 100.0, 900.0, "10% discount (quantity 10-49)"),
        (49, 100.0, 4410.0, "10% discount (quantity 49)"),
        (50, 100.0, 3750.0, "25% discount (quantity >= 50)")
    ]
    
    # Epsilon for float comparison
    EPS = 1e-9
    
    for qty, price, expected, desc in test_cases:
        actual = calculate_discount(price, qty)
        if abs(actual - expected) < EPS:
            print(f"PASS: {desc}")
        else:
            print(f"FAIL: {desc} (expected: {expected}, got: {actual})")


if __name__ == "__main__":
    run_tests()