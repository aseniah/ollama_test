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
    # Test case 1: quantity < 10 (no discount) - boundary at 9
    price = 10.0
    quantity = 9
    expected = price * quantity  # 90.0
    result = calculate_discount(price, quantity)
    if abs(result - expected) < 1e-9:
        print(f"PASSESS: quantity {quantity} has no discount")
    else:
        print(f"FAIL: quantity {quantity} has no discount (expected: {expected}, got: {result})")

    # Test case 2: quantity = 10 (10% discount) - boundary at 10
    price = 10.0
    quantity = 10
    expected = price * quantity * 0.9  # 90.0
    result = calculate_discount(price, quantity)
    if abs(result - expected) < 1e-9:
        print(f"PASSED: quantity {quantity} has 10% discount")
    else:
        print(f"FAIL: quantity {quantity} has 10% discount (expected: {expected}, got: {result})")

    # Test case 3: quantity = 49 (10% discount) - boundary at 49
    price = 10.0
    quantity = 49
    expected = price * quantity * 0.9  # 441.0
    result = calculate_discount(price, quantity)
    if abs(result - expected) < 1e-9:
        print(f"PASSED: quantity {quantity} has 10% discount")
    else:
        print(f"FAIL: quantity {quantity} has 10% discount (expected: {expected}, got: {result})")

    # Test case 4: quantity = 50 (25% discount) - boundary at 50
    price = 10.0
    quantity = 50
    expected = price * quantity * 0.75  # 375.0
    result = calculate_discount(price, quantity)
    if abs(result - expected) < 1e-9:
        print(f"PASSED: quantity {quantity} has 25% discount")
    else:
        print(f"FAIL: quantity {quantity} has 25% discount (expected: {expected}, got: {result})")


if __name__ == "__main__":
    test_calculate_discount()