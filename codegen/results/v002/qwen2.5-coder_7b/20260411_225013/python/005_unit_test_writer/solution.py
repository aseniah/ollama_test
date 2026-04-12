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
    # Test case 1: quantity < 10 (boundary condition)
    if calculate_discount(10, 9) == 90:
        print("PASS: Boundary condition - quantity = 9")
    else:
        print("FAIL: Boundary condition - quantity = 9")

    # Test case 2: quantity = 10
    if calculate_discount(10, 10) == 90:
        print("PASS: Edge case - quantity = 10")
    else:
        print("FAIL: Edge case - quantity = 10")

    # Test case 3: quantity < 50 (general case)
    if calculate_discount(10, 25) == 225:
        print("PASS: General case - quantity = 25")
    else:
        print("FAIL: General case - quantity = 25")

    # Test case 4: quantity >= 50 (boundary condition)
    if calculate_discount(10, 50) == 375:
        print("PASS: Boundary condition - quantity = 50")
    else:
        print("FAIL: Boundary condition - quantity = 50")


test_calculate_discount()