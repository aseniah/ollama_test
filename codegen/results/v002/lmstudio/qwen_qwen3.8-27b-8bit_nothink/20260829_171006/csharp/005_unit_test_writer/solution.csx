double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void RunTest(string description, double price, int quantity, double expected) {
    double got = CalculateDiscount(price, quantity);
    // Use a small tolerance for floating point comparison
    bool pass = Math.Abs(got - expected) < 0.0001;
    if (pass) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {got})");
    }
}

// Test cases with boundary conditions
// 1. quantity = 9 (just below 10, no discount)
RunTest("Quantity 9 (no discount)", 10.0, 9, 10.0 * 9 * 1.0);

// 2. quantity = 10 (boundary: starts 10% discount)
RunTest("Quantity 10 (10% discount starts)", 10.0, 10, 10.0 * 10 * 0.9);

// 3. quantity = 49 (just below 50, still 10% discount)
RunTest("Quantity 49 (10% discount)", 10.0, 49, 10.0 * 49 * 0.9);

// 4. quantity = 50 (boundary: starts 25% discount)
RunTest("Quantity 50 (25% discount starts)", 10.0, 50, 10.0 * 50 * 0.75);