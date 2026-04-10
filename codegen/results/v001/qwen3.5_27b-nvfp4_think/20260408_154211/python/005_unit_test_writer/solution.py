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
    # Define test cases as tuples: (price, quantity, expected_result, description)
    # Boundary conditions included: 9, 10, 49, 50
    test_cases = [
        (100.0, 9, 900.0, "Quantity boundary (< 10): No discount"),
        (100.0, 10, 900.0, "Quantity boundary (>= 10): 10% discount"),
        (100.0, 49, 4410.0, "Quantity boundary (< 50): 10% discount"),
        (100.0, 50, 3750.0, "Quantity boundary (>= 50): 25% discount")
    ]

    for price, quantity, expected, description in test_cases:
        actual = calculate_discount(price, quantity)
        
        # Using a small epsilon for floating point comparison safety
        if abs(actual - expected) < 1e-9:
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {actual})")

if __name__ == "__main__":
    run_tests()