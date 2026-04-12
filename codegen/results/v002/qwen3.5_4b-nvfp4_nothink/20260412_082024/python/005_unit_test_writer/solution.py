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
    # Test Case 1: quantity = 9 (boundary low, no discount)
    # Expected: 10 * 9 * 1.0 = 90.0
    price, qty = 10.0, 9
    expected = price * qty * 1.0
    got = calculate_discount(price, qty)
    status = "PASS: Test case with quantity=9 (no discount)" if abs(expected - got) < 1e-6 else f"FAIL: Test case with quantity=9 (expected: {expected}, got: {got})"
    print(status)

    # Test Case 2: quantity = 10 (boundary low/mid, 10% off)
    # Expected: 10 * 10 * 0.9 = 90.0
    price, qty = 10.0, 10
    expected = price * qty * 0.9
    got = calculate_discount(price, qty)
    status = "PASS: Test case with quantity=10 (10% off)" if abs(expected - got) < 1e-6 else f"FAIL: Test case with quantity=10 (expected: {expected}, got: {got})"
    print(status)

    # Test Case 3: quantity = 49 (boundary high/mid, 10% off)
    # Expected: 10 * 49 * 0.9 = 441.0
    price, qty = 10.0, 49
    expected = price * qty * 0.9
    got = calculate_discount(price, qty)
    status = "PASS: Test case with quantity=49 (10% off)" if abs(expected - got) < 1e-6 else f"FAIL: Test case with quantity=49 (expected: {expected}, got: {got})"
    print(status)

    # Test Case 4: quantity = 50 (boundary low/high, 25% off)
    # Expected: 10 * 50 * 0.75 = 375.0
    price, qty = 10.0, 50
    expected = price * qty * 0.75
    got = calculate_discount(price, qty)
    status = "PASS: Test case with quantity=50 (25% off)" if abs(expected - got) < 1e-6 else f"FAIL: Test case with quantity=50 (expected: {expected}, got: {got})"
    print(status)

if __name__ == "__main__":
    test_calculate_discount()