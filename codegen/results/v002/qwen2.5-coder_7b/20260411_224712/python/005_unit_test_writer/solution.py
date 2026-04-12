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
    (10, 5, "quantity below threshold"),
    (20, 10, "quantity at lower boundary of discount"),
    (30, 49, "quantity within discount range"),
    (40, 50, "quantity at upper boundary of discount")
]

for price, quantity, description in test_cases:
    expected = calculate_discount(price, quantity)
    result = calculate_discount(price, quantity)
    if expected == result:
        print(f"PASS: {description}")
    else:
        print(f"FAIL: {description} (expected: {expected}, got: {result})")