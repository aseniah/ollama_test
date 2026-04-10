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

# Define test cases covering boundary conditions (quantity 9, 10, 49, 50)
test_cases = [
    # (price, quantity, expected_result, description)
    (100.0, 9, 900.0, "Quantity 9 (No Discount)"),
    (100.0, 10, 900.0, "Quantity 10 (10% Off Boundary)"),
    (100.0, 49, 4410.0, "Quantity 49 (10% Off)"),
    (100.0, 50, 3750.0, "Quantity 50 (25% Off Boundary)")
]

# Execute tests
for price, quantity, expected, description in test_cases:
    result = calculate_discount(price, quantity)
    
    # Compare with rounding to handle floating point precision issues
    if round(result, 2) == round(expected, 2):
        print(f"PASS: {description}")
    else:
        print(f"FAIL: {description} (expected: {expected}, got: {result})")