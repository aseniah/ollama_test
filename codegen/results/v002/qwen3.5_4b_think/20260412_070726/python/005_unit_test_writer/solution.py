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
            "name": "Quantity 9 (no discount boundary)",
            "price": 10.0,
            "quantity": 9,
            "expected": 90.0,
        },
        {
            "name": "Quantity 10 (10% discount boundary)",
            "price": 10.0,
            "quantity": 10,
            "expected": 90.0,
        },
        {
            "name": "Quantity 49 (10% discount boundary)",
            "price": 10.0,
            "quantity": 49,
            "expected": 441.0,
        },
        {
            "name": "Quantity 50 (25% discount boundary)",
            "price": 10.0,
            "quantity": 50,
            "expected": 375.0,
        },
    ]
    
    for test in test_cases:
        name = test["name"]
        price = test["price"]
        quantity = test["quantity"]
        expected = test["expected"]
        result = calculate_discount(price, quantity)
        tolerance = 0.001  # For floating point comparison
        
        if abs(result - expected) < tolerance:
            print(f"PASS: {name}")
        else:
            print(f"FAIL: {name} (expected: {expected}, got: {result})")


if __name__ == "__main__":
    run_tests()