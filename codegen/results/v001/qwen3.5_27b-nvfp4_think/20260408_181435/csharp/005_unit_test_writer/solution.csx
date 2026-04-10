using System;
using System.Collections.Generic;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

var tests = new List<(string Description, double Price, int Quantity, double Expected)> {
    ("Quantity 9 (no discount)", 100.0, 9, 900.0),
    ("Quantity 10 (10% off)", 100.0, 10, 900.0),
    ("Quantity 49 (10% off)", 100.0, 49, 4410.0),
    ("Quantity 50 (25% off)", 100.0, 50, 3750.0)
};

foreach (var test in tests) {
    double actual = CalculateDiscount(test.Price, test.Quantity);
    
    // Use tolerance for floating point comparisons
    if (Math.Abs(actual - test.Expected) < 0.001) {
        Console.WriteLine($"PASS: {test.Description}");
    } else {
        Console.WriteLine($"FAIL: {test.Description} (expected: {test.Expected}, got: {actual})");
    }
}