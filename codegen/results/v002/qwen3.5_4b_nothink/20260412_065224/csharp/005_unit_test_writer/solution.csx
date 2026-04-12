using System;
using System.Collections.Generic;
using System.Linq;

// Copy the function definition
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases
var testCases = new List<(string description, double expected, double actual)> {
    ("quantity 9 (no discount)", 100.0, CalculateDiscount(100.0, 9)),
    ("quantity 10 (90% off)", 900.0, CalculateDiscount(100.0, 10)),
    ("quantity 49 (90% off)", 4410.0, CalculateDiscount(100.0, 49)),
    ("quantity 50 (75% off)", 3750.0, CalculateDiscount(100.0, 50))
};

foreach (var testCase in testCases) {
    var (description, expected, actual) = testCase;
    if (Math.Abs(expected - actual) < 0.001) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
    }
}