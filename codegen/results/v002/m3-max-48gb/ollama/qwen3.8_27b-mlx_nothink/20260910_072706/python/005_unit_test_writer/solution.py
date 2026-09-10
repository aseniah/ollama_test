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
    # Test case 1: quantity = 9 (no discount, boundary just below 10)
    result = calculate_discount(10.0, 9)
    expected = 10.0 * 9
    if abs(result - expected) < 1e-9:
        print("PASS: quantity=9, no discount")
    else:
        print(f"FAIL: quantity=9, no discount (expected: {expected}, got: {result})")

    # Test case 2: quantity = 10 (10% discount, boundary at 10)
    result = calculate_discount(10.0, 10)
    expected = 10.0 * 10 * 0.9
    if abs(result - expected) < 1e-9:
        print("PASS: quantity=10, 10% discount")
    else:
        print(f"FAIL: quantity=10, 10% discount (expected: {expected}, got: {result})")

    # Test case 3: quantity = 49 (10% discount, boundary just below 50)
    result = calculate_discount(10.0, 49)
    expected = 10.0 * 49 * 0.9
    if abs(result - expected) < 1e-9:
        print("PASS: quantity=49, 10% discount")
    else:
        print(f"FAIL: quantity=49, 10% discount (expected: {expected}, got: {result})")

    # Test case 4: quantity = 50 (25% discount, boundary at 50)
    result = calculate_discount(10.0, 50)
    expected = 10.0 * 50 * 0.75
    if abs(result - expected) < 1e-9:
        print("PASS: quantity=50, 25% discount")
    else:
        print(f"FAIL: quantity=50, 25% discount (expected: {expected}, got: {result})")


if __name__ == "__main__":
    test_calculate_discount()