double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void RunTest(string description, double price, int quantity, double expected) {
    double actual = CalculateDiscount(price, quantity);
    
    // Using a small epsilon for double comparison to avoid precision issues
    if (Math.Abs(actual - expected) < 0.0001) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
    }
}

// Test cases covering boundaries
RunTest("Quantity just below 10 (no discount)", 10.0, 9, 90.0);
RunTest("Quantity exactly 10 (10% discount)", 10.0, 10, 90.0);
RunTest("Quantity just below 50 (10% discount)", 10.0, 49, 441.0);
RunTest("Quantity exactly 50 (25% discount)", 10.0, 50, 375.0);

// Additional test case for general functionality
RunTest("Quantity high above 50 (25% discount)", 10.0, 100, 750.0);