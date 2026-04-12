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

RunTest("quantity 9 (no discount)", 10.0 * 9, CalculateDiscount(10.0, 9));
RunTest("quantity 10 (10% discount)", 10.0 * 10 * 0.9, CalculateDiscount(10.0, 10));
RunTest("quantity 49 (10% discount)", 10.0 * 49 * 0.9, CalculateDiscount(10.0, 49));
RunTest("quantity 50 (25% discount)", 10.0 * 50 * 0.75, CalculateDiscount(10.0, 50));
RunTest("quantity 1 (no discount)", 5.0 * 1, CalculateDiscount(5.0, 1));
RunTest("quantity 100 (25% discount)", 20.0 * 100 * 0.75, CalculateDiscount(20.0, 100));
