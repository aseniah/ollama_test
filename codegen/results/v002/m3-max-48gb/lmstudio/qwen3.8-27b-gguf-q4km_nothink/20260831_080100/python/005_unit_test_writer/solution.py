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
    (10.0, 9, "quantity 9 (below 10, no discount)", 10.0 * 9),
    (10.0, 10, "quantity 10 (boundary, 10% off)", 10.0 * 10 * 0.9),
    (10.0, 49, "quantity 49 (boundary, 10% off)", 10.0 * 49 * 0.9),
    (10.0, 50, "quantity 50 (boundary, 25% off)", 10.0 * 50 * 0.75),
]

for price, quantity, description, expected in test_cases:
    result = calculate_discount(price, quantity)
    # Use a tolerance for floating point comparison
    if abs(result - expected) < 1e-9:
        print(f"PASS: {description}")
    else:
        print(f"FAIL: {description} (expected: {expected}, got: {result})")