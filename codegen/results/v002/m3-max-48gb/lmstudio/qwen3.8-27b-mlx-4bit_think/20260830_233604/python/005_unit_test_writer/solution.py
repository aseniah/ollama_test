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


def run_test(description: str, price: float, quantity: int, expected: float) -> None:
    result = calculate_discount(price, quantity)
    if abs(result - expected) < 1e-9:
        print(f"PASS: {description}")
    else:
        print(f"FAIL: {description} (expected: {expected}, got: {result})")


if __name__ == "__main__":
    # Boundary: quantity just below 10 (no discount)
    run_test("quantity=9 (no discount), price=10", 10.0, 9, 10.0 * 9)

    # Boundary: quantity=10 (10% discount starts)
    run_test("quantity=10 (10% discount), price=10", 10.0, 10, 10.0 * 10 * 0.9)

    # Boundary: quantity=49 (10% discount, just below 50)
    run_test("quantity=49 (10% discount), price=10", 10.0, 49, 10.0 * 49 * 0.9)

    # Boundary: quantity=50 (25% discount starts)
    run_test("quantity=50 (25% discount), price=10", 10.0, 50, 10.0 * 50 * 0.75)

    # Additional: quantity well above 50
    run_test("quantity=100 (25% discount), price=5.5", 5.5, 100, 5.5 * 100 * 0.75)

    # Additional: quantity=1 (minimum)
    run_test("quantity=1 (no discount), price=3.0", 3.0, 1, 3.0 * 1)