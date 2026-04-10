using System;
using System.Text.Json;
using System.Text.Json.Nodes;

// The requested function copied into the script
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Helper to run a test and print the result
void RunTest(string description, double price, int quantity, double expected) {
    double actual = CalculateDiscount(price, quantity);
    
    // Use a small epsilon for float comparison safety, though simple == usually works for these clean decimals
    bool passed = Math.Abs(actual - expected) < 0.0001;

    if (passed) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
    }
}

// Test Case 1: Boundary condition (quantity 9) - No discount
// Price: 100, Qty: 9 -> 100 * 9 = 900
RunTest("Quantity 9 (no discount)", 100, 9, 900);

// Test Case 2: Boundary condition (quantity 10) - 10% discount
// Price: 100, Qty: 10 -> 100 * 10 * 0.9 = 900
RunTest("Quantity 10 (10% discount)", 100, 10, 900);

// Test Case 3: Boundary condition (quantity 49) - 10% discount
// Price: 100, Qty: 49 -> 100 * 49 * 0.9 = 4410
RunTest("Quantity 49 (10% discount)", 100, 49, 4410);

// Test Case 4: Boundary condition (quantity 50) - 25% discount
// Price: 100, Qty: 50 -> 100 * 50 * 0.75 = 3750
RunTest("Quantity 50 (25% discount)", 100, 50, 3750);

// Test Case 5: Additional check (Quantity 100) - 25% discount
// Price: 50, Qty: 100 -> 50 * 100 * 0.75 = 3750
RunTest("Quantity 100 (25% discount, diff price)", 50, 100, 3750);