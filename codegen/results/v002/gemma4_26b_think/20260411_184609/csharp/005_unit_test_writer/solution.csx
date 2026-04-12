using System;

// The function to be tested
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Helper method to execute and format test results
void RunTest(string description, double price, int quantity, double expected)
{
    double actual = CalculateDiscount(price, quantity);
    
    // Using a small epsilon for double comparison to handle precision issues
    bool passed = Math.Abs(actual - expected) < 1e-10;

    if (passed)
    {
        Console.WriteLine($"PASS: {description}");
    }
    else
    {
        // Formatting numbers to avoid long decimal strings in output
        Console.WriteLine($"FAIL: {description} (expected: {expected:G}, got: {actual:G})");
    }
}

// Test Cases
// Boundary Condition: quantity 9 (Should have no discount)
RunTest("Quantity 9 (No discount)", 100.0, 9, 900.0);

// Boundary Condition: quantity 10 (Should trigger 10% discount)
// Note: 100 * 10 * 0.9 = 900. The value happens to be same as quantity 9, 
// but the logic path changes.
RunTest("Quantity 10 (10% discount)", 100.0, 10, 900.0);

// Boundary Condition: quantity 49 (Should trigger 10% discount)
RunTest("Quantity 49 (10% discount)", 100.0, 49, 4410.0);

// Boundary Condition: quantity 50 (Should trigger 25% discount)
RunTest("Quantity 50 (25% discount)", 100.0, 50, 3750.0);

// Extra Test: Large quantity
RunTest("Quantity 100 (25% discount)", 10.0, 100, 750.0);