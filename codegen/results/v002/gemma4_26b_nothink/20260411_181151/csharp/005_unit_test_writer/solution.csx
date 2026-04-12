using System;
using System.Collections.Generic;

// The function to be tested
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Define a test case structure
record TestCase(string Description, double Price, int Quantity, double Expected);

// Prepare the list of test cases including boundary conditions
var testCases = new List<TestCase>
{
    new TestCase("Boundary: quantity 9 (No discount)", 100.0, 9, 900.0),
    new TestCase("Boundary: quantity 10 (10% discount)", 100.0, 10, 900.0),
    new TestCase("Boundary: quantity 49 (10% discount)", 100.0, 49, 4410.0),
    new TestCase("Boundary: quantity 50 (25% discount)", 100.0, 50, 3750.0),
    new TestCase("General: quantity 100 (25% discount)", 10.0, 100, 750.0),
    new TestCase("General: quantity 1 (No discount)", 50.0, 1, 50.0)
};

// Iterate through test cases and evaluate
foreach (var test in testCases)
{
    double actual = CalculateDiscount(test.Price, test.Quantity);
    
    // Using a small epsilon for floating point comparison
    const double epsilon = 0.000001;
    bool passed = Math.Abs(actual - test.Expected) < epsilon;

    if (passed)
    {
        Console.WriteLine($"PASS: {test.Description}");
    }
    else
    {
        // Formatting numbers to 2 decimal places for readability in output
        Console.WriteLine($"FAIL: {test.Description} (expected: {test.Expected:F2}, got: {actual:F2})");
    }
}