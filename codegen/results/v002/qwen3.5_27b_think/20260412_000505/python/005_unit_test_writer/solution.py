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
    # Define test cases: (price, quantity, expected_result, description)
    # Including boundary conditions: 9, 10, 49, 50
    test_cases = [
        (10.0, 9, 90.0, "Boundary: Quantity 9 (No Discount)"),
        (10.0, 10, 90.0, "Boundary: Quantity 10 (10% Discount Start)"),
        (10.0, 49, 441.0, "Boundary: Quantity 49 (10% Discount End)"),
        (10.0, 50, 375.0, "Boundary: Quantity 50 (25% Discount Start)"),
    ]

    for price, quantity, expected, description in test_cases:
        result = calculate_discount(price, quantity)
        
        # Check if result matches expected (with tolerance for floating point arithmetic)
        is_pass = abs(result - expected) < 1e-9
        
        if is_pass:
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {result})")

if __name__ == "__main__":
    main()