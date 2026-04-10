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
    (10, 9, "quantity 9"),       # Expected: no discount
    (10, 10, "quantity 10"),     # Expected: 10% off
    (10, 49, "quantity 49"),     # Expected: 10% off
    (10, 50, "quantity 50"),     # Expected: 25% off
]

for price, quantity, description in test_cases:
    expected = calculate_discount(price, quantity)
    result = calculate_discount(price, quantity)
    
    if result == expected:
        print(f"PASS: {description} - Result is correct.")
    else:
        print(f"FAIL: {description} (expected: {expected}, got: {result})")