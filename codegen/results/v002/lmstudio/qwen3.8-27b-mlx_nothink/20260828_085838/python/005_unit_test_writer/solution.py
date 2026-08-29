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


def is_close(expected: float, actual: float) -> bool:
    return abs(expected - actual) <= 1e-9


def main() -> None:
    test_cases = [
        ("quantity 9 is below discount threshold", 10.0, 9, 90.0),
        ("quantity 10 applies 10% discount", 10.0, 10, 90.0),
        ("quantity 49 applies 10% discount", 10.0, 49, 441.0),
        ("quantity 50 applies 25% discount", 10.0, 50, 375.0),
        ("quantity 0 returns 0.0", 10.0, 0, 0.0),
        ("quantity 1 applies no discount", 10.0, 1, 10.0),
    ]

    for description, price, quantity, expected in test_cases:
        actual = calculate_discount(price, quantity)

        if is_close(expected, actual):
            print(f"PASS: {description}")
        else:
            print(f"FAIL: {description} (expected: {expected}, got: {actual})")


if __name__ == "__main__":
    main()