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
        {"price": 10.0, "quantity": 9, "expected": 90.0, "description": "Quantity 9 (no discount)"},
        {"price": 10.0, "quantity": 10, "expected": 90.0, "description": "Quantity 10 (10% discount threshold)"},
        {"price": 10.0, "quantity": 49, "expected": 441.0, "description": "Quantity 49 (10% discount, just below 50)"},
        {"price": 10.0, "quantity": 50, "expected": 375.0, "description": "Quantity 50 (25% discount threshold)"},
        {"price": 20.0, "quantity": 75, "expected": 1125.0, "description": "Quantity 75 (25% discount)"},
        {"price": 5.0, "quantity": 5, "expected": 25.0, "description": "Small price and quantity (no discount)"},
    ]

    for i, test in enumerate(test_cases):
        result = calculate_discount(test["price"], test["quantity"])
        if abs(result - test["expected"]) < 1e-9:
            print(f"PASS: {test['description']}")
        else:
            print(f"FAIL: {test['description']} (expected: {test['expected']}, got: {result})")


if __name__ == "__main__":
    run_tests()