double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void RunTest(string description, double expected, double actual) {
    if (Math.Abs(expected - actual) < 1e-9)
        Console.WriteLine($"PASS: {description}");
    else
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
}

// quantity < 10: no discount
RunTest("quantity 9 (no discount)", 10.0 * 9, CalculateDiscount(10.0, 9));

// quantity == 10: first boundary, 10% discount
RunTest("quantity 10 (10% discount)", 10.0 * 10 * 0.9, CalculateDiscount(10.0, 10));

// quantity == 49: just below second boundary, 10% discount
RunTest("quantity 49 (10% discount)", 10.0 * 49 * 0.9, CalculateDiscount(10.0, 49));

// quantity == 50: second boundary, 25% discount
RunTest("quantity 50 (25% discount)", 10.0 * 50 * 0.75, CalculateDiscount(10.0, 50));

// quantity == 1: minimum, no discount
RunTest("quantity 1 (no discount)", 5.0 * 1, CalculateDiscount(5.0, 1));

// quantity == 100: well above second boundary, 25% discount
RunTest("quantity 100 (25% discount)", 2.5 * 100 * 0.75, CalculateDiscount(2.5, 100));
