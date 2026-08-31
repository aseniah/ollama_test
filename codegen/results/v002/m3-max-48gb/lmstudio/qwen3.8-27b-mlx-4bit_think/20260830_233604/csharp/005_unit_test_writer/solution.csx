using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

var tests = new (string desc, double price, int qty, double expected)[] {
    ("quantity 9 (below threshold, no discount)", 10.0, 9, 90.0),
    ("quantity 10 (just reaches 10%, discount)", 10.0, 10, 90.0),
    ("quantity 49 (just below 25%, discount)", 10.0, 49, 441.0),
    ("quantity 50 (reaches 25%, discount)", 10.0, 50, 375.0),
    ("quantity 1 (minimum)", 5.0, 1, 5.0),
    ("quantity 100 (well above 50)", 3.0, 100, 225.0),
};

bool allPassed = true;

foreach (var (desc, price, qty, expected) in tests) {
    double got = CalculateDiscount(price, qty);
    bool pass = Math.Abs(got - expected) < 1e-9;
    if (pass) {
        Console.WriteLine($"PASS: {desc}");
    } else {
        allPassed = false;
        Console.WriteLine($"FAIL: {desc} (expected: {expected}, got: {got})");
    }
}

Console.WriteLine();
Console.WriteLine(allPassed ? "All tests passed." : "Some tests failed.");