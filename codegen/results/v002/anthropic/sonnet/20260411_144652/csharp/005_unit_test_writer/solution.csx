double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void Check(string description, double expected, double actual) {
    if (Math.Abs(expected - actual) < 1e-9)
        Console.WriteLine($"PASS: {description}");
    else
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
}

Check("quantity 9 (no discount)", 10.0 * 9, CalculateDiscount(10.0, 9));
Check("quantity 10 (10% discount starts)", 10.0 * 10 * 0.9, CalculateDiscount(10.0, 10));
Check("quantity 49 (10% discount, upper boundary)", 10.0 * 49 * 0.9, CalculateDiscount(10.0, 49));
Check("quantity 50 (25% discount starts)", 10.0 * 50 * 0.75, CalculateDiscount(10.0, 50));
Check("quantity 1 (single unit, no discount)", 10.0 * 1, CalculateDiscount(10.0, 1));
Check("quantity 100 (25% discount, large order)", 10.0 * 100 * 0.75, CalculateDiscount(10.0, 100));
