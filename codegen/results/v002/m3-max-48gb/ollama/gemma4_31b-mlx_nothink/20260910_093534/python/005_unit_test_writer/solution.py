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
    # Test cases defined as: (description, price, quantity, expected_result)
    test_cases = [
        ("Quantity 9 (No discount boundary)", 100.0, 9, 900.0),
        ("Quantity 10 (10% discount start boundary)", 100.0, 10, 900.0),
        ("Quantity 49 (10% discount upper boundary)", 100.0, 49, 4410.0),
        ("Quantity 50 (25% discount start boundary)", 100.0, 50, 3750.0),
        ("Quantity 0 (Edge case)", 100.0, 0, 0.0),
    ]

    for description, price, quantity, expected in test_cases:
        actual = calculate_discount(price, quantity)
        
        # Using round() to avoid floating point precision issues during comparison
        if round(actual, 2) == round(expected, 2):
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {actual})")

if __name__ == "__main__":
    run_tests()