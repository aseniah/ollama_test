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
    # Test Case 1: quantity = 9 (less than 10, no discount)
    # Expected: price * 9
    result_1 = calculate_discount(10.0, 9)
    expected_1 = 10.0 * 9
    passed_1 = abs(result_1 - expected_1) < 1e-6
    print(f"{'PASS' if passed_1 else 'FAIL'}: quantity=9, price=10.0 (expected: {expected_1}, got: {result_1})")

    # Test Case 2: quantity = 10 (minimum for 10% off)
    # Expected: price * 10 * 0.9
    result_2 = calculate_discount(10.0, 10)
    expected_2 = 10.0 * 10 * 0.9
    passed_2 = abs(result_2 - expected_2) < 1e-6
    print(f"{'PASS' if passed_2 else 'FAIL'}: quantity=10, price=10.0 (expected: {expected_2}, got: {result_2})")

    # Test Case 3: quantity = 49 (maximum for 10% off)
    # Expected: price * 49 * 0.9
    result_3 = calculate_discount(10.0, 49)
    expected_3 = 10.0 * 49 * 0.9
    passed_3 = abs(result_3 - expected_3) < 1e-6
    print(f"{'PASS' if passed_3 else 'FAIL'}: quantity=49, price=10.0 (expected: {expected_3}, got: {result_3})")

    # Test Case 4: quantity = 50 (minimum for 25% off)
    # Expected: price * 50 * 0.75
    result_4 = calculate_discount(10.0, 50)
    expected_4 = 10.0 * 50 * 0.75
    passed_4 = abs(result_4 - expected_4) < 1e-6
    print(f"{'PASS' if passed_4 else 'FAIL'}: quantity=50, price=10.0 (expected: {expected_4}, got: {result_4})")

if __name__ == "__main__":
    run_tests()