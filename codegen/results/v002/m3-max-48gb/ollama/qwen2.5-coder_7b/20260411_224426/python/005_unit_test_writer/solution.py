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
    result = calculate_discount(10, 9)
    expected = 90
    if result == expected:
        print(f"PASS: Quantity less than 10 - Expected {expected}, got {result}")
    else:
        print(f"FAIL: Quantity less than 10 - Expected {expected}, got {result}")

    # Test case 2: quantity exactly 10 (boundary condition)
    result = calculate_discount(10, 10)
    expected = 90
    if result == expected:
        print(f"PASS: Quantity equals 10 - Expected {expected}, got {result}")
    else:
        print(f"FAIL: Quantity equals 10 - Expected {expected}, got {result}")

    # Test case 3: quantity < 50 (within discount range)
    result = calculate_discount(10, 49)
    expected = 441
    if result == expected:
        print(f"PASS: Quantity between 10 and 49 - Expected {expected}, got {result}")
    else:
        print(f"FAIL: Quantity between 10 and 49 - Expected {expected}, got {result}")

    # Test case 4: quantity exactly 50 (boundary condition)
    result = calculate_discount(10, 50)
    expected = 375
    if result == expected:
        print(f"PASS: Quantity equals 50 - Expected {expected}, got {result}")
    else:
        print(f"FAIL: Quantity equals 50 - Expected {expected}, got {result}")

# Run the tests
test_calculate_discount()