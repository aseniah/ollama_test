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
    price = 100.0
    quantity = 9
    expected = price * quantity  # 900.0
    got = calculate_discount(price, quantity)
    if abs(got - expected) < 1e-9:
        print(f"PASSTest: quantity=9, no discount (expected: {expected}, got: {got})")
    else:
        print(f"FAIL: quantity=9, no discount (expected: {expected}, got: {got})")

    # Test case 2: quantity = 10 (10% discount) - lower boundary of discount tier
    price = 100.0
    quantity = 10
    expected = price * quantity * 0.9  # 900.0
    got = calculate_discount(price, quantity)
    if abs(got - expected) < 1e-9:
        print(f"PASSTest: quantity=10, 10% discount (expected: {expected}, got: {got})")
    else:
        print(f"FAIL: quantity=10, 10% discount (expected: {expected}, got: {got})")

    # Test case 3: quantity = 49 (10% discount) - upper boundary of first discount tier
    price = 100.0
    quantity = 49
    expected = price * quantity * 0.9  # 4410.0
    got = calculate_discount(price, quantity)
    if abs(got - expected) < 1e-9:
        print(f"PASSTest: quantity=49, 10% discount (expected: {expected}, got: {got})")
    else:
        print(f"FAIL: quantity=49, 10% discount (expected: {expected}, got: {got})")

    # Test case 4: quantity = 50 (25% discount) - lower boundary of highest discount tier
    price = 100.0
    quantity = 50
    expected = price * quantity * 0.75  # 3750.0
    got = calculate_discount(price, quantity)
    if abs(got - expected) < 1e-9:
        print(f"PASSTest: quantity=50, 25% discount (expected: {expected}, got: {got})")
    else:
        print(f"FAIL: quantity=50, 25% discount (expected: {expected}, got: {got})")


if __name__ == "__main__":
    test_calculate_discount()