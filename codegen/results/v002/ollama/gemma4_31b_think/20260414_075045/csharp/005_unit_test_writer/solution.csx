using System;
using System.Collections.Generic;

// The function to be tested
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test case definition
var testCases = new[]
{
    new { Description = "Below first threshold (quantity 9)", Price = 100.0, Quantity = 9, Expected = 900.0 },
    new { Description = "Exactly first threshold (quantity 10)", Price = 100.0, Quantity = 10, Expected = 900.0 }, // 100 * 10 * 0.9
    new { Description = "Below second threshold (quantity 49)", Price = 100.0, Quantity = 49, Expected = 4410.0 }, // 100 * 49 * 0.9
    new { Description = "Exactly second threshold (quantity 50)", Price = 100.0, Quantity = 50, Expected = 3750.0 }, // 100 * 50 * 0.75
    new { Description = "Well above second threshold (quantity 100)", Price = 10.0, Quantity = 100, Expected = 750.0 } // 10 * 100 * 0.75
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