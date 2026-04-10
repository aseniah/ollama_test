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

if __name__ == "__main__":
    BASE_PRICE = 100.0
    
    test_cases = [
        (9, "quantity < 10 (no discount)"),
        (10, "quantity 10-49 (10% off)"),
        (49, "quantity 10-49 (10% off)"),
        (50, "quantity >= 50 (25% off)")
    ]

    for qty, desc in test_cases:
        expected = 0.0
        if qty < 10:
            expected = BASE_PRICE * qty
        elif qty < 50:
            expected = BASE_PRICE * qty * 0.9
        else:
            expected = BASE_PRICE * qty * 0.75
        
        result = calculate_discount(BASE_PRICE, qty)
        
        # Floating point comparison with tolerance
        if abs(expected - result) < 1e-6:
            print(f"PASS: {desc}")
        else:
            print(f"FAIL: {desc} (expected: {expected}, got: {result})")