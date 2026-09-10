double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void RunTest(string description, double price, int quantity, double expected) {
    double result = CalculateDiscount(price, quantity);
    
    // Using a small epsilon for double comparison to avoid precision issues
    if (Math.Abs(result - expected) < 0.0001) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {result})");
    }
}

// Test Price
double testPrice = 10.0;

// Boundary Condition Tests
RunTest("Quantity 9 (No discount)", testPrice, 9, 90.0);
RunTest("Quantity 10 (10% discount boundary)", testPrice, 10, 90.0); // 10 * 10 * 0.9 = 90
RunTest("Quantity 49 (10% discount boundary)", testPrice, 49, 441.0); // 10 * 49 * 0.9 = 441
RunTest("Quantity 50 (25% discount boundary)", testPrice, 50, 375.0); // 10 * 50 * 0.75 = 375

// Additional test case
RunTest("Quantity 5 (Small order)", testPrice, 5, 50.0);