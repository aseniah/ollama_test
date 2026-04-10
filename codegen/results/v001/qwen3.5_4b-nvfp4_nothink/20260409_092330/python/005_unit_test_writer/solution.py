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

# Test cases
test_cases = [
    {"price": 100, "quantity": 9, "expected": 900, "description": "Single item (q=9): no discount"},
    {"price": 100, "quantity": 10, "expected": 900, "description": "Minimum bulk discount (q=10): 10% off"},
    {"price": 100, "quantity": 49, "expected": 4410, "description": "Maximum bulk discount tier (q=49): 10% off"},
    {"price": 100, "quantity": 50, "expected": 3750, "description": "Premium bulk discount (q=50): 25% off"}
]

passed = 0
failed = 0

for tc in test_cases:
    price = tc["price"]
    quantity = tc["quantity"]
    expected = tc["expected"]
    description = tc["description"]
    
    result = calculate_discount(price, quantity)
    
    # Handle floating point comparison with a small epsilon for safety
    if abs(result - expected) < 1e-9:
        print(f"PASS: {description}")
        passed += 1
    else:
        print(f"FAIL: {description} (expected: {expected}, got: {result})")
        failed += 1

print(f"\nSummary: {passed} passed, {failed} failed")