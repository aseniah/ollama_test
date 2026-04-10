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
    """Test cases for calculate_discount function."""
    
    # Test case 1: quantity < 10 (no discount)
    print("Test 1: Basic pricing with no discount")
    expected = 100 * 9  # price=100, quantity=9
    actual = calculate_discount(100, 9)
    if actual == expected:
        print(f"PASS: Basic pricing (price=100, qty=9), expected={expected}, got={actual}")
    else:
        print(f"FAIL: Basic pricing (price=100, qty=9), expected={expected}, got={actual}")
    
    # Test case 2: quantity = 10 (start of 10% discount)
    print("\nTest 2: Boundary at 10% discount threshold")
    expected = 100 * 10 * 0.9
    actual = calculate_discount(100, 10)
    if actual == expected:
        print(f"PASS: Boundary at 10% discount (price=100, qty=10), expected={expected}, got={actual}")
    else:
        print(f"FAIL: Boundary at 10% discount (price=100, qty=10), expected={expected}, got={actual}")
    
    # Test case 3: quantity = 49 (end of 10% discount)
    print("\nTest 3: Boundary at end of 10% discount range")
    expected = 100 * 49 * 0.9
    actual = calculate_discount(100, 49)
    if actual == expected:
        print(f"PASS: Boundary at end of 10% discount (price=100, qty=49), expected={expected}, got={actual}")
    else:
        print(f"FAIL: Boundary at end of 10% discount (price=100, qty=49), expected={expected}, got={actual}")
    
    # Test case 4: quantity = 50 (start of 25% discount)
    print("\nTest 4: Boundary at 25% discount threshold")
    expected = 100 * 50 * 0.75
    actual = calculate_discount(100, 50)
    if actual == expected:
        print(f"PASS: Boundary at 25% discount (price=100, qty=50), expected={expected}, got={actual}")
    else:
        print(f"FAIL: Boundary at 25% discount (price=100, qty=50), expected={expected}, got={actual}")


if __name__ == "__main__":
    test_calculate_discount()