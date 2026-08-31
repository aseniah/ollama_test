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
        (10.0, 9, 90.0, "Quantity 9 (below 10, no discount)"),
        (10.0, 10, 90.0, "Quantity 10 (boundary, 10% off)"),
        (10.0, 49, 441.0, "Quantity 49 (boundary, 10% off)"),
        (10.0, 50, 375.0, "Quantity 50 (boundary, 25% off)"),
    ]
    
    for price, quantity, expected, description in test_cases:
        got = calculate_discount(price, quantity)
        if abs(got - expected) < 1e-9:
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {got})")


test_calculate_discount()