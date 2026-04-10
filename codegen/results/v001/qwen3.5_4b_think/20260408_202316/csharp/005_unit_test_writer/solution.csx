using System;
using System.Collections.Generic;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases: {description, price, quantity, expected}
var testCases = new List<(string desc, double price, int quantity, double expected)> {
    ("Quantity 9 - no discount", 100.0, 9, 900.0),
    ("Quantity 10 - 10% discount", 100.0, 10, 900.0),
    ("Quantity 49 - 10% discount", 100.0, 49, 4410.0),
    ("Quantity 50 - 25% discount", 100.0, 50, 3750.0),
};

foreach (var (desc, price, quantity, expected) in testCases) {
    var actual = CalculateDiscount(price, quantity);
    var match = (Math.Abs(actual - expected) < 0.01);
    
    if (match) {
        Console.WriteLine($"PASS: {desc}");
    } else {
        Console.WriteLine($"FAIL: {desc} (expected: {expected}, got: {actual})");
    }
}