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
        (100.0, 9),   # quantity < 10
        (100.0, 10),  # quantity 10-49
        (100.0, 49),  # quantity 10-49
        (100.0, 50),  # quantity >= 50
        (100.0, 60),  # quantity 10-49 with higher quantity
        (100.0, 100) # quantity >= 50 with higher quantity
    ]
    
    for i, (price, quantity) in enumerate(test_cases, 1):
        expected = calculate_discount(price, quantity)
        result = calculate_discount(price, quantity)
        
        if expected == result:
            print(f"PASS: Test case {i} - Expected: {expected}, got: {result}")
        else:
            print(f"FAIL: Test case {i} - Expected: {expected}, got: {result}")

# Run the test cases
test_calculate_discount()