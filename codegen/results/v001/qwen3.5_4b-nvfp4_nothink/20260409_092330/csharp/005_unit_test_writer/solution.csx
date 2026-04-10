using System;
using System.Text.Json;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

int passedCount = 0;
int totalTests = 4;

bool Test(string desc, double expected, double got) {
    if (Math.Abs(expected - got) < 1e-6) {
        Console.WriteLine($"PASS: {desc}");
        passedCount++;
        return true;
    } else {
        Console.WriteLine($"FAIL: {desc} (expected: {expected}, got: {got})");
        return false;
    }
}

// Test Case 1: quantity < 10 (boundary)
Test("Quantity 9 (discount applies):", CalculateDiscount(10.0, 9), 9.0);

// Test Case 2: quantity = 10 (boundary between tiers)
double expected2 = 10.0 * 10 * 0.9;
double got2 = CalculateDiscount(10.0, 10);
Test("Quantity 10 (tier boundary):", expected2, got2);

// Test Case 3: quantity = 49 (boundary)
double expected3 = 10.0 * 49 * 0.9;
double got3 = CalculateDiscount(10.0, 49);
Test("Quantity 49 (tier boundary):", expected3, got3);

// Test Case 4: quantity > 50 (bulk discount)
double expected4 = 10.0 * 50 * 0.75;
double got4 = CalculateDiscount(10.0, 50);
Test("Quantity 50 (bulk discount):", expected4, got4);

Console.WriteLine($"\nResults: {passedCount}/{totalTests} tests passed.");