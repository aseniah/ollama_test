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
    # Test Case 1: Boundary - Just below discount (quantity 9) -> No discount
    test_1_qty = 9
    test_1_price = 20.0
    expected_1 = test_1_price * test_1_qty  # 180.0
    got_1 = calculate_discount(test_1_price, test_1_qty)
    
    if abs(got_1 - expected_1) < 0.01:
        print("PASS: Quantity 9 (boundary below discount threshold)")
    else:
        print(f"FAIL: Quantity 9 (boundary below discount threshold) (expected: {expected_1}, got: {got_1})")

    # Test Case 2: Boundary - Start of 10% discount (quantity 10)
    test_2_qty = 10
    test_2_price = 20.0
    expected_2 = test_2_price * test_2_qty * 0.9  # 180.0
    got_2 = calculate_discount(test_2_price, test_2_qty)

    if abs(got_2 - expected_2) < 0.01:
        print("PASS: Quantity 10 (start of 10% discount)")
    else:
        print(f"FAIL: Quantity 10 (start of 10% discount) (expected: {expected_2}, got: {got_2})")

    # Test Case 3: Boundary - End of 10% discount (quantity 49)
    test_3_qty = 49
    test_3_price = 10.0
    expected_3 = test_3_price * test_3_qty * 0.9  # 441.0
    got_3 = calculate_discount(test_3_price, test_3_qty)

    if abs(got_3 - expected_3) < 0.01:
        print("PASS: Quantity 49 (end of 10% discount)")
    else:
        print(f"FAIL: Quantity 49 (end of 10% discount) (expected: {expected_3}, got: {got_3})")

    # Test Case 4: Boundary - Start of 25% discount (quantity 50)
    test_4_qty = 50
    test_4_price = 10.0
    expected_4 = test_4_price * test_4_qty * 0.75  # 375.0
    got_4 = calculate_discount(test_4_price, test_4_qty)

    if abs(got_4 - expected_4) < 0.01:
        print("PASS: Quantity 50 (start of 25% discount)")
    else:
        print(f"FAIL: Quantity 50 (start of 25% discount) (expected: {expected_4}, got: {got_4})")

if __name__ == "__main__":
    run_tests()