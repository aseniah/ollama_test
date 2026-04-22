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

# Define test cases covering the required boundary conditions
test_cases = [
    {"price": 10.0, "quantity": 9, "expected": 90.0, "description": "Quantity 9 (no discount)"},
    {"price": 10.0, "quantity": 10, "expected": 90.0, "description": "Quantity 10 (10% off)"},
    {"price": 10.0, "quantity": 49, "expected": 441.0, "description": "Quantity 49 (10% off)"},
    {"price": 10.0, "quantity": 50, "expected": 375.0, "description": "Quantity 50 (25% off)"},
]

# Execute tests and print results
for tc in test_cases:
    result = calculate_discount(tc["price"], tc["quantity"])
    if result == tc["expected"]:
        print(f"PASS: {tc['description']}")
    else:
        print(f"FAIL: {tc['description']} (expected: {tc['expected']}, got: {result})")