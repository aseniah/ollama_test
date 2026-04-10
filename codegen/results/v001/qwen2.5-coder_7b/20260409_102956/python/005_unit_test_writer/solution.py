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
    # Test case 1: quantity less than 10 (boundary condition)
    expected = calculate_discount(10, 9)  # Expected result without discount
    got = calculate_discount(10, 9)
    if got == expected:
        print(f"PASS: Quantity < 10 ({got})")
    else:
        print(f"FAIL: Quantity < 10 (expected: {expected}, got: {got})")

    # Test case 2: quantity equal to 10
    expected = calculate_discount(10, 10) * 0.9  # Expected result with 10% discount
    got = calculate_discount(10, 10)
    if got == expected:
        print(f"PASS: Quantity = 10 ({got})")
    else:
        print(f"FAIL: Quantity = 10 (expected: {expected}, got: {got})")

    # Test case 3: quantity equal to 49
    expected = calculate_discount(10, 49) * 0.9  # Expected result with 10% discount
    got = calculate_discount(10, 49)
    if got == expected:
        print(f"PASS: Quantity = 49 ({got})")
    else:
        print(f"FAIL: Quantity = 49 (expected: {expected}, got: {got})")

    # Test case 4: quantity equal to 50
    expected = calculate_discount(10, 50) * 0.75  # Expected result with 25% discount
    got = calculate_discount(10, 50)
    if got == expected:
        print(f"PASS: Quantity = 50 ({got})")
    else:
        print(f"FAIL: Quantity = 50 (expected: {expected}, got: {got})")

# Run the test cases
test_calculate_discount()