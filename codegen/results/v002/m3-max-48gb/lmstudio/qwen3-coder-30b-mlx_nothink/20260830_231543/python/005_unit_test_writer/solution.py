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
    (20.0, 3, "quantity 3 (less than 10) - no discount with higher price"),
    (15.0, 60, "quantity 60 (more than 50) - 25% off")
]

for price, quantity, description in test_cases:
    # Calculate expected result
    if quantity < 10:
        expected = price * quantity
    elif quantity < 50:
        expected = price * quantity * 0.9
    else:
        expected = price * quantity * 0.75
    
    # Calculate actual result
    actual = calculate_discount(price, quantity)
    
    # Check if test passes
    if abs(actual - expected) < 0.0001:  # Using small epsilon for float comparison
        print(f"PASS: {description}")
    else:
        print(f"FAIL: {description} (expected: {expected}, got: {actual})")