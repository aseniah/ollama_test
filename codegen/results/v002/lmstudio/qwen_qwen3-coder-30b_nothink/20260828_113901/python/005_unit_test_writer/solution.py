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
    (10.0, 10, "quantity 10 (boundary case) - 10% off"),
    (10.0, 49, "quantity 49 (boundary case) - 10% off"),
    (10.0, 50, "quantity 50 (boundary case) - 25% off"),
    (5.0, 25, "quantity 25 (middle range) - 10% off"),
    (15.0, 60, "quantity 60 (high range) - 25% off")
]

for price, quantity, description in test_cases:
    result = calculate_discount(price, quantity)
    
    # Calculate expected result manually
    if quantity < 10:
        expected = price * quantity
    elif quantity < 50:
        expected = price * quantity * 0.9
    else:
        expected = price * quantity * 0.75
    
    if abs(result - expected) < 0.0001:  # Using small tolerance for float comparison
        print(f"PASS: {description}")
    else:
        print(f"FAIL: {description} (expected: {expected}, got: {result})")