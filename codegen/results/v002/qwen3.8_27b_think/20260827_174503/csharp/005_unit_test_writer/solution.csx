double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

var tests = new (double price, int quantity, double expected, string desc)[] {
    (100.0,  9, 100.0 * 9,          "quantity=9 (below 10, no discount)"),
    (100.0, 10, 100.0 * 10 * 0.9,   "quantity=10 (boundary: exactly 10, 10% discount)"),
    (100.0, 49, 100.0 * 49 * 0.9,   "quantity=49 (below 50, 10% discount)"),
    (100.0, 50, 100.0 * 50 * 0.75,  "quantity=50 (boundary: exactly 50, 25% discount)"),
};

foreach (var (price, quantity, expected, desc) in tests) {
    double result = CalculateDiscount(price, quantity);
    double diff = Math.Abs(result - expected);
    if (diff < 1e-9) {
        Console.WriteLine($"PASS: {desc}");
    } else {
        Console.WriteLine($"FAIL: {desc} (expected: {expected}, got: {result})");
    }
}