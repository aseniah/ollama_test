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


def run_tests():
    test_cases = [
        (10.0, 9, 10.0 * 9, "quantity 9 (no discount, boundary below 10)"),
        (10.0, 10, 10.0 * 10 * 0.9, "quantity 10 (10% discount, boundary at 10)"),
        (10.0, 49, 10.0 * 49 * 0.9, "quantity 49 (10% discount, boundary below 50)"),
        (10.0, 50, 10.0 * 50 * 0.75, "quantity 50 (25% discount, boundary at 50)"),
        (5.0, 1, 5.0 * 1, "quantity 1 (no discount, minimum quantity)"),
        (100.0, 25, 100.0 * 25 * 0.9, "quantity 25 (10% discount, middle range)"),
        (10.0, 100, 10.0 * 100 * 0.75, "quantity 100 (25% discount, well above 50)"),
    ]
    
    for price, quantity, expected, description in test_cases:
        got = calculate_discount(price, quantity)
        # Use a tolerance for floating point comparison
        if abs(got - expected) < 1e-9:
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {got})")

if __name__ == "__main__":
    run_tests()