// The function to test
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Helper method for testing
void TestCase(double price, int quantity, double expected, string description) {
    double actual = CalculateDiscount(price, quantity);
    // Use a small epsilon for floating point comparison
    bool passed = Math.Abs(actual - expected) < 0.001;
    
    if (passed) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
    }
}

// Test cases with boundary conditions
TestCase(10.0, 9, 90.0, "Quantity 9 (no discount)");
TestCase(10.0, 10, 90.0, "Quantity 10 (10% discount)");
TestCase(10.0, 49, 441.0, "Quantity 49 (10% discount)");
TestCase(10.0, 50, 375.0, "Quantity 50 (25% discount)");

// Additional test cases for thoroughness
TestCase(20.0, 5, 100.0, "Quantity 5 (no discount, different price)");
TestCase(20.0, 100, 1500.0, "Quantity 100 (25% discount)");