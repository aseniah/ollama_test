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

def test_calculate_discount():
    test_cases = [
        {"price": 10, "quantity": 9,  "expected": 90,    "description": "Quantity less than 10 (no discount)"},
        {"price": 10, "quantity": 10, "expected": 90,    "description": "Minimum quantity for 10% discount"},
        {"price": 10, "quantity": 49, "expected": 44.1,  "description": "Maximum quantity for 10% discount"},
        {"price": 10, "quantity": 50, "expected": 37.5,  "description": "Minimum quantity for 25% discount"},
        {"price": 50, "quantity": 20, "expected": 90,    "description": "Mid-range price and quantity"},
    ]
    
    passed_count = 0
    total_cases = len(test_cases)

    print("Running tests for calculate_discount...\n")
    
    for case in test_cases:
        p = case["price"]
        q = case["quantity"]
        exp = case["expected"]
        desc = case["description"]
        
        result = calculate_discount(p, q)
        # Compare with a small epsilon for float comparisons
        if abs(result - exp) < 1e-9:
            print(f"PASS: {desc}")
            passed_count += 1
        else:
            print(f"FAIL: {desc} (expected: {exp}, got: {result})")

    print("\n" + "="*30)
    print(f"Tests completed: {passed_count}/{total_cases} passed")
    if passed_count == total_cases:
        print("All tests passed successfully!")

if __name__ == "__main__":
    test_calculate_discount()