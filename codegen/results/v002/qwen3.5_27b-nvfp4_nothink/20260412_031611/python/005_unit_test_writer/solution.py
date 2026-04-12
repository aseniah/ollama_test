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
    test_cases = [
        {
            "description": "Quantity below threshold (9 units): no discount",
            "price": 10.0,
            "quantity": 9,
            "expected": 10.0 * 9  # 90.0
        },
        {
            "description": "Lower boundary of 10% discount (10 units)",
            "price": 10.0,
            "quantity": 10,
            "expected": 10.0 * 10 * 0.9  # 90.0
        },
        {
            "description": "Upper boundary of 10% discount (49 units)",
            "price": 2.0,
            "quantity": 49,
            "expected": 2.0 * 49 * 0.9  # 88.2
        },
        {
            "description": "Lower boundary of 25% discount (50 units)",
            "price": 100.0,
            "quantity": 50,
            "expected": 100.0 * 50 * 0.75  # 3750.0
        },
        {
            "description": "High quantity above 50 (60 units)",
            "price": 10.0,
            "quantity": 60,
            "expected": 10.0 * 60 * 0.75  # 450.0
        }
    ]

    for i, case in enumerate(test_cases):
        price = case["price"]
        quantity = case["quantity"]
        expected = case["expected"]
        description = case["description"]
        
        result = calculate_discount(price, quantity)
        
        # Use a small epsilon for float comparison to handle precision issues
        if abs(result - expected) < 1e-9:
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {result})")

if __name__ == "__main__":
    run_tests()