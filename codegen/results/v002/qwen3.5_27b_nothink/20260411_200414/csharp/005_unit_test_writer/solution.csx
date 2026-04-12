#r "System.Text.Json"

using System;
using System.Text.Json;
using System.Text.Json.Nodes;

// The function to test
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Helper to run a test case
void RunTest(string description, double price, int quantity, double expected) {
    double result = CalculateDiscount(price, quantity);
    // Use a small epsilon for floating point comparison
    double epsilon = 0.0001;
    bool passed = Math.Abs(result - expected) < epsilon;

    if (passed) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {result})");
    }
}

// Test Case 1: Quantity < 10 (No discount)
// Price: 10, Qty: 9 -> Expected: 90.0
RunTest("Quantity 9 (No discount)", 10.0, 9, 90.0);

// Test Case 2: Quantity == 10 (Boundary: 10% discount starts)
// Price: 10, Qty: 10 -> Expected: 100 * 0.9 = 90.0
RunTest("Quantity 10 (Boundary: 10% discount)", 10.0, 10, 90.0);

// Test Case 3: Quantity == 49 (Boundary: Just before 25% discount)
// Price: 10, Qty: 49 -> Expected: 490 * 0.9 = 441.0
RunTest("Quantity 49 (Boundary: Just before 25% discount)", 10.0, 49, 441.0);

// Test Case 4: Quantity == 50 (Boundary: 25% discount starts)
// Price: 10, Qty: 50 -> Expected: 500 * 0.75 = 375.0
RunTest("Quantity 50 (Boundary: 25% discount)", 10.0, 50, 375.0);