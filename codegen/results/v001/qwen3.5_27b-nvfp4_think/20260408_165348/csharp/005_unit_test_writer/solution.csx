using System;
using System.Text.Json;
using System.Text.Json.Nodes;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void Test(string description, double price, int quantity, double expected) {
    double actual = CalculateDiscount(price, quantity);
    bool passed = Math.Abs(actual - expected) < 0.01;
    
    if (passed) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected:F2}, got: {actual:F2})");
    }
}

Console.WriteLine("Testing CalculateDiscount function...\n");

// Test case 1: Quantity less than 10 (no discount) - boundary condition
Test("Quantity 9, price $100 (no discount)", 100, 9, 900);

// Test case 2: Quantity exactly 10 (10% discount starts) - boundary condition
Test("Quantity 10, price $100 (10% discount)", 100, 10, 900);

// Test case 3: Quantity just below 50 (still 10% discount) - boundary condition
Test("Quantity 49, price $100 (10% discount)", 100, 49, 4410);

// Test case 4: Quantity exactly 50 (25% discount starts) - boundary condition
Test("Quantity 50, price $100 (25% discount)", 100, 50, 3750);

// Additional test cases for completeness
Test("Quantity 1, price $50 (no discount)", 50, 1, 50);

Test("Quantity 25, price $80 (10% discount)", 80, 25, 1800);

Test("Quantity 100, price $100 (25% discount)", 100, 100, 7500);

Console.WriteLine("\nAll tests completed.");