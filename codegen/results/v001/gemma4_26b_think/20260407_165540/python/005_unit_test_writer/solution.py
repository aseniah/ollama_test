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
    # Define test cases: (price, quantity, expected_output, description)
    test_cases = [
        # Boundary: quantity 9 (No discount)
        (100.0, 9, 900.0, "Quantity 9 (No discount)"),
        
        # Boundary: quantity 10 (10% discount starts)
        (100.0, 10, 900.0, "Quantity 10 (10% discount)"),
        
        # Boundary: quantity 49 (End of 10% discount range)
        (100.0, 49, 4410.0, "Quantity 49 (10% discount)"),
        
        # Boundary: quantity 50 (25% discount starts)
        (100.0, 50, 3750.0, "Quantity 50 (25% discount)"),
        
        # Extra case: Large quantity
        (10.0, 100, 750.0, "Quantity 100 (25% discount)"),
        
        # Extra case: Small quantity
        (50.0, 1, 50.0, "Quantity 1 (No discount)")
    ]

    for price, quantity, expected, description in test_cases:
        actual = calculate_discount(price, quantity)
        
        # Using round to handle potential floating point precision issues in comparison
        # Though for these specific numbers, direct comparison works.
        if round(actual, 2) == round(expected, 2):
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {actual})")

if __name__ == "__main__":
    run_tests()