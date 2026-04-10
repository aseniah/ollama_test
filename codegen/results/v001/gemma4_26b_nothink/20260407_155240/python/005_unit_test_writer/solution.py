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
    Runs test cases for the calculate_discount function and prints 
    the results to stdout.
    """
    # Test cases definition: (price, quantity, expected_output, description)
    test_cases = [
        (100.0, 9, 900.0, "Boundary: quantity 9 (No discount)"),
        (100.0, 10, 900.0, "Boundary: quantity 10 (10% discount)"),
        (100.0, 49, 4410.0, "Boundary: quantity 49 (10% discount)"),
        (100.0, 50, 3750.0, "Boundary: quantity 50 (25% discount)"),
        (50.0, 20, 900.0, "Standard: quantity 20 (10% discount)"),
        (10.0, 100, 750.0, "Standard: quantity 100 (25% discount)")
    ]

    for price, quantity, expected, description in test_cases:
        # We use round() to handle floating point precision issues common in float math
        actual = round(calculate_discount(price, quantity), 2)
        expected_rounded = round(expected, 2)

        if actual == expected_rounded:
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected_rounded}, got: {actual})")

if __name__ == "__main__":
    run_tests()