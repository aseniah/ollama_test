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
    tests_passed = []
    tests_failed = []
    
    # Test case 1: quantity 9 (< 10, no discount)
    price = 100.0
    quantity = 9
    expected = price * quantity
    got = calculate_discount(price, quantity)
    if abs(expected - got) < 0.0001:
        tests_passed.append(f"PASS: quantity 9 - no discount")
    else:
        tests_failed.append(f"FAIL: quantity 9 - expected: {expected}, got: {got}")
    
    # Test case 2: quantity 10 (>= 10 and < 50, 10% off)
    quantity = 10
    expected = price * quantity * 0.9
    got = calculate_discount(price, quantity)
    if abs(expected - got) < 0.0001:
        tests_passed.append(f"PASS: quantity 10 - 10% discount")
    else:
        tests_failed.append(f"FAIL: quantity 10 - expected: {expected}, got: {got}")
    
    # Test case 3: quantity 49 (>= 10 and < 50, 10% off)
    quantity = 49
    expected = price * quantity * 0.9
    got = calculate_discount(price, quantity)
    if abs(expected - got) < 0.0001:
        tests_passed.append(f"PASS: quantity 49 - 10% discount")
    else:
        tests_failed.append(f"FAIL: quantity 49 - expected: {expected}, got: {got}")
    
    # Test case 4: quantity 50 (>= 50, 25% off)
    quantity = 50
    expected = price * quantity * 0.75
    got = calculate_discount(price, quantity)
    if abs(expected - got) < 0.0001:
        tests_passed.append(f"PASS: quantity 50 - 25% discount")
    else:
        tests_failed.append(f"FAIL: quantity 50 - expected: {expected}, got: {got}")
    
    # Print all test results
    for test in tests_passed:
        print(test)
    for test in tests_failed:
        print(test)


if __name__ == "__main__":
    run_tests()