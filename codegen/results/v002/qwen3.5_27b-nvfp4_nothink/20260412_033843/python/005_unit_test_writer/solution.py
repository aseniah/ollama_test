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
    # Test Case 1: Boundary - quantity just below 10 (no discount)
    description = "Quantity 9 (no discount)"
    price = 10.0
    quantity = 9
    expected = price * quantity  # 90.0
    result = calculate_discount(price, quantity)
    
    if abs(result - expected) < 0.0001:
        print(f"PASS: {description}")
    else:
        print(f"FAIL: {description} (expected: {expected}, got: {result})")

    # Test Case 2: Boundary - quantity exactly 10 (starts 10% discount)
    description = "Quantity 10 (10% discount)"
    price = 10.0
    quantity = 10
    expected = price * quantity * 0.9  # 90.0
    result = calculate_discount(price, quantity)
    
    if abs(result - expected) < 0.0001:
        print(f"PASS: {description}")
    else:
        print(f"FAIL: {description} (expected: {expected}, got: {result})")

    # Test Case 3: Boundary - quantity just below 50 (still 10% discount)
    description = "Quantity 49 (10% discount)"
    price = 10.0
    quantity = 49
    expected = price * quantity * 0.9  # 441.0
    result = calculate_discount(price, quantity)
    
    if abs(result - expected) < 0.0001:
        print(f"PASS: {description}")
    else:
        print(f"FAIL: {description} (expected: {expected}, got: {result})")

    # Test Case 4: Boundary - quantity exactly 50 (starts 25% discount)
    description = "Quantity 50 (25% discount)"
    price = 10.0
    quantity = 50
    expected = price * quantity * 0.75  # 375.0
    result = calculate_discount(price, quantity)
    
    if abs(result - expected) < 0.0001:
        print(f"PASS: {description}")
    else:
        print(f"FAIL: {description} (expected: {expected}, got: {result})")

if __name__ == "__main__":
    test_calculate_discount()