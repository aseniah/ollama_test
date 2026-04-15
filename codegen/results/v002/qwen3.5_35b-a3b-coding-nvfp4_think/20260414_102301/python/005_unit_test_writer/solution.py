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


def run_test(description: str, price: float, quantity: int, expected: float):
    """Run a single test case and print PASS or FAIL."""
    result = calculate_discount(price, quantity)
    tolerance = 0.01  # Allow small floating point differences
    if abs(result - expected) < tolerance:
        print(f"PASS: {description}")
    else:
        print(f"FAIL: {description} (expected: {expected}, got: {result})")


def main():
    """Run test cases for calculate_discount function."""
    print("Testing calculate_discount function:")
    print("-" * 50)

    # Test case 1: quantity = 9 (boundary - no discount)
    run_test(
        "quantity=9 (no discount)",
        price=100.0,
        quantity=9,
        expected=900.0
    )

    # Test case 2: quantity = 10 (boundary - 10% discount)
    run_test(
        "quantity=10 (10% discount)",
        price=100.0,
        quantity=10,
        expected=900.0
    )

    # Test case 3: quantity = 49 (boundary - 10% discount)
    run_test(
        "quantity=49 (10% discount)",
        price=100.0,
        quantity=49,
        expected=4410.0
    )

    # Test case 4: quantity = 50 (boundary - 25% discount)
    run_test(
        "quantity=50 (25% discount)",
        price=100.0,
        quantity=50,
        expected=3750.0
    )

    print("-" * 50)


if __name__ == "__main__":
    main()