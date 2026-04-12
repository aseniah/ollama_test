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


# Run the tests
if __name__ == "__main__":
    # Test Case 1: quantity < 10 (no discount) - boundary at 9
    test_price = 100.0
    test_quantity_1 = 9
    expected_price_1 = test_price * test_quantity_1
    got_price_1 = calculate_discount(test_price, test_quantity_1)
    if abs(expected_price_1 - got_price_1) < 1e-6:
        print("PASS: quantity=9, price=100.0 (no discount expected)")
    else:
        print("FAIL: quantity=9, price=100.0 (expected: " + str(expected_price_1) + ", got: " + str(got_price_1) + ")")

    # Test Case 2: quantity = 10 (enter bulk discount tier) - boundary
    test_quantity_2 = 10
    expected_price_2 = test_price * test_quantity_2 * 0.9
    got_price_2 = calculate_discount(test_price, test_quantity_2)
    if abs(expected_price_2 - got_price_2) < 1e-6:
        print("PASS: quantity=10, price=100.0 (10% discount expected)")
    else:
        print("FAIL: quantity=10, price=100.0 (expected: " + str(expected_price_2) + ", got: " + str(got_price_2) + ")")

    # Test Case 3: quantity = 49 (upper limit of 10% discount tier)
    test_quantity_3 = 49
    expected_price_3 = test_price * test_quantity_3 * 0.9
    got_price_3 = calculate_discount(test_price, test_quantity_3)
    if abs(expected_price_3 - got_price_3) < 1e-6:
        print("PASS: quantity=49, price=100.0 (10% discount expected)")
    else:
        print("FAIL: quantity=49, price=100.0 (expected: " + str(expected_price_3) + ", got: " + str(got_price_3) + ")")

    # Test Case 4: quantity = 50 (enter 25% discount tier)
    test_quantity_4 = 50
    expected_price_4 = test_price * test_quantity_4 * 0.75
    got_price_4 = calculate_discount(test_price, test_quantity_4)
    if abs(expected_price_4 - got_price_4) < 1e-6:
        print("PASS: quantity=50, price=100.0 (25% discount expected)")
    else:
        print("FAIL: quantity=50, price=100.0 (expected: " + str(expected_price_4) + ", got: " + str(got_price_4) + ")")

    # Test Case 5: quantity > 50 (ensure highest discount tier works)
    test_quantity_5 = 100
    expected_price_5 = test_price * test_quantity_5 * 0.75
    got_price_5 = calculate_discount(test_price, test_quantity_5)
    if abs(expected_price_5 - got_price_5) < 1e-6:
        print("PASS: quantity=100, price=100.0 (25% discount expected)")
    else:
        print("FAIL: quantity=100, price=100.0 (expected: " + str(expected_price_5) + ", got: " + str(got_price_5) + ")")

    # Test Case 6: test with different price values
    test_price_2 = 50.0
    test_quantity_6 = 7
    expected_price_6 = test_price_2 * test_quantity_6
    got_price_6 = calculate_discount(test_price_2, test_quantity_6)
    if abs(expected_price_6 - got_price_6) < 1e-6:
        print("PASS: price=50.0, quantity=7 (no discount expected)")
    else:
        print("FAIL: price=50.0, quantity=7 (expected: " + str(expected_price_6) + ", got: " + str(got_price_6) + ")")

    # Test Case 7: edge case near threshold 10
    test_quantity_7 = 8
    expected_price_7 = test_price * test_quantity_7
    got_price_7 = calculate_discount(test_price, test_quantity_7)
    if abs(expected_price_7 - got_price_7) < 1e-6:
        print("PASS: quantity=8, price=100.0 (no discount expected)")
    else:
        print("FAIL: quantity=8, price=100.0 (expected: " + str(expected_price_7) + ", got: " + str(got_price_7) + ")")