// CalculateDiscount function
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test runner
void RunTest(string description, double price, int quantity, double expected) {
    double result = CalculateDiscount(price, quantity);
    
    // Allow small floating point tolerance
    bool passed = Math.Abs(result - expected) < 0.001;
    
    if (passed) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {result})");
    }
}

// Test cases with boundary conditions
RunTest("Quantity 9 (no discount)", 100, 9, 900);
RunTest("Quantity 10 (10% discount boundary)", 100, 10, 900);
RunTest("Quantity 49 (10% discount)", 100, 49, 4410);
RunTest("Quantity 50 (25% discount boundary)", 100, 50, 3750);