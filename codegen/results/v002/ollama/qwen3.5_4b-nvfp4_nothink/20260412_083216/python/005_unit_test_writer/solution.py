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

def main():
    # Define expected values based on the logic
    test_cases = [
        {"price": 10.0, "quantity": 9, "expected": 90.0, "desc": "no discount (boundary < 10)"},
        {"price": 10.0, "quantity": 10, "expected": 90.0, "desc": "start of 10% discount"},
        {"price": 10.0, "quantity": 49, "expected": 441.0, "desc": "end of 10% discount"},
        {"price": 10.0, "quantity": 50, "expected": 375.0, "desc": "start of 25% discount"}
    ]

    print("Running tests for calculate_discount...")

    for test in test_cases:
        expected = test["expected"]
        desc = test["desc"]
        qty = test["quantity"]
        price = test["price"]
        
        actual = calculate_discount(price, qty)
        
        # Using a small epsilon for float comparison safety, though inputs are clean here
        if abs(actual - expected) < 1e-9:
            print(f"PASS: {desc}")
        else:
            print(f"FAIL: {desc} (expected: {expected}, got: {actual})")

if __name__ == "__main__":
    main()