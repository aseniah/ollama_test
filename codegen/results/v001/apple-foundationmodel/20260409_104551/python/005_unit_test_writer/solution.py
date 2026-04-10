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
    # Test case 1: quantity = 9 (no discount)
    price = 100.0
    quantity = 9
    expected = 900.0
    result = calculate_discount(price, quantity)
    if result == expected:
        print(f"PASS: Test case 1 - quantity = 9 (no discount)")
    else:
        print(f"FAIL: Test case 1 - quantity = 9 (expected: {expected}, got: {result})")
    
    # Test case 2: quantity = 10 (10% discount)
    price = 100.0
    quantity = 10
    expected = 900.0
    result = calculate_discount(price, quantity)
    if result == expected:
        print(f"PASS: Test case 2 - quantity = 10 (10% discount)")
    else:
        print(f"FAIL: Test case 2 - quantity = 10 (expected: {expected}, got: {result})")
    
    # Test case 3: quantity = 49 (9% discount)
    price = 100.0
    quantity = 49
    expected = 4410.0
    result = calculate_discount(price, quantity)
    if result == expected:
        print(f"PASS: Test case 3 - quantity = 49 (9% discount)")
    else:
        print(f"FAIL: Test case 3 - quantity = 49 (expected: {expected}, got: {result})")
    
    # Test case 4: quantity = 50 (25% discount)
    price = 100.0
    quantity = 50
    expected = 7500.0
    result = calculate_discount(price, quantity)
    if result == expected:
        print(f"PASS: Test case 4 - quantity = 50 (25% discount)")
    else:
        print(f"FAIL: Test case 4 - quantity = 50 (expected: {expected}, got: {result})")

# Run the tests
test_calculate_discount()