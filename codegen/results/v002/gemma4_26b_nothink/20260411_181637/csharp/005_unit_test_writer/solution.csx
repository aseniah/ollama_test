using System;
using System.Collections.Generic;

// The function to be tested
double CalculateDiscount(double price, int quantity)
{
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Define a structure for our test cases
var testCases = new List<(string Description, double Price, int Quantity, double Expected)>
{
    ("Below threshold (quantity 9)", 100.0, 9, 900.0),
    ("Boundary threshold (quantity 10)", 100.0, 10, 900.0), // 10 * 100 * 0.9
    ("Below upper threshold (quantity 49)", 100.0, 49, 4410.0), // 49 * 100 * 0.9
    ("Upper threshold (quantity 50)", 100.0, 50, 3750.0), // 50 * 100 * 0.75
    ("Large quantity (quantity 100)", 10.0, 100, 750.0)
};

int passed = 0;
int failed = 0;

foreach (var test in testCases)
{
    double actual = CalculateDiscount(test.Price, test.Quantity);
    
    // Using a small epsilon for double comparison to avoid precision issues
    double epsilon = 0.000001;
    bool isSuccess = Math.Abs(actual - test.Expected) < epsilon;

    if (isSuccess)
    {
        Console.WriteLine($"PASS: {test.Description}");
        passed++;
    }
    else
    {
        Console.WriteLine($"FAIL: {test.Description} (expected: {test.Expected}, got: {actual})");
        failed++;
    }
}

Console.WriteLine("\nSummary:");
Console.WriteLine($"Total: {testCases.Count}, Passed: {passed}, Failed: {failed}");