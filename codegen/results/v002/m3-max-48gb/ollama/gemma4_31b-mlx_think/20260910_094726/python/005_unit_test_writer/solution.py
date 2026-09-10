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
    # Define test cases: (price, quantity, expected_result, description)
    test_cases = [
        (10.0, 9, 90.0, "Boundary case: quantity 9 (no discount)"),
        (10.0, 10, 90.0, "Boundary case: quantity 10 (10% discount)"),
        (10.0, 49, 441.0, "Boundary case: quantity 49 (10% discount)"),
        (10.0, 50, 375.0, "Boundary case: quantity 50 (25% discount)"),
        (100.0, 5, 500.0, "Standard case: small quantity"),
        (100.0, 100, 7500.0, "Standard case: large quantity"),
    ]

    for price, quantity, expected, description in test_cases:
        actual = calculate_discount(price, quantity)
        
        # Using round to avoid floating point precision issues during comparison
        if round(actual, 2) == round(expected, 2):
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {actual})")

if __name__ == "__main__":
    run_tests()