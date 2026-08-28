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
    """Run tests on calculate_discount function."""
    
    def run_test(description, actual_price, expected_price, passed=None):
        if passed is None:
            print(f"FAIL: {description}")
            print(f"expected: {expected_price}, got: {actual_price}")
        else:
            print(f"PASS: {description}")
    
    # Test case 1: quantity < 10 (no discount) - boundary condition 9
    price_1 = 100.0
    quantity_1 = 9
    expected_1 = 100.0 * 9
    actual_1 = calculate_discount(price_1, quantity_1)
    print(f"Test: quantity=9 (boundary no discount)")
    if abs(actual_1 - expected_1) < 0.0001:
        print("PASS: quantity=9 with price=100.0 - no discount applied")
    else:
        print("FAIL: quantity=9 with price=100.0 - no discount applied")
        print(f"expected: {expected_1}, got: {actual_1}")
    
    # Test case 2: quantity >= 10 and < 50 (10% discount) - boundary condition 10
    price_2 = 100.0
    quantity_2 = 10
    expected_2 = 100.0 * 10 * 0.9
    actual_2 = calculate_discount(price_2, quantity_2)
    print(f"\nTest: quantity=10 (boundary 10-49 discount)")
    if abs(actual_2 - expected_2) < 0.0001:
        print("PASS: quantity=10 with price=100.0 - 10% discount applied")
    else:
        print("FAIL: quantity=10 with price=100.0 - 10% discount applied")
        print(f"expected: {expected_2}, got: {actual_2}")
    
    # Test case 3: quantity = 49 (still 10% discount) - boundary condition 49
    price_3 = 50.0
    quantity_3 = 49
    expected_3 = 50.0 * 49 * 0.9
    actual_3 = calculate_discount(price_3, quantity_3)
    print(f"\nTest: quantity=49 (boundary max for 10% discount)")
    if abs(actual_3 - expected_3) < 0.0001:
        print("PASS: quantity=49 with price=50.0 - 10% discount applied")
    else:
        print("FAIL: quantity=49 with price=50.0 - 10% discount applied")
        print(f"expected: {expected_3}, got: {actual_3}")
    
    # Test case 4: quantity >= 50 (25% discount) - boundary condition 50
    price_4 = 50.0
    quantity_4 = 50
    expected_4 = 50.0 * 50 * 0.75
    actual_4 = calculate_discount(price_4, quantity_4)
    print(f"\nTest: quantity=50 (boundary >= 50 discount)")
    if abs(actual_4 - expected_4) < 0.0001:
        print("PASS: quantity=50 with price=50.0 - 25% discount applied")
    else:
        print("FAIL: quantity=50 with price=50.0 - 25% discount applied")
        print(f"expected: {expected_4}, got: {actual_4}")


if __name__ == "__main__":
    test_calculate_discount()