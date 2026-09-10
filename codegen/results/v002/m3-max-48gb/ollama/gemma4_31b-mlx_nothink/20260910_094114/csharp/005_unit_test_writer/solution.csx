using System;
using System.Collections.Generic;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Define a simple structure for test cases
var tests = new List<(string Description, double Price, int Quantity, double Expected)>
{
    ("Below 10 units (9)", 10.0, 9, 90.0),
    ("Boundary 10 units (10% discount start)", 10.0, 10, 90.0), // 10 * 10 * 0.9 = 90
    ("Below 50 units (49)", 10.0, 49, 441.0),                // 10 * 49 * 0.9 = 441
    ("Boundary 50 units (25% discount start)", 10.0, 50, 375.0) // 10 * 50 * 0.75 = 375
};

foreach (var test in tests)
{
    double result = CalculateDiscount(test.Price, test.Quantity);
    
    // Using a small epsilon for double comparison to avoid precision issues
    if (Math.Abs(result - test.Expected) < 0.0001)
    {
        Console.WriteLine($"PASS: {test.Description}");
    }
    else
    {
        Console.WriteLine($"FAIL: {test.Description} (expected: {test.Expected}, got: {result})");
    }
}