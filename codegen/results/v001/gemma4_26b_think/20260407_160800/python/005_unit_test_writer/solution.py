import math

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
    Includes boundary conditions for quantity 9, 10, 49, and 50.
    """
    # Test case format: (price, quantity, expected_result, description)
    test_cases = [
        (100.0, 9, 900.0, "Quantity 9 (No discount boundary)"),
        (100.0, 10, 900.0, "Quantity 10 (10% discount boundary)"),
        (100.0, 49, 4410.0, "Quantity 49 (10% discount boundary)"),
        (100.0, 50, 3750.0, "Quantity 50 (25% discount boundary)"),
        (50.0, 100, 3750.0, "Large quantity (25% discount check)"),
        (20.0, 1, 20.0, "Single item (No discount check)")
    ]

    for price, qty, expected, description in test_cases:
        actual = calculate_discount(price, qty)
        
        # Using math.isclose to handle potential floating point precision issues
        if math.isclose(actual, expected, rel_tol=1e-9):
            print(f"PASS: {description}")
        else:
            # Formatting to 2 decimal places for readable output in case of failure
            print(f"FAIL: {description} (expected: {expected}, got: {actual})")

if __name__ == "__main__":
    run_tests()