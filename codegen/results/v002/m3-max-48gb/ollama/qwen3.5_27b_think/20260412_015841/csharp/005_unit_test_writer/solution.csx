using System;
using System.Collections.Generic;

// Copy the CalculateDiscount function
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Define test cases with boundary conditions (9, 10, 49, 50)
var testCases = new List<(string description, double price, int quantity, double expected)>() {
    ("No discount for quantity 9 (below 10)", 10.0, 9, 90.0),
    ("10% discount for quantity 10 (boundary)", 10.0, 10, 90.0),
    ("10% discount for quantity 49 (below 50)", 10.0, 49, 441.0),
    ("25% discount for quantity 50 (boundary)", 10.0, 50, 375.0)
};

// Run tests
foreach (var testCase in testCases) {
    var result = CalculateDiscount(testCase.price, testCase.quantity);
    var description = testCase.description;
    var expected = testCase.expected;
    
    // Use small tolerance for floating point comparison
    var tolerance = 0.001;
    
    if (Math.Abs(result - expected) < tolerance) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {result})");
    }
}