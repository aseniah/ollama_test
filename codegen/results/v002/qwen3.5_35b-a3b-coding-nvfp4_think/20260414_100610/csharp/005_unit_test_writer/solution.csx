// CalculateDiscount function
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases
void RunTest(string description, double price, int quantity, double expected) {
    double actual = CalculateDiscount(price, quantity);
    // Use a small epsilon for floating point comparison
    double epsilon = 0.0001;
    bool passed = Math.Abs(actual - expected) < epsilon;
    
    if (passed) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
    }
}

// Test case 1: quantity = 9 (no discount - below threshold)
RunTest("quantity = 9 (no discount)", 100.0, 9, 900.0);

// Test case 2: quantity = 10 (10% discount - boundary)
RunTest("quantity = 10 (10% discount)", 100.0, 10, 900.0);

// Test case 3: quantity = 49 (10% discount - below 50 threshold)
RunTest("quantity = 49 (10% discount)", 100.0, 49, 4410.0);

// Test case 4: quantity = 50 (25% discount - boundary)
RunTest("quantity = 50 (25% discount)", 100.0, 50, 3750.0);

// Additional test case for edge case
RunTest("quantity = 0 (edge case)", 100.0, 0, 0.0);