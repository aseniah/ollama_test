using System;
using System.Collections.Generic;

// The function to be tested
double CalculateDiscount(double price, int quantity)
{
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Helper class to define test cases
var testCases = new List<(string Description, double Price, int Quantity, double Expected)>
{
    ("Quantity 9 (No discount)", 100.0, 9, 900.0),
    ("Quantity 10 (Boundary - 10% discount)", 100.0, 10, 900.0),
    ("Quantity 49 (Boundary - 10% discount)", 100.0, 49, 4410.0),
    ("Quantity 50 (Boundary - 25% discount)", 100.0, 50, 3750.0),
    ("Quantity 100 (Large quantity)", 10.0, 100, 750.0),
    ("Quantity 1 (Smallest valid)", 50.0, 1, 50.0)
};

int passCount = 0;
int failCount = 0;

foreach (var test in testCases)
{
    double actual = CalculateDiscount(test.Price, test.Quantity);
    
    // Using a small epsilon for floating point comparison
    bool isPassed = Math.Abs(actual - test.Expected) < 0.000001;

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