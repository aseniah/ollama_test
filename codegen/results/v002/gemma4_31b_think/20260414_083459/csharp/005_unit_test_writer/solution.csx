using System;
using System.Collections.Generic;

// The function to be tested
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test Case structure
var testCases = new[]
{
    new { Price = 100.0, Quantity = 9,  Expected = 900.0, Description = "Quantity just below 10 (no discount)" },
    new { Price = 100.0, Quantity = 10, Expected = 900.0, Description = "Quantity exactly 10 (10% discount)" },
    new { Price = 100.0, Quantity = 49, Expected = 4410.0, Description = "Quantity just below 50 (10% discount)" },
    new { Price = 100.0, Quantity = 50, Expected = 3750.0, Description = "Quantity exactly 50 (25% discount)" },
    new { Price = 20.0,  Quantity = 5,  Expected = 100.0, Description = "Small quantity (no discount)" },
    new { Price = 10.0,  Quantity = 60, Expected = 450.0, Description = "Large quantity (25% discount)" }
};

foreach (var tc in testCases)
{
    double result = CalculateDiscount(tc.Price, tc.Quantity);
    
    // Using a small epsilon for double comparison to avoid precision issues
    if (Math.Abs(result - tc.Expected) < 0.0001)
    {
        Console.WriteLine($"PASS: {tc.Description}");
    }
    else
    {
        Console.WriteLine($"FAIL: {tc.Description} (expected: {tc.Expected}, got: {result})");
    }
}