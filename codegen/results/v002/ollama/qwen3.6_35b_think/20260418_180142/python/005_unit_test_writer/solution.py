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
        (10.0, 9, 90.0, "No discount for quantity 9"),
        (10.0, 10, 90.0, "10% discount for quantity 10 (lower boundary)"),
        (10.0, 49, 441.0, "10% discount for quantity 49 (upper boundary)"),
        (10.0, 50, 375.0, "25% discount for quantity 50 (lower boundary)"),
    ]

    for price, quantity, expected, description in test_cases:
        result = calculate_discount(price, quantity)
        if result == expected:
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {result})")


if __name__ == "__main__":
    run_tests()