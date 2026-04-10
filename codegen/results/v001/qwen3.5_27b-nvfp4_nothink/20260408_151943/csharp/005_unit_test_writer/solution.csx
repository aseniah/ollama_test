#r "System.Text.Json"
using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

// The function to test
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test runner logic
void RunTest(string description, double price, int quantity, double expected) {
    double actual = CalculateDiscount(price, quantity);
    
    // Use a small epsilon for floating point comparison to handle precision issues
    bool passed = Math.Abs(actual - expected) < 0.0001;

    if (passed) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
    }
}

// Execute Tests
Console.WriteLine("Running Bulk Discount Tests...");
Console.WriteLine("-------------------------------");

// Case 1: Boundary below 10 (No discount)
RunTest("Quantity 9 (Below 10, No Discount)", 100.0, 9, 900.0);

// Case 2: Boundary at 10 (10% discount starts)
RunTest("Quantity 10 (Start of 10% Discount)", 100.0, 10, 900.0);

// Case 3: Boundary below 50 (Still 10% discount)
RunTest("Quantity 49 (Below 50, 10% Discount)", 100.0, 49, 4410.0);

// Case 4: Boundary at 50 (25% discount starts)
RunTest("Quantity 50 (Start of 25% Discount)", 100.0, 50, 3750.0);

// Additional Edge Cases
RunTest("Quantity 0", 100.0, 0, 0.0);
RunTest("Large Quantity (Over 50)", 20.5, 100, 1537.5);

Console.WriteLine("-------------------------------");
Console.WriteLine("Tests completed.");