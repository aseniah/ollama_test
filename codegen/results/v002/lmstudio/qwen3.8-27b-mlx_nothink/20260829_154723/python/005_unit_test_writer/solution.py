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
    # Define test cases: (price, quantity, expected, description)
    test_cases = [
        (5.0, 9, 45.0, "quantity 9 (below 10, no discount)"),
        (5.0, 10, 45.0, "quantity 10 (boundary, 10% off)"),
        (5.0, 49, 220.5, "quantity 49 (upper 10-49, 10% off)"),
        (5.0, 50, 187.5, "quantity 50 (boundary, 25% off)"),
    ]
    
    # Additional test cases for completeness
    test_cases.append((10.0, 1, 10.0, "quantity 1 (no discount)"))
    test_cases.append((10.0, 50, 375.0, "quantity 50 with price 10 (25% off)"))
    test_cases.append((2.5, 25, 56.25, "quantity 25 with price 2.5 (10% off)"))
    test_cases.append((0.0, 100, 0.0, "price 0 with quantity 100 (free)"))
    
    for price, quantity, expected, description in test_cases:
        got = calculate_discount(price, quantity)
        # Use a tolerance for floating point comparison
        if abs(got - expected) < 1e-9:
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {got}")


if __name__ == "__main__":
    run_tests()