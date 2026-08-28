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


if __name__ == "__main__":
    test_cases = [
        (10.0, 9, 90.0, "quantity 9 should have no discount"),
        (10.0, 10, 90.0, "quantity 10 should get 10% discount"),
        (10.0, 49, 441.0, "quantity 49 should get 10% discount"),
        (10.0, 50, 375.0, "quantity 50 should get 25% discount"),
        (2.5, 100, 187.5, "quantity 100 should get 25% discount"),
    ]

    for price, quantity, expected, description in test_cases:
        actual = calculate_discount(price, quantity)

        if abs(actual - expected) < 1e-9:
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {actual})")