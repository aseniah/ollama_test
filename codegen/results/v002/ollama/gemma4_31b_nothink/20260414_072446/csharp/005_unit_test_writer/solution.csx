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
    ("Boundary: Just below first discount tier (9)", 10.0, 9, 90.0),
    ("Boundary: Start of 10% discount tier (10)", 10.0, 10, 90.0), // 10 * 10 * 0.9 = 90
    ("Boundary: Just below second discount tier (49)", 10.0, 49, 441.0), // 49 * 10 * 0.9 = 441
    ("Boundary: Start of 25% discount tier (50)", 10.0, 50, 375.0), // 50 * 10 * 0.75 = 375
    ("General: Low quantity", 20.0, 2, 40.0),
    ("General: High quantity", 20.0, 100, 1500.0) // 100 * 20 * 0.75 = 1500
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