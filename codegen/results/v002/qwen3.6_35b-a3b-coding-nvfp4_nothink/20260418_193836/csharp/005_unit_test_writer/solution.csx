using System;
using System.Collections.Generic;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void Test(string description, double price, int quantity, double expected) {
    double got = CalculateDiscount(price, quantity);
    double tolerance = 0.0001;
    if (Math.Abs(got - expected) < tolerance) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {got})");
    }
}

// Test case 1: quantity < 10 (quantity = 9)
Test("quantity = 9, price = 100", 100.0, 9, 100.0 * 9);

// Test case 2: quantity = 10 (boundary, 10 <= quantity < 50, 10% discount)
Test("quantity = 10, price = 100", 100.0, 10, 100.0 * 10 * 0.9);

// Test case 3: quantity = 49 (boundary, 10 <= quantity < 50, 10% discount)
Test("quantity = 49, price = 100", 100.0, 49, 100.0 * 49 * 0.9);

// Test case 4: quantity = 50 (boundary, quantity >= 50, 25% discount)
Test("quantity = 50, price = 100", 100.0, 50, 100.0 * 50 * 0.75);