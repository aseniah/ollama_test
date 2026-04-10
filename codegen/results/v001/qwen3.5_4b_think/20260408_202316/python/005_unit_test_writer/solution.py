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


# Test configuration
test_cases = [
    {"price": 100.0, "quantity": 9, "expected": 900.0, "desc": "quantity 9 (no discount)"},
    {"price": 100.0, "quantity": 10, "expected": 900.0, "desc": "quantity 10 (10% off)"},
    {"price": 100.0, "quantity": 49, "expected": 4410.0, "desc": "quantity 49 (10% off)"},
    {"price": 100.0, "quantity": 50, "expected": 3750.0, "desc": "quantity 50 (25% off)"},
]

# Tolerance for float comparison
EPSILON = 1e-6

for test in test_cases:
    result = calculate_discount(test["price"], test["quantity"])
    
    # Check if result is within acceptable float precision
    if abs(result - test["expected"]) < EPSILON:
        print(f"PASS: {test['desc']}")
    else:
        print(f"FAIL: {test['desc']} (expected: {test['expected']}, got: {result})")