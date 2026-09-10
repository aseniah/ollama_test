using System;
using System.Collections.Generic;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Define a test case structure
var testCases = new List<(string Description, double Price, int Quantity, double Expected)>
{
    ("Below first threshold (9 items)", 10.0, 9, 90.0),
    ("Exact first threshold boundary (10 items)", 10.0, 10, 90.0), // 10 * 10 * 0.9
    ("Below second threshold (49 items)", 10.0, 49, 441.0),       // 49 * 10 * 0.9
    ("Exact second threshold boundary (50 items)", 10.0, 50, 375.0) // 50 * 10 * 0.75
};

foreach (var test in testCases)
{
    double actual = CalculateDiscount(test.Price, test.Quantity);
    
    // Using a small epsilon for double comparison to avoid precision issues
    if (Math.Abs(actual - test.Expected) < 0.0001)
    {
        Console.WriteLine($"PASS: {test.Description}");
    }
    else
    {
        Console.WriteLine($"FAIL: {test.Description} (expected: {test.Expected}, got: {actual})");
    }
}