// CalculateDiscount function
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test helper function
void TestCase(string description, double price, int quantity, double expected) {
    double result = CalculateDiscount(price, quantity);
    // Use small epsilon for floating point comparison
    double epsilon = 0.0001;
    if (Math.Abs(result - expected) < epsilon) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {result})");
    }
}

// Run test cases - covering boundary conditions
TestCase("Quantity 9 (no discount)", 10.0, 9, 90.0);
TestCase("Quantity 10 (10% discount starts)", 10.0, 10, 90.0);
TestCase("Quantity 49 (10% discount)", 10.0, 49, 441.0);
TestCase("Quantity 50 (25% discount starts)", 10.0, 50, 375.0);