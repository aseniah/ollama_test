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
    def print_result(expected, got, description):
        if expected == got:
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {got})")
    
    # Test case 1: quantity < 10 (boundary: quantity = 9)
    price = 10.0
    expected = 90.0
    got = calculate_discount(price, 9)
    print_result(expected, got, f"price=10.0, quantity=9 (no discount)")

    # Test case 2: quantity 10-49 (boundary: quantity = 10)
    price = 10.0
    expected = 90.0
    got = calculate_discount(price, 10)
    print_result(expected, got, f"price=10.0, quantity=10 (10% off)")

    # Test case 3: quantity 10-49 (boundary: quantity = 49)
    price = 100.0
    expected = 810.0
    got = calculate_discount(price, 49)
    print_result(expected, got, f"price=100.0, quantity=49 (10% off)")

    # Test case 4: quantity >= 50 (boundary: quantity = 50)
    price = 20.0
    expected = 300.0
    got = calculate_discount(price, 50)
    print_result(expected, got, f"price=20.0, quantity=50 (25% off)")

if __name__ == "__main__":
    test_calculate_discount()