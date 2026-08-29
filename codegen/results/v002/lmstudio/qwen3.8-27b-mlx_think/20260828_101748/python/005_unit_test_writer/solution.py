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


def is_close(a: float, b: float) -> bool:
    return abs(a - b) <= 1e-9 * max(1.0, abs(a), abs(b))


tests = [
    ("quantity 9 gets no discount", 10.0, 9, 90.0),
    ("quantity 10 gets 10% discount", 10.0, 10, 90.0),
    ("quantity 49 gets 10% discount", 10.0, 49, 441.0),
    ("quantity 50 gets 25% discount", 10.0, 50, 375.0),
]

for description, price, quantity, expected in tests:
    got = calculate_discount(price, quantity)
    if is_close(got, expected):
        print(f"PASS: {description}")
    else:
        print(f"FAIL: {description} (expected: {expected}, got: {got})")