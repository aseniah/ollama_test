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
    price = 100.0
    
    # Test case 1: quantity < 10 (no discount), boundary: 9
    result = calculate_discount(price, 9)
    expected = price * 9
    status = "PASS" if abs(result - expected) < 1e-9 else "FAIL"
    description = f"No discount for quantity 9. Expected: {expected}, Got: {result}"
    print(f"{status}: {description}")
    
    # Test case 2: 10 <= quantity < 50 (10% off), boundary: 10
    result = calculate_discount(price, 10)
    expected = price * 10 * 0.9
    status = "PASS" if abs(result - expected) < 1e-9 else "FAIL"
    description = f"10% discount for quantity 10. Expected: {expected}, Got: {result}"
    print(f"{status}: {description}")
    
    # Test case 3: 10 <= quantity < 50 (10% off), boundary: 49
    result = calculate_discount(price, 49)
    expected = price * 49 * 0.9
    status = "PASS" if abs(result - expected) < 1e-9 else "FAIL"
    description = f"10% discount for quantity 49. Expected: {expected}, Got: {result}"
    print(f"{status}: {description}")
    
    # Test case 4: quantity >= 50 (25% off), boundary: 50
    result = calculate_discount(price, 50)
    expected = price * 50 * 0.75
    status = "PASS" if abs(result - expected) < 1e-9 else "FAIL"
    description = f"25% discount for quantity 50. Expected: {expected}, Got: {result}"
    print(f"{status}: {description}")

if __name__ == "__main__":
    test_calculate_discount()