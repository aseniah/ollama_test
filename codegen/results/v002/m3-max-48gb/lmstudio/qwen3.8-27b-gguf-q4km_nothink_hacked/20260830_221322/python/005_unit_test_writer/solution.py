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
        # (price, quantity, description)
        (10.0, 9, "Boundary: quantity 9 (no discount)"),
        (10.0, 10, "Boundary: quantity 10 (10% discount)"),
        (10.0, 49, "Boundary: quantity 49 (10% discount)"),
        (10.0, 50, "Boundary: quantity 50 (25% discount)"),
        (5.0, 5, "Below discount threshold"),
        (5.0, 25, "Mid-range discount"),
        (5.0, 100, "High quantity discount"),
        (0.0, 10, "Zero price"),
        (100.0, 1, "Single item no discount"),
    ]
    
    for price, quantity, desc in test_cases:
        result = calculate_discount(price, quantity)
        
        # Calculate expected value
        if quantity < 10:
            expected = price * quantity
        elif quantity < 50:
            expected = price * quantity * 0.9
        else:
            expected = price * quantity * 0.75
        
        # Use a small tolerance for floating point comparison
        if abs(result - expected) < 1e-9:
            print(f"PASS: {desc}")
        else:
            print(f"FAIL: {desc} (expected: {expected}, got: {result}")


if __name__ == "__main__":
    test_calculate_discount()