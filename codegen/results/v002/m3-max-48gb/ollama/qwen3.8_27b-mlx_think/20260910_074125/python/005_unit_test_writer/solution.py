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
    epsilon = 1e-9

    # Test case 1: quantity 9 - just below 10, no discount
    desc1 = "quantity=9, price=10.0 (no discount, boundary below 10)"
    expected1 = 10.0 * 9  # 90.0
    got1 = calculate_discount(10.0, 9)
    if abs(got1 - expected1) < epsilon:
        print(f"PASS: {desc1}")
    else:
        print(f"FAIL: {desc1} (expected: {expected1}, got: {got1})")

    # Test case 2: quantity 10 - boundary for 10% discount
    desc2 = "quantity=10, price=10.0 (10% off, boundary at 10)"
    expected2 = 10.0 * 10 * 0.9  # 90.0
    got2 = calculate_discount(10.0, 10)
    if abs(got2 - expected2) < epsilon:
        print(f"PASS: {desc2}")
    else:
        print(f"FAIL: {desc2} (expected: {expected2}, got: {got2})")

    # Test case 3: quantity 49 - just below 50, still 10% discount
    desc3 = "quantity=49, price=10.0 (10% off, boundary below 50)"
    expected3 = 10.0 * 49 * 0.9  # 441.0
    got3 = calculate_discount(10.0, 49)
    if abs(got3 - expected3) < epsilon:
        print(f"PASS: {desc3}")
    else:
        print(f"FAIL: {desc3} (expected: {expected3}, got: {got3})")

    # Test case 4: quantity 50 - boundary for 25% discount
    desc4 = "quantity=50, price=10.0 (25% off, boundary at 50)"
    expected4 = 10.0 * 50 * 0.75  # 375.0
    got4 = calculate_discount(10.0, 50)
    if abs(got4 - expected4) < epsilon:
        print(f"PASS: {desc4}")
    else:
        print(f"FAIL: {desc4} (expected: {expected4}, got: {got4})")

    # Test case 5: quantity 100 - well above 50, 25% discount
    desc5 = "quantity=100, price=5.0 (25% off, well above 50)"
    expected5 = 5.0 * 100 * 0.75  # 375.0
    got5 = calculate_discount(5.0, 100)
    if abs(got5 - expected5) < epsilon:
        print(f"PASS: {desc5}")
    else:
        print(f"FAIL: {desc5} (expected: {expected5}, got: {got5})")

    # Test case 6: quantity 1 - minimum quantity, no discount
    desc6 = "quantity=1, price=20.0 (no discount, minimum quantity)"
    expected6 = 20.0 * 1  # 20.0
    got6 = calculate_discount(20.0, 1)
    if abs(got6 - expected6) < epsilon:
        print(f"PASS: {desc6}")
    else:
        print(f"FAIL: {desc6} (expected: {expected6}, got: {got6})")


if __name__ == "__main__":
    run_tests()