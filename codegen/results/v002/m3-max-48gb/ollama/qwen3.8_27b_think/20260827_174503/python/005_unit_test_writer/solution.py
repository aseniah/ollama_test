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
        (100.0, 9, 900.0, "quantity=9 (no discount, boundary below 10)"),
        (100.0, 10, 900.0, "quantity=10 (10% discount, boundary at 10)"),
        (100.0, 49, 4410.0, "quantity=49 (10% discount, boundary below 50)"),
        (100.0, 50, 3750.0, "quantity=50 (25% discount, boundary at 50)"),
        (25.5, 3, 76.5, "quantity=3 (no discount, small quantity)"),
        (50.0, 25, 1125.0, "quantity=25 (10% discount, mid-range)"),
        (10.0, 100, 750.0, "quantity=100 (25% discount, large quantity)"),
    ]

    for price, quantity, expected, description in test_cases:
        result = calculate_discount(price, quantity)
        if abs(result - expected) < 1e-9:
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {result}")


if __name__ == "__main__":
    run_tests()