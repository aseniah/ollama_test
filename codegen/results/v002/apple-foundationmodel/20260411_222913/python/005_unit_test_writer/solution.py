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
    # Test cases with descriptions
    test_cases = [
        ("test_quantity_9", 9, 10.0, 9.0),  # No discount, quantity < 10
        ("test_quantity_10", 10, 10.0, 10.0),  # 10% off, quantity 10-49
        ("test_quantity_49", 49, 10.0, 44.5),  # 10% off, quantity 10-49
        ("test_quantity_50", 50, 10.0, 50.0),  # 25% off, quantity >= 50
        ("test_quantity_99", 99, 10.0, 89.25),  # 25% off, quantity >= 50
    ]

    for description, quantity, expected, result in test_cases:
        if description == "test_quantity_9":
            got = calculate_discount(10.0, 9)
        elif description == "test_quantity_10":
            got = calculate_discount(10.0, 10)
        elif description == "test_quantity_49":
            got = calculate_discount(10.0, 49)
        elif description == "test_quantity_50":
            got = calculate_discount(10.0, 50)
        elif description == "test_quantity_99":
            got = calculate_discount(10.0, 99)
        else:
            continue
        
        if got == expected:
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {got})")

test_calculate_discount()