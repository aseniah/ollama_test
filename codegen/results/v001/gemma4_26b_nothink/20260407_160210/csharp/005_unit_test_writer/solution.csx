using System;
using System.Collections.Generic;

// The function to be tested
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test Case definition
record TestCase(string Description, double Price, int Quantity, double Expected);

// Define the test suite
var testCases = new List<TestCase>
{
    new TestCase("Boundary: quantity 9 (No discount)", 100.0, 9, 900.0),
    new TestCase("Boundary: quantity 10 (10% discount applied)", 100.0, 10, 900.0),
    new TestCase("Boundary: quantity 49 (10% discount applied)", 100.0, 49, 4410.0),
    new TestCase("Boundary: quantity 50 (25% discount applied)", 100.0, 50, 3750.0),
    new TestCase("Standard: quantity 100 (25% discount applied)", 10.0, 100, 750.0),
    new TestCase("Standard: quantity 1 (No discount)", 50.0, 1, 50.0)
};

int passed = 0;
int failed = 0;

Console.WriteLine($"Running {testCases.Count} tests...\n");

foreach (var test in testCases)
{
    double actual = CalculateDiscount(test.Price, test.Quantity);
    
    // Using a small epsilon for double comparison to avoid precision issues
    bool isSuccess = Math.Abs(actual - test.Expected) < 0.000001;

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

Console.WriteLine($"\nSummary: {passed} passed, {failed} failed.");

if (failed > 0)
{
    Environment.Exit(1);
}