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
    # Helper to assert and print result
    def test(name, expected, price, quantity, tolerance=1e-9):
        result = calculate_discount(price, quantity)
        if abs(result - expected) < tolerance:
            print(f"PASS: {name}")
        else:
            print(f"FAIL: {name} (expected: {expected}, got: {result})")

    # Test cases covering boundary conditions
    # Test 1: quantity < 10 (no discount)
    test("Test 1: quantity=9, price=100 (no discount)", 
         900.0, 100, 9)

    # Test 2: quantity = 10 (10% discount)
    test("Test 2: quantity=10, price=100 (10% discount)", 
         900.0, 100, 10)

    # Test 3: quantity = 49 (10% discount, upper boundary of this range)
    test("Test 3: quantity=49, price=100 (10% discount)", 
         4410.0, 100, 49)

    # Test 4: quantity = 50 (25% discount, lower boundary of this range)
    test("Test 4: quantity=50, price=100 (25% discount)", 
         3750.0, 100, 50)

if __name__ == "__main__":
    run_tests()