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
    result = calculate_discount(10, 8)
    expected = 80
    if result == expected:
        print("PASS: Quantity 9 should not have a discount")
    else:
        print(f"FAIL: Expected {expected}, got {result}")

    # Test case 2: quantity = 10, 10% off
    result = calculate_discount(10, 10)
    expected = 90
    if result == expected:
        print("PASS: Quantity 10 should have a 10% discount")
    else:
        print(f"FAIL: Expected {expected}, got {result}")

    # Test case 3: quantity = 49, 10% off
    result = calculate_discount(10, 49)
    expected = 441
    if result == expected:
        print("PASS: Quantity 49 should have a 10% discount")
    else:
        print(f"FAIL: Expected {expected}, got {result}")

    # Test case 4: quantity >= 50, 25% off
    result = calculate_discount(10, 50)
    expected = 375
    if result == expected:
        print("PASS: Quantity 50 should have a 25% discount")
    else:
        print(f"FAIL: Expected {expected}, got {result}")

# Run the test cases
test_calculate_discount()