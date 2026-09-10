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


def run_test(description, actual, expected):
    if abs(actual - expected) < 1e-9:
        print(f"PASS: {description}")
    else:
        print(f"FAIL: {description} (expected: {expected}, got: {actual})")


# Test 1: quantity = 9 (just below the 10% discount threshold - no discount)
# 10.0 * 9 = 90.0
run_test("quantity 9: no discount applied",
         calculate_discount(10.0, 9), 90.0)

# Test 2: quantity = 10 (boundary: enters 10% discount bracket)
# 10.0 * 10 * 0.9 = 90.0
run_test("quantity 10: 10% discount boundary",
         calculate_discount(10.0, 10), 90.0)

# Test 3: quantity = 49 (boundary: last value in 10% discount bracket)
# 10.0 * 49 * 0.9 = 441.0
run_test("quantity 49: 10% discount upper boundary",
         calculate_discount(10.0, 49), 441.0)

# Test 4: quantity = 50 (boundary: enters 25% discount bracket)
# 10.0 * 50 * 0.75 = 375.0
run_test("quantity 50: 25% discount boundary",
         calculate_discount(10.0, 50), 375.0)

# Test 5: quantity = 1 (minimum quantity, no discount)
# 5.5 * 1 = 5.5
run_test("quantity 1: no discount, single item",
         calculate_discount(5.5, 1), 5.5)

# Test 6: quantity = 100 (well into 25% discount bracket)
# 20.0 * 100 * 0.75 = 1500.0
run_test("quantity 100: 25% discount applied",
         calculate_discount(20.0, 100), 1500.0)