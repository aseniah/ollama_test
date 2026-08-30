double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void RunTest(int quantity, double price, double expected, string description) {
    double result = CalculateDiscount(price, quantity);
    // Use a small epsilon for floating point comparison safety, though inputs here are clean
    if (Math.Abs(result - expected) < 0.0001) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {result})");
    }
}

// Test Case 1: Quantity 9 (Boundary - No discount)
// Expected: 10.0 * 9 = 90.0
RunTest(9, 10.0, 90.0, "Quantity 9 (No discount)");

// Test Case 2: Quantity 10 (Boundary - 10% discount)
// Expected: 10.0 * 10 * 0.9 = 90.0
RunTest(10, 10.0, 90.0, "Quantity 10 (10% discount)");

// Test Case 3: Quantity 49 (Boundary - 10% discount)
// Expected: 10.0 * 49 * 0.9 = 441.0
RunTest(49, 10.0, 441.0, "Quantity 49 (10% discount)");

// Test Case 4: Quantity 50 (Boundary - 25% discount)
// Expected: 10.0 * 50 * 0.75 = 375.0
RunTest(50, 10.0, 375.0, "Quantity 50 (25% discount)");