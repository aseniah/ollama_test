using System;
using System.Collections.Generic;

// The function to be tested
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Define a structure for test cases
var testCases = new List<(string description, double price, int quantity, double expected)>
{
    ("Below threshold (9 units)", 10.0, 9, 90.0),
    ("Boundary: Exactly threshold (10 units)", 10.0, 10, 90.0), // 10 * 10 * 0.9
    ("Mid-tier (49 units)", 10.0, 49, 441.0), // 10 * 49 * 0.9
    ("Boundary: Next threshold (50 units)", 10.0, 50, 375.0), // 10 * 50 * 0.75
    ("Large quantity (100 units)", 10.0, 100, 750.0)
};

int passed = 0;
int failed = 0;

foreach (var test in testCases)
{
    double actual = CalculateDiscount(test.price, test.quantity);
    
    // Use a small epsilon for double comparison to avoid floating point precision issues
    bool isSuccess = Math.Abs(actual - test.expected) < 0.000001;

    if (isSuccess)
    {
        Console.WriteLine($"PASS: {test.description}");
        passed++;
    }
    else
    {
        Console.WriteLine($"FAIL: {test.description} (expected: {test.expected}, got: {actual})");
        failed++;
    }
}

Console.WriteLine($"\nTest Summary: {passed} passed, {failed} failed.");