using System;

// The function provided for testing
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Helper method to run and format test results
void RunTest(string description, double price, int quantity, double expected) {
    double actual = CalculateDiscount(price, quantity);
    
    // Using a small epsilon to handle potential floating point precision issues
    bool passed = Math.Abs(actual - expected) < 1e-10;

    if (passed) {
        Console.WriteLine($"PASS: {description}");
    } else {
        // Format expected and actual to avoid excessive decimal places in the output
        Console.WriteLine($"FAIL: {description} (expected: {expected:G}, got: {actual:G})");
    }
}

// Test Case 1: Boundary Condition - Quantity 9 (Should have no discount)
RunTest("Quantity 9 (No discount applied)", 100.0, 9, 900.0);

// Test Case 2: Boundary Condition - Quantity 10 (Threshold for 10% discount)
RunTest("Quantity 10 (10% discount applied)", 100.0, 10, 900.0);

// Test Case 3: Boundary Condition - Quantity 49 (Still in 10% discount range)
RunTest("Quantity 49 (10% discount applied)", 100.0, 49, 4410.0);

// Test Case 4: Boundary Condition - Quantity 50 (Threshold for 25% discount)
RunTest("Quantity 50 (25% discount applied)", 100.0, 50, 3750.0);

// Test Case 5: Arbitrary value within high discount range
RunTest("Quantity 100 (25% discount applied)", 10.0, 100, 750.0);