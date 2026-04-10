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
    Runs a suite of test cases against calculate_discount.
    """
    # Test case structure: (description, price, quantity, expected_output)
    test_cases = [
        ("Boundary: Quantity 9 (No discount)", 100.0, 9, 900.0),
        ("Boundary: Quantity 10 (10% discount start)", 100.0, 10, 900.0),
        ("Boundary: Quantity 49 (10% discount end)", 100.0, 49, 4410.0),
        ("Boundary: Quantity 50 (25% discount start)", 100.0, 50, 3750.0),
        ("Regular case: Quantity 100 (25% discount)", 10.0, 100, 750.0),
    ]

    for description, price, quantity, expected in test_cases:
        actual = calculate_discount(price, quantity)
        
        # Using math.isclose to handle potential floating point precision issues
        if math.isclose(actual, expected, rel_tol=1e-9):
            print(f"PASS: {description}")
        else:
            # Format numbers to avoid long trailing decimals in error messages
            expected_fmt = f"{expected:.2f}"
            actual_fmt = f"{actual:.2f}"
            print(f"FAIL: {description} (expected: {expected_fmt}, got: {actual_fmt})")

if __name__ == "__main__":
    run_tests()