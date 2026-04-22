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
    """Test the calculate_discount function with various test cases."""
    
    # Test case 1: No discount (quantity < 10)
    price = 100.0
    quantity = 9
    expected = price * quantity  # 900.0
    result = calculate_discount(price, quantity)
    if abs(result - expected) < 1e-9:
        print(f"PASS: No discount for quantity 9 at price {price}")
    else:
        print(f"FAIL: No discount for quantity 9 at price {price} (expected: {expected}, got: {result})")
    
    # Test case 2: 10% discount boundary (quantity == 10)
    price = 100.0
    quantity = 10
    expected = price * quantity * 0.9  # 900.0
    result = calculate_discount(price, quantity)
    if abs(result - expected) < 1e-9:
        print(f"PASS: 10% discount for quantity 10 at price {price}")
    else:
        print(f"FAIL: 10% discount for quantity 10 at price {price} (expected: {expected}, got: {result})")
    
    # Test case 3: 10% discount at upper boundary (quantity == 49)
    price = 100.0
    quantity = 49
    expected = price * quantity * 0.9  # 4410.0
    result = calculate_discount(price, quantity)
    if abs(result - expected) < 1e-9:
        print(f"PASS: 10% discount for quantity 49 at price {price}")
    else:
        print(f"FAIL: 10% discount for quantity 49 at price {price} (expected: {expected}, got: {result})")
    
    # Test case 4: 25% discount at boundary (quantity == 50)
    price = 100.0
    quantity = 50
    expected = price * quantity * 0.75  # 3750.0
    result = calculate_discount(price, quantity)
    if abs(result - expected) < 1e-9:
        print(f"PASS: 25% discount for quantity 50 at price {price}")
    else:
        print(f"FAIL: 25% discount for quantity 50 at price {price} (expected: {expected}, got: {result})")
    
    # Test case 5: Extra test - quantity 0
    price = 100.0
    quantity = 0
    expected = price * quantity  # 0.0
    result = calculate_discount(price, quantity)
    if abs(result - expected) < 1e-9:
        print(f"PASS: No cost for quantity 0 at price {price}")
    else:
        print(f"FAIL: No cost for quantity 0 at price {price} (expected: {expected}, got: {result})")
    
    # Test case 6: Extra test - quantity 100
    price = 50.0
    quantity = 100
    expected = price * quantity * 0.75  # 3750.0
    result = calculate_discount(price, quantity)
    if abs(result - expected) < 1e-9:
        print(f"PASS: 25% discount for quantity 100 at price {price}")
    else:
        print(f"FAIL: 25% discount for quantity 100 at price {price} (expected: {expected}, got: {result})")


if __name__ == "__main__":
    test_calculate_discount()