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
    test_cases = [
        {"price": 10.0, "quantity": 5, "expected": 50.0, "description": "quantity < 10 no discount"},
        {"price": 10.0, "quantity": 9, "expected": 90.0, "description": "quantity = 9 (boundary, no discount)"},
        {"price": 10.0, "quantity": 10, "expected": 90.0, "description": "quantity = 10 (minimum for 10% off)"},
        {"price": 10.0, "quantity": 49, "expected": 810.0, "description": "quantity = 49 (maximum for 10% off)"},
        {"price": 10.0, "quantity": 50, "expected": 375.0, "description": "quantity = 50 (minimum for 25% off)"},
        {"price": 25.0, "quantity": 15, "expected": 360.0, "description": "middle quantity within 10-49 range"},
    ]
    
    passed = 0
    failed = 0
    
    print("=" * 70)
    print("Testing calculate_discount function")
    print("=" * 70)
    print()
    
    for case in test_cases:
        price = case["price"]
        quantity = case["quantity"]
        expected = case["expected"]
        description = case["description"]
        
        result = calculate_discount(price, quantity)
        
        if abs(result - expected) < 0.001:
            print(f"PASS: {description}")
            passed += 1
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {result})")
            failed += 1
    
    print()
    print("=" * 70)
    print(f"Results: {passed} passed, {failed} failed")
    print("=" * 70)

if __name__ == "__main__":
    test_calculate_discount()