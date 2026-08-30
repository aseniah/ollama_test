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
    """Run a single test case and print result."""
    result = calculate_discount(price, quantity)
    # Use approximate equality for float comparison
    if abs(result - expected) < 0.01:
        print(f"PASS: {description}")
    else:
        print(f"FAIL: {description} (expected: {expected}, got: {result})")


def main():
    """Run all test cases."""
    # Test case 1: quantity < 10 (no discount)
    run_test(
        "No discount for quantity < 10",
        price=100.0,
        quantity=9,
        expected=900.0  # 100 * 9 = 900
    )

    # Test case 2: quantity = 10 (10% discount boundary)
    run_test(
        "10% discount for quantity = 10",
        price=100.0,
        quantity=10,
        expected=900.0  # 100 * 10 * 0.9 = 900
    )

    # Test case 3: quantity = 49 (10% discount boundary)
    run_test(
        "10% discount for quantity = 49",
        price=50.0,
        quantity=49,
        expected=2205.0  # 50 * 49 * 0.9 = 2205
    )

    # Test case 4: quantity = 50 (25% discount boundary)
    run_test(
        "25% discount for quantity = 50",
        price=50.0,
        quantity=50,
        expected=1875.0  # 50 * 50 * 0.75 = 1875
    )

    # Additional test cases for coverage
    run_test(
        "No discount for quantity = 5",
        price=200.0,
        quantity=5,
        expected=1000.0  # 200 * 5 = 1000
    )

    run_test(
        "10% discount for quantity = 25",
        price=100.0,
        quantity=25,
        expected=2250.0  # 100 * 25 * 0.9 = 2250
    )

    run_test(
        "25% discount for quantity = 100",
        price=100.0,
        quantity=100,
        expected=7500.0  # 100 * 100 * 0.75 = 7500
    )


if __name__ == "__main__":
    main()