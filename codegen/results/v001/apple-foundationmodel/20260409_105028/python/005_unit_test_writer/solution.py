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
        (100.0, 9),  # No discount
        (100.0, 10), # 10% discount
        (100.0, 49), # 10% discount
        (100.0, 50)  # 25% discount
    ]
    
    for price, quantity in test_cases:
        expected_result = price * quantity
        actual_result = calculate_discount(price, quantity)
        
        if actual_result == expected_result:
            print(f"PASS: Test case - expected: {expected_result}, got: {actual_result}")
        else:
            print(f"FAIL: Test case - expected: {expected_result}, got: {actual_result}")

test_calculate_discount()