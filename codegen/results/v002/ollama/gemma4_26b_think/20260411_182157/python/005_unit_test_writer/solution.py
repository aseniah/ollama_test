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

def run_test_suite():
    """
    Executes a series of test cases against the calculate_discount function
    and prints the results to stdout.
    """
    # Test cases defined as: (price, quantity, expected_result, description)
    test_cases = [
        (10.0, 9, 90.0, "Boundary: quantity 9 (no discount)"),
        (10.0, 10, 90.0, "Boundary: quantity 10 (10% discount)"),
        (10.0, 49, 441.0, "Boundary: quantity 49 (10% discount)"),
        (10.0, 50, 375.0, "Boundary: quantity 50 (25% discount)"),
        (20.0, 100, 1500.0, "Large quantity (25% discount)")
    ]

    for price, quantity, expected, description in test_cases:
        actual = calculate_discount(price, quantity)
        
        # Using math.isclose to handle potential floating point precision issues
        if math.isclose(actual, expected, rel_tol=1e-9):
            print(f"PASS: {description}")
        else:
            # Formatting to 2 decimal places for cleaner error reporting
            print(f"FAIL: {description} (expected: {expected:.2f}, got: {actual:.2f})")

if __name__ == "__main__":
    run_test_suite()