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
    test_cases = [
        (9, 10.0, "Quantity 9 (below 10, no discount)", 90.0),
        (10, 10.0, "Quantity 10 (10-49, 10% off)", 90.0),
        (49, 10.0, "Quantity 49 (10-49, 10% off)", 441.0),
        (50, 10.0, "Quantity 50 (>= 50, 25% off)", 375.0),
    ]
    
    for qty, price, desc, expected in test_cases:
        result = calculate_discount(price, qty)
        # Use a small epsilon for float comparison to handle precision differences
        if abs(result - expected) < 1e-9:
            print(f"PASS: {desc}")
        else:
            print(f"FAIL: {desc} (expected: {expected}, got: {result})")