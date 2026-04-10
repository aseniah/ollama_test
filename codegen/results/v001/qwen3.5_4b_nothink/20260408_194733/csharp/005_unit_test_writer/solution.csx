using System;
using System.Collections.Generic;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

double GetExpected(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

List<(string desc, double expected, int qty, double price)> tests = new List<(string, double, int, double)> {
    ("Quantity 9 (no discount)", 90.0, 9, 10.0),
    ("Quantity 10 (90% of full price)", 90.0, 10, 10.0),
    ("Quantity 49 (90% of full price)", 441.0, 49, 10.0),
    ("Quantity 50 (75% of full price)", 1875.0, 50, 10.0)
};

bool RunTest(string desc, double expected, int qty, double price) {
    double got = CalculateDiscount(price, qty);
    if (Math.Abs(got - expected) < 0.0001) {
        Console.WriteLine($"PASS: {desc}");
        return true;
    } else {
        Console.WriteLine($"FAIL: {desc} (expected: {expected}, got: {got})");
        return false;
    }
}

foreach (var test in tests) {
    RunTest(test.desc, test.expected, test.qty, test.price);
}