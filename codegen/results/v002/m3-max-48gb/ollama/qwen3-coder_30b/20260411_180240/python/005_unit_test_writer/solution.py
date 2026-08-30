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

# Test cases
test_cases = [
    (10.0, 5, "price 10.0, quantity 5 (no discount)"),
    (10.0, 10, "price 10.0, quantity 10 (10% discount)"),
    (10.0, 49, "price 10.0, quantity 49 (10% discount)"),
    (10.0, 50, "price 10.0, quantity 50 (25% discount)"),
    (5.0, 9, "price 5.0, quantity 9 (no discount)"),
    (15.0, 25, "price 15.0, quantity 25 (10% discount)"),
    (8.0, 75, "price 8.0, quantity 75 (25% discount)")
]

# Run tests
for price, quantity, description in test_cases:
    result = calculate_discount(price, quantity)
    
    # Calculate expected result
    if quantity < 10:
        expected = price * quantity
    elif quantity < 50:
        expected = price * quantity * 0.9
    else:
        expected = price * quantity * 0.75
    
    if result == expected:
        print(f"PASS: {description}")
    else:
        print(f"FAIL: {description} (expected: {expected}, got: {result})")