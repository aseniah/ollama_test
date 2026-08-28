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


def test_calculate_discount():
    # Test case 1: quantity < 10 (no discount)
    def test_case_1():
        price = 10.0
        quantity = 9
        expected = price * quantity
        actual = calculate_discount(price, quantity)
        status = "PASS" if abs(actual - expected) < 0.001 else "FAIL"
        description = f"quantity 9 (no discount)"
        print(f"{status}: {description} (expected: {expected}, got: {actual})")
    
    # Test case 2: quantity 10-49 (10% off)
    def test_case_2():
        price = 10.0
        quantity = 10
        expected = price * quantity * 0.9
        actual = calculate_discount(price, quantity)
        status = "PASS" if abs(actual - expected) < 0.001 else "FAIL"
        description = f"quantity 10 (10% off)"
        print(f"{status}: {description} (expected: {expected}, got: {actual})")
    
    # Test case 3: quantity 49 (boundary of 10-49 range)
    def test_case_3():
        price = 10.0
        quantity = 49
        expected = price * quantity * 0.9
        actual = calculate_discount(price, quantity)
        status = "PASS" if abs(actual - expected) < 0.001 else "FAIL"
        description = f"quantity 49 (10% off)"
        print(f"{status}: {description} (expected: {expected}, got: {actual})")
    
    # Test case 4: quantity >= 50 (25% off)
    def test_case_4():
        price = 10.0
        quantity = 50
        expected = price * quantity * 0.75
        actual = calculate_discount(price, quantity)
        status = "PASS" if abs(actual - expected) < 0.001 else "FAIL"
        description = f"quantity 50 (25% off)"
        print(f"{status}: {description} (expected: {expected}, got: {actual})")
    
    # Run all test cases
    test_case_1()
    test_case_2()
    test_case_3()
    test_case_4()


if __name__ == "__main__":
    test_calculate_discount()