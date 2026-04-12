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
    """
    Executes a suite of test cases for calculate_discount and prints results.
    """
    # Test cases definition: (description, price, quantity, expected_result)
    test_cases = [
        ("Boundary condition: quantity 9 (no discount)", 100.0, 9, 900.0),
        ("Boundary condition: quantity 10 (10% discount)", 100.0, 10, 900.0),
        ("Boundary condition: quantity 49 (10% discount)", 100.0, 49, 4410.0),
        ("Boundary condition: quantity 50 (25% discount)", 100.0, 50, 3750.0),
        ("Standard case: quantity 25 (10% discount)", 20.0, 25, 450.0),
        ("Standard case: quantity 100 (25% discount)", 10.0, 100, 750.0)
    ]

    for description, price, quantity, expected in test_cases:
        actual = calculate_discount(price, quantity)
        
        # Using round to handle potential floating point precision issues during comparison
        # although for these specific decimals (0.9, 0.75), standard equality usually works.
        if round(actual, 10) == round(expected, 10):
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {actual})")

if __name__ == "__main__":
    run_tests()