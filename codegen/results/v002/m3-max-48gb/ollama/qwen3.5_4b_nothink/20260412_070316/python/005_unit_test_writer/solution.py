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
    # Test case 1: quantity < 10 (no discount) - boundary 9
    result = calculate_discount(10.0, 9)
    expected = 90.0
    if result == expected:
        print("PASS: No discount applied for quantity 9")
    else:
        print(f"FAIL: No discount applied for quantity 9 (expected: {expected}, got: {result})")
    
    # Test case 2: quantity >= 10 and < 50 (10% off) - boundary 10
    result = calculate_discount(10.0, 10)
    expected = 90.0
    if result == expected:
        print("PASS: 10% discount applied for quantity 10")
    else:
        print(f"FAIL: 10% discount applied for quantity 10 (expected: {expected}, got: {result})")
    
    # Test case 3: quantity >= 10 and < 50 (10% off) - boundary 49
    result = calculate_discount(10.0, 49)
    expected = 441.0
    if result == expected:
        print("PASS: 10% discount applied for quantity 49")
    else:
        print(f"FAIL: 10% discount applied for quantity 49 (expected: {expected}, got: {result})")
    
    # Test case 4: quantity >= 50 (25% off) - boundary 50
    result = calculate_discount(10.0, 50)
    expected = 375.0
    if result == expected:
        print("PASS: 25% discount applied for quantity 50")
    else:
        print(f"FAIL: 25% discount applied for quantity 50 (expected: {expected}, got: {result})")


if __name__ == "__main__":
    run_tests()