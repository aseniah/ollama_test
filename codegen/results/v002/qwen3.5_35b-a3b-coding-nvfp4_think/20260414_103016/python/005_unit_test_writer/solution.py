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


def run_test(price: float, quantity: int, expected: float, description: str) -> None:
    """Run a single test case and print the result."""
    actual = calculate_discount(price, quantity)
    if abs(actual - expected) < 0.0001:  # Allow for floating point tolerance
        print(f"PASS: {description}")
    else:
        print(f"FAIL: {description} (expected: {expected}, got: {actual})")


def main():
    """Run all test cases for calculate_discount function."""
    # Test case 1: quantity < 10 (no discount)
    run_test(100.0, 9, 900.0, "No discount for quantity 9")

    # Test case 2: quantity = 10 (10% off - boundary)
    run_test(100.0, 10, 900.0, "10% discount for quantity 10")

    # Test case 3: quantity = 49 (10% off - boundary)
    run_test(100.0, 49, 4410.0, "10% discount for quantity 49")

    # Test case 4: quantity = 50 (25% off - boundary)
    run_test(100.0, 50, 3750.0, "25% discount for quantity 50")


if __name__ == "__main__":
    main()