using System;
using System.Collections.Generic;

// The function to test
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Define a structure for our test cases
var testCases = new List<(string Description, double Price, int Quantity, double Expected)>
{
    ("Below threshold (9 items)", 10.0, 9, 90.0),
    ("Lower boundary threshold (10 items)", 10.0, 10, 90.0), // 10 * 10 * 0.9
    ("Upper boundary threshold (49 items)", 10.0, 49, 441.0), // 10 * 49 * 0.9
    ("Bulk threshold (50 items)", 10.0, 50, 375.0), // 10 * 50 * 0.75
    ("Large quantity (100 items)", 10.0, 100, 750.0)
};

int passCount = 0;
int failCount = 0;

foreach (var test in testCases)
{
    double actual = CalculateDiscount(test.Price, test.Quantity);
    
    // Using a small epsilon for double comparison to avoid floating point precision issues
    bool isPassed = Math.Abs(actual - test.Expected) < 0.0001;

    if (isPassed)
    {
        Console.WriteLine($"PASS: {test.Description}");
        passCount++;
    }
    else
    {
        Console.WriteLine($"FAIL: {test.Description} (expected: {test.Expected}, got: {actual})");
        failCount++;
    }
}

Console.WriteLine("\nSummary:");
Console.WriteLine($"Total: {testCases.Count}, Passed: {passCount}, Failed: {failCount}");

if (failCount > 0)
{
    Environment.Exit(1);
}