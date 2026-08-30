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
    (10.0, 5, "quantity 5 (less than 10) - no discount"),
    (10.0, 10, "quantity 10 (boundary between no discount and 10% off)"),
    (10.0, 49, "quantity 49 (boundary between 10% off and 25% off)"),
    (10.0, 50, "quantity 50 (boundary for 25% off)"),
    (5.0, 25, "quantity 25 (10% off) - normal case"),
    (15.0, 75, "quantity 75 (25% off) - normal case")
]

# Run tests
for price, quantity, description in test_cases:
    result = calculate_discount(price, quantity)
    
    # Calculate expected result manually
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