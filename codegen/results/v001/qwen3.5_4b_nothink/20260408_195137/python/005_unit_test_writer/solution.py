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
    # Test cases: (price, quantity, expected_description)
    test_cases = [
        ("boundary_low", 10.0, 9),
        ("low", 10.0, 5),
        ("boundary_mid", 10.0, 10),
        ("mid", 10.0, 25),
        ("boundary_high", 10.0, 49),
        ("high", 10.0, 100),
        ("boundary_very_high", 10.0, 50),
    ]

    for name, price, quantity in test_cases:
        expected = calculate_discount(price, quantity)
        # Use 10^-6 as tolerance for float comparison
        tolerance = 1e-6
        if abs(expected - round(expected, 10)) < tolerance:
            print(f"PASS: {name} - price {price}, quantity {quantity}")
        else:
            print(f"FAIL: {name} - expected: {round(expected, 2)}, got: {round(round(expected, 10), 2)}")


if __name__ == "__main__":
    run_tests()