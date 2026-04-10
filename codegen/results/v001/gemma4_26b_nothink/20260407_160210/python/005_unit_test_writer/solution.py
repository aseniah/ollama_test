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
    # Format: (price, quantity, expected_result, description)
    test_cases = [
        (100.0, 9, 900.0, "Boundary: quantity 9 (No discount)"),
        (100.0, 10, 900.0, "Boundary: quantity 10 (10% discount)"),
        (100.0, 49, 4410.0, "Boundary: quantity 49 (10% discount)"),
        (100.0, 50, 3750.0, "Boundary: quantity 50 (25% discount)"),
        (20.0, 100, 1500.0, "Bulk: quantity 100 (25% discount)"),
    ]

    passed_count = 0
    failed_count = 0

    for price, quantity, expected, description in test_cases:
        # Use round to handle potential floating point precision issues in comparison
        actual = round(calculate_discount(price, quantity), 2)
        expected_rounded = round(expected, 2)

        if actual == expected_rounded:
            print(f"PASS: {description}")
            passed_count += 1
        else:
            print(f"FAIL: {description} (expected: {expected_rounded}, got: {actual})")
            failed_count += 1

    print(f"\nTest Summary: {passed_count} passed, {failed_count} failed.")

if __name__ == "__main__":
    run_tests()