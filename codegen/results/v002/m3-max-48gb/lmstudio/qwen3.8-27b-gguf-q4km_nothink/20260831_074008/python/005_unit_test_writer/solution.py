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
    # Test case 1: quantity < 10, no discount
    price = 10.0
    quantity = 9
    expected = price * quantity  # 90.0
    result = calculate_discount(price, quantity)
    if abs(result - expected) < 1e-9:
        print(f"PASS: quantity 9 (no discount) - price=10, qty=9, total={result}")
    else:
        print(f"FAIL: quantity 9 (no discount) (expected: {expected}, got: {result})")

    # Test case 2: quantity = 10, 10% off (boundary between no discount and 10% off)
    price = 10.0
    quantity = 10
    expected = price * quantity * 0.9  # 90.0
    result = calculate_discount(price, quantity)
    if abs(result - expected) < 1e-9:
        print(f"PASS: quantity 10 (10% off) - price=10, qty=10, total={result}")
    else:
        print(f"FAIL: quantity 10 (10% off) (expected: {expected}, got: {result})")

    # Test case 3: quantity = 49, 10% off (upper boundary of 10-49 range)
    price = 10.0
    quantity = 49
    expected = price * quantity * 0.9  # 441.0
    result = calculate_discount(price, quantity)
    if abs(result - expected) < 1e-9:
        print(f"PASS: quantity 49 (10% off) - price=10, qty=49, total={result}")
    else:
        print(f"FAIL: quantity 49 (10% off) (expected: {expected}, got: {result})")

    # Test case 4: quantity = 50, 25% off (boundary between 10% off and 25% off)
    price = 10.0
    quantity = 50
    expected = price * quantity * 0.75  # 375.0
    result = calculate_discount(price, quantity)
    if abs(result - expected) < 1e-9:
        print(f"PASS: quantity 50 (25% off) - price=10, qty=50, total={result}")
    else:
        print(f"FAIL: quantity 50 (25% off) (expected: {expected}, got: {result})")

    # Test case 5: quantity > 50, 25% off
    price = 5.0
    quantity = 100
    expected = price * quantity * 0.75  # 375.0
    result = calculate_discount(price, quantity)
    if abs(result - expected) < 1e-9:
        print(f"PASS: quantity 100 (25% off) - price=5, qty=100, total={result}")
    else:
        print(f"FAIL: quantity 100 (25% off) (expected: {expected}, got: {result})")


if __name__ == "__main__":
    test_calculate_discount()