using System;

// The function to test
double CalculateDiscount(double price, int quantity)
{
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void RunTest(double price, int quantity, double expected, string description)
{
    double result = CalculateDiscount(price, quantity);
    
    // Use a small epsilon for floating-point comparison to handle precision issues
    if (Math.Abs(result - expected) < 0.0001)
    {
        Console.WriteLine($"PASS: {description}");
    }
    else
    {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {result})");
    }
}

// Test Case 1: Quantity 9 (Boundary: Less than 10)
// Expected: 10 * 9 = 90
RunTest(10, 9, 90, "Quantity 9: No discount");

// Test Case 2: Quantity 10 (Boundary: 10% discount threshold)
// Expected: 10 * 10 * 0.9 = 90
RunTest(10, 10, 90, "Quantity 10: 10% discount");

// Test Case 3: Quantity 49 (Boundary: Just below 25% discount)
// Expected: 10 * 49 * 0.9 = 441
RunTest(10, 49, 441, "Quantity 49: 10% discount");

// Test Case 4: Quantity 50 (Boundary: 25% discount threshold)
// Expected: 10 * 50 * 0.75 = 375
RunTest(10, 50, 375, "Quantity 50: 25% discount");