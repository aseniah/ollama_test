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
    expected = 10.0 * 9
    got = calculate_discount(price, quantity)
    if abs(got - expected) < 1e-9:
        print(f"PASS: quantity 9 (below 10), no discount")
    else:
        print(f"FAIL: quantity 9 (below 10), no discount (expected: {expected}, got: {got})")
    
    # Test case 2: quantity = 10, 10% discount (boundary)
    price = 10.0
    quantity = 10
    expected = 10.0 * 10 * 0.9
    got = calculate_discount(price, quantity)
    if abs(got - expected) < 1e-9:
        print(f"PASS: quantity 10 (boundary), 10% discount")
    else:
        print(f"FAIL: quantity 10 (boundary), 10% discount (expected: {expected}, got: {got})")
    
    # Test case 3: quantity = 49, 10% discount (boundary)
    price = 10.0
    quantity = 49
    expected = 10.0 * 49 * 0.9
    got = calculate_discount(price, quantity)
    if abs(got - expected) < 1e-9:
        print(f"PASS: quantity 49 (boundary), 10% discount")
    else:
        print(f"FAIL: quantity 49 (boundary), 10% discount (expected: {expected}, got: {got})")
    
    # Test case 4: quantity = 50, 25% discount (boundary)
    price = 10.0
    quantity = 50
    expected = 10.0 * 50 * 0.75
    got = calculate_discount(price, quantity)
    if abs(got - expected) < 1e-9:
        print(f"PASS: quantity 50 (boundary), 25% discount")
    else:
        print(f"FAIL: quantity 50 (boundary), 25% discount (expected: {expected}, got: {got})")
    
    # Test case 5: quantity > 50, 25% discount
    price = 20.0
    quantity = 100
    expected = 20.0 * 100 * 0.75
    got = calculate_discount(price, quantity)
    if abs(got - expected) < 1e-9:
        print(f"PASS: quantity 100 (above 50), 25% discount")
    else:
        print(f"FAIL: quantity 100 (above 50), 25% discount (expected: {expected}, got: {got})")

if __name__ == "__main__":
    test_calculate_discount()