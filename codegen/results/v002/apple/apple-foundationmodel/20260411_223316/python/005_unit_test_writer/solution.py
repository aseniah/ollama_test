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
    # Test case 1: quantity 9 (no discount)
    price = 100.0
    quantity = 9
    expected = price * quantity
    result = calculate_discount(price, quantity)
    print(f"Test case 1: PASS: {result == expected}")
    
    # Test case 2: quantity 10 (10% off)
    price = 100.0
    quantity = 10
    expected = price * quantity * 0.9
    result = calculate_discount(price, quantity)
    print(f"Test case 2: PASS: {result == expected}")
    
    # Test case 3: quantity 49 (9% off)
    price = 100.0
    quantity = 49
    expected = price * quantity * 0.9
    result = calculate_discount(price, quantity)
    print(f"Test case 3: PASS: {result == expected}")
    
    # Test case 4: quantity 50 (25% off)
    price = 100.0
    quantity = 50
    expected = price * quantity * 0.75
    result = calculate_discount(price, quantity)
    print(f"Test case 4: PASS: {result == expected}")

# Run the tests
test_calculate_discount()