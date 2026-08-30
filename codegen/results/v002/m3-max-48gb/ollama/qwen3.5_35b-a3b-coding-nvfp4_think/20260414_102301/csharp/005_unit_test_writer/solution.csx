// CalculateDiscount function
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test runner
void RunTest(string description, double price, int quantity, double expected) {
    double actual = CalculateDiscount(price, quantity);
    // Use approximate comparison for floating point
    bool passed = Math.Abs(actual - expected) < 0.0001;
    
    if (passed) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
    }
}

// Test cases with boundary conditions
RunTest("Test 1: quantity = 9 (no discount)", 100.0, 9, 900.0);
RunTest("Test 2: quantity = 10 (10% discount)", 100.0, 10, 900.0);
RunTest("Test 3: quantity = 49 (10% discount)", 100.0, 49, 4410.0);
RunTest("Test 4: quantity = 50 (25% discount)", 100.0, 50, 3750.0);