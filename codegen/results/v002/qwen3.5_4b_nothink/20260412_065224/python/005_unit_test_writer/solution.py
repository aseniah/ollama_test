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
    # Test case 1: quantity = 9 (no discount)
    price = 10.0
    quantity = 9
    expected = price * quantity
    result = calculate_discount(price, quantity)
    description = f"Test quantity={quantity}, price={price}: expected {expected}, got {result}"
    if abs(result - expected) < 1e-9:
        print(f"PASS: {description}")
    else:
        print(f"FAIL: {description} (expected: {expected}, got: {result})")

    # Test case 2: quantity = 10 (10% off)
    price = 10.0
    quantity = 10
    expected = price * quantity * 0.9
    result = calculate_discount(price, quantity)
    description = f"Test quantity={quantity}, price={price}: expected {expected}, got {result}"
    if abs(result - expected) < 1e-9:
        print(f"PASS: {description}")
    else:
        print(f"FAIL: {description} (expected: {expected}, got: {result})")

    # Test case 3: quantity = 49 (10% off)
    price = 20.0
    quantity = 49
    expected = price * quantity * 0.9
    result = calculate_discount(price, quantity)
    description = f"Test quantity={quantity}, price={price}: expected {expected}, got {result}"
    if abs(result - expected) < 1e-9:
        print(f"PASS: {description}")
    else:
        print(f"FAIL: {description} (expected: {expected}, got: {result})")

    # Test case 4: quantity = 50 (25% off)
    price = 30.0
    quantity = 50
    expected = price * quantity * 0.75
    result = calculate_discount(price, quantity)
    description = f"Test quantity={quantity}, price={price}: expected {expected}, got {result}"
    if abs(result - expected) < 1e-9:
        print(f"PASS: {description}")
    else:
        print(f"FAIL: {description} (expected: {expected}, got: {result})")


if __name__ == "__main__":
    test_calculate_discount()