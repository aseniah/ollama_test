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
    Includes boundary conditions for quantities 9, 10, 49, and 50.
    """
    # List of test cases: (price, quantity, expected_output, description)
    test_cases = [
        (10.0, 9, 90.0, "Boundary: quantity 9 (No discount)"),
        (10.0, 10, 90.0, "Boundary: quantity 10 (10% discount)"),
        (10.0, 49, 441.0, "Boundary: quantity 49 (10% discount)"),
        (10.0, 50, 375.0, "Boundary: quantity 50 (25% discount)"),
        (100.0, 100, 7500.0, "General: Large quantity (25% discount)"),
        (5.0, 1, 5.0, "General: Single item (No discount)")
    ]

    for price, quantity, expected, description in test_cases:
        # Calculate actual result
        actual = calculate_discount(price, quantity)
        
        # Using round to handle potential floating point precision issues in comparison
        # We round to 2 decimal places as is standard for currency
        if round(actual, 2) == round(expected, 2):
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {actual})")


if __name__ == "__main__":
    run_tests()