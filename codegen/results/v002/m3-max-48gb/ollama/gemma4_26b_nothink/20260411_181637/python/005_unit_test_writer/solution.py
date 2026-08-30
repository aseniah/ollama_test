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
    Runs test cases for the calculate_discount function.
    Tests boundary conditions: 9, 10, 49, 50.
    """
    # Define test cases as a list of tuples:
    # (price, quantity, expected_output, description)
    test_cases = [
        (100.0, 9, 900.0, "Boundary: quantity 9 (No discount)"),
        (100.0, 10, 900.0, "Boundary: quantity 10 (10% discount)"),
        (100.0, 49, 4410.0, "Boundary: quantity 49 (10% discount)"),
        (100.0, 50, 3750.0, "Boundary: quantity 50 (25% discount)"),
        (50.0, 100, 3750.0, "Case: quantity 100 (25% discount)"),
        (20.0, 1, 20.0, "Case: quantity 1 (No discount)")
    ]

    for price, quantity, expected, description in test_cases:
        try:
            actual = calculate_discount(price, quantity)
            
            # Using round() to handle potential floating point precision issues
            # in comparison (e.g., 4410.000000000001 vs 4410.0)
            if round(actual, 2) == round(expected, 2):
                print(f"PASS: {description}")
            else:
                print(f"FAIL: {description} (expected: {expected}, got: {actual})")
        except Exception as e:
            print(f"FAIL: {description} (Error: {str(e)})")

if __name__ == "__main__":
    run_tests()