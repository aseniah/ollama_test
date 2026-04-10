using System;
using System.Collections.Generic;
using System.Text.Json;

// Function definition
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test runner
void RunTest(string description, double price, int quantity, double expected) {
    double result = CalculateDiscount(price, quantity);
    // Allow for small floating point differences
    if (Math.Abs(result - expected) < 0.0001) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {result})");
    }
}

// Test cases covering boundary conditions
RunTest("Quantity 9 (No discount)", 10.0, 9, 90.0);
RunTest("Quantity 10 (First boundary - 10% discount)", 10.0, 10, 90.0);
RunTest("Quantity 49 (Second boundary - 10% discount)", 10.0, 49, 441.0);
RunTest("Quantity 50 (Third boundary - 25% discount)", 10.0, 50, 375.0);

// Additional test case for extra coverage
RunTest("Quantity 60 (Above threshold - 25% discount)", 10.0, 60, 450.0);