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
    ("Boundary condition: quantity 9 (No discount)", 100.0, 9, 900.0),
    ("Boundary condition: quantity 10 (Start of 10% discount)", 100.0, 10, 900.0),
    ("Boundary condition: quantity 49 (End of 10% discount)", 100.0, 49, 4410.0),
    ("Boundary condition: quantity 50 (Start of 25% discount)", 100.0, 50, 3750.0),
    ("Standard case: quantity 100 (25% discount)", 10.0, 100, 750.0),
    ("Edge case: quantity 1 (No discount)", 50.0, 1, 50.0)
};

// Use a small epsilon for double comparison to handle floating point precision issues
const double epsilon = 1e-9;

foreach (var test in testCases)
{
    double actual = CalculateDiscount(test.Price, test.Quantity);
    bool isPassed = Math.Abs(actual - test.Expected) < epsilon;

    if (isPassed)
    {
        Console.WriteLine($"PASS: {test.Description}");
    }
    else
    {
        // Format expected and actual to avoid scientific notation for simple numbers
        Console.WriteLine($"FAIL: {test.Description} (expected: {test.Expected}, got: {actual})");
    }
}