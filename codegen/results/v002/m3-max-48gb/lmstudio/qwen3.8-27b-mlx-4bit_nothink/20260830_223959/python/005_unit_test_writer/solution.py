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
        (10.0, 9, 90.0, "quantity=9 (no discount)"),
        (10.0, 10, 90.0, "quantity=10 (10% off boundary)"),
        (10.0, 49, 441.0, "quantity=49 (10% off upper boundary)"),
        (10.0, 50, 375.0, "quantity=50 (25% off boundary)"),
        (5.0, 25, 112.5, "quantity=25, price=5 (10% off)"),
        (100.0, 100, 7500.0, "quantity=100, price=100 (25% off)"),
        (1.0, 0, 0.0, "quantity=0 (no discount)"),
    ]
    
    for price, quantity, expected, description in test_cases:
        got = calculate_discount(price, quantity)
        # Use a small tolerance for floating point comparison
        if abs(got - expected) < 1e-9:
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {got})")


if __name__ == "__main__":
    test_calculate_discount()