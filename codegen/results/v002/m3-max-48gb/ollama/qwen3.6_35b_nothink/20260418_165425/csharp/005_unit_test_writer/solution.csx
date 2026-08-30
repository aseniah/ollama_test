double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases covering boundary conditions
var testCases = new (double price, int quantity, double expected, string description)[] {
    (10.0, 9, 90.0, "quantity < 10 (boundary: 9)"),
    (10.0, 10, 90.0, "quantity == 10 (boundary: 10, 10% off)"),
    (10.0, 49, 441.0, "quantity < 50 (boundary: 49, 10% off)"),
    (10.0, 50, 375.0, "quantity == 50 (boundary: 50, 25% off)"),
    (10.0, 51, 382.5, "quantity > 50 (25% off)"),
    (100.0, 1, 100.0, "small quantity (no discount)"),
    (0.0, 10, 0.0, "zero price"),
    (10.0, 100, 750.0, "large quantity (25% off)")
};

foreach (var tc in testCases) {
    double result = CalculateDiscount(tc.price, tc.quantity);
    // Use a small epsilon for floating point comparison
    double epsilon = 1e-9;
    double diff = Math.Abs(result - tc.expected);
    if (diff < epsilon) {
        Console.WriteLine($"PASS: {tc.description}");
    } else {
        Console.WriteLine($"FAIL: {tc.description} (expected: {tc.expected}, got: {result})");
    }
}