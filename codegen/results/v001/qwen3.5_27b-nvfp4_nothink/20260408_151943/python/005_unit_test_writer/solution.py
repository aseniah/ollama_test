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
    print("Running tests for calculate_discount function...")

    test_cases = [
        {
            "description": "Quantity below discount threshold (9 items)",
            "price": 10.0,
            "quantity": 9,
            "expected": 10.0 * 9 * 1.0
        },
        {
            "description": "Lower boundary of first discount tier (10 items)",
            "price": 10.0,
            "quantity": 10,
            "expected": 10.0 * 10 * 0.9
        },
        {
            "description": "Upper boundary of first discount tier (49 items)",
            "price": 10.0,
            "quantity": 49,
            "expected": 10.0 * 49 * 0.9
        },
        {
            "description": "Lower boundary of highest discount tier (50 items)",
            "price": 10.0,
            "quantity": 50,
            "expected": 10.0 * 50 * 0.75
        },
    ]

    for i, test in enumerate(test_cases, 1):
        result = calculate_discount(test["price"], test["quantity"])
        
        # Use a small epsilon for floating point comparison to handle precision issues
        if abs(result - test["expected"]) < 0.0001:
            print(f"PASS: {test['description']}")
        else:
            print(f"FAIL: {test['description']} (expected: {test['expected']}, got: {result})")

if __name__ == "__main__":
    run_tests()