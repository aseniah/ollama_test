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
        (100.0, 9),  # quantity 9
        (100.0, 10), # quantity 10
        (100.0, 49), # quantity 49
        (100.0, 50), # quantity 50
    ]
    
    expected_results = [
        (90.0, 9),  # 100 * 9
        (90.0, 10), # 100 * 10 * 0.9
        (81.0, 49), # 100 * 49 * 0.9
        (75.0, 50), # 100 * 50 * 0.75
    ]
    
    for i, (price, quantity) in enumerate(test_cases):
        actual_result = calculate_discount(price, quantity)
        expected_value = expected_results[i][0]
        
        if actual_result == expected_value:
            print(f"PASS: Test case {i+1}: expected {expected_value}, got {actual_result}")
        else:
            print(f"FAIL: Test case {i+1}: expected {expected_value}, got {actual_result}")

# Run the test function
test_calculate_discount()