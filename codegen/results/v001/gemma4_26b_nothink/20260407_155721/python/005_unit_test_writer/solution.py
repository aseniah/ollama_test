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
    Runs a suite of test cases against the calculate_discount function.
    Each test case defines a description, price, quantity, and expected result.
    """
    # List of test cases: (description, price, quantity, expected_result)
    test_cases = [
        ("Boundary: quantity 9 (No discount)", 100.0, 9, 900.0),
        ("Boundary: quantity 10 (10% discount applied)", 100.0, 10, 900.0),
        ("Boundary: quantity 49 (10% discount applied)", 100.0, 49, 4410.0),
        ("Boundary: quantity 50 (25% discount applied)", 100.0, 50, 3750.0),
        ("General case: quantity 100 (25% discount applied)", 10.0, 100, 750.0),
    ]

    for description, price, quantity, expected in test_cases:
        # Calculate actual result using the function
        actual = calculate_discount(price, quantity)
        
        # Using round to handle potential floating point precision issues in comparison
        # This ensures that 4410.000000000001 is treated as 4410.0
        if round(actual, 10) == round(expected, 10):
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {actual})")


if __name__ == "__main__":
    run_tests()