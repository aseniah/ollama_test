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
    # Define test cases as (price, quantity, expected_result, description)
    test_cases = [
        (100.0, 9, 900.0, "Lower boundary for no discount (9 units)"),
        (100.0, 10, 900.0, "Lower boundary for 10% discount (10 units)"),
        (100.0, 49, 4410.0, "Upper boundary for 10% discount (49 units)"),
        (100.0, 50, 3750.0, "Lower boundary for 25% discount (50 units)"),
        (10.0, 5, 50.0, "Standard case no discount"),
        (10.0, 100, 750.0, "Large quantity 25% discount"),
    ]

    for price, quantity, expected, description in test_cases:
        actual = calculate_discount(price, quantity)
        
        # Using round() to avoid floating point precision issues during comparison
        if round(actual, 2) == round(expected, 2):
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {actual})")

if __name__ == "__main__":
    run_tests()