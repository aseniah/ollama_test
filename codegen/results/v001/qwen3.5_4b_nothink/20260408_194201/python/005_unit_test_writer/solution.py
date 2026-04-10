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


def run_tests():
    price = 100.0
    
    # Test Case 1: quantity = 9 (no discount)
    expected = price * 9
    result = calculate_discount(price, 9)
    if abs(result - expected) < 1e-6:
        print(f"PASS: quantity 9 results in no discount")
    else:
        print(f"FAIL: quantity 9 results in no discount (expected: {expected}, got: {result})")

    # Test Case 2: quantity = 10 (10% discount)
    expected = price * 10 * 0.9
    result = calculate_discount(price, 10)
    if abs(result - expected) < 1e-6:
        print(f"PASS: quantity 10 results in 10% discount")
    else:
        print(f"FAIL: quantity 10 results in 10% discount (expected: {expected}, got: {result})")

    # Test Case 3: quantity = 49 (10% discount)
    expected = price * 49 * 0.9
    result = calculate_discount(price, 49)
    if abs(result - expected) < 1e-6:
        print(f"PASS: quantity 49 results in 10% discount")
    else:
        print(f"FAIL: quantity 49 results in 10% discount (expected: {expected}, got: {result})")

    # Test Case 4: quantity = 50 (25% discount)
    expected = price * 50 * 0.75
    result = calculate_discount(price, 50)
    if abs(result - expected) < 1e-6:
        print(f"PASS: quantity 50 results in 25% discount")
    else:
        print(f"FAIL: quantity 50 results in 25% discount (expected: {expected}, got: {result})")

if __name__ == "__main__":
    run_tests()