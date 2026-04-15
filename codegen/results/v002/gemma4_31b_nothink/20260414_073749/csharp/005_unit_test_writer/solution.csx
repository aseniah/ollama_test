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
    ("Below first threshold (9 items)", 100.0, 9, 900.0),
    ("Exactly first threshold (10 items)", 100.0, 10, 900.0), // 10 * 100 * 0.9
    ("Below second threshold (49 items)", 100.0, 49, 4410.0), // 49 * 100 * 0.9
    ("Exactly second threshold (50 items)", 100.0, 50, 3750.0), // 50 * 100 * 0.75
    ("Very large quantity (100 items)", 10.0, 100, 750.0)      // 100 * 10 * 0.75
};

foreach (var test in testCases)
{
    double result = CalculateDiscount(test.Price, test.Quantity);
    
    // Using a small epsilon for floating point comparison
    if (Math.Abs(result - test.Expected) < 0.001)
    {
        Console.WriteLine($"PASS: {test.Description}");
    }
    else
    {
        Console.WriteLine($"FAIL: {test.Description} (expected: {test.Expected}, got: {result})");
    }
}