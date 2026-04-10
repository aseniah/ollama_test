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
    print("=" * 60)
    print("Testing calculate_discount function")
    print("=" * 60)
    
    test_cases = [
        # (price, quantity, expected_result, description)
        (10.0, 9, 90.0, "quantity < 10: no discount"),
        (10.0, 10, 90.0, "boundary: quantity == 10 (10% off)"),
        (10.0, 49, 441.0, "boundary: quantity == 49 (10% off)"),
        (10.0, 50, 375.0, "boundary: quantity == 50 (25% off)"),
    ]
    
    failed = []
    passed_count = 0
    
    for price, quantity, expected, description in test_cases:
        actual = calculate_discount(price, quantity)
        
        # Use floating point comparison with small tolerance
        if abs(actual - expected) < 1e-6:
            status = "PASS"
            pass_count += 1
        else:
            status = "FAIL"
            failed.append((expected, actual))
        
        print(f"{status}: {description}")
        if status == "FAIL":
            print(f"   Expected: {expected}, Got: {actual}")
    
    print("=" * 60)
    print(f"Results: {passed_count} passed, {len(failed)} failed")
    print("=" * 60)
    
    return len(failed) == 0


if __name__ == "__main__":
    test_calculate_discount()