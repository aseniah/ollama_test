using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void RunTest(double price, int quantity, double expected, string description) {
    double got = CalculateDiscount(price, quantity);
    if (Math.Abs(got - expected) < 1e-9) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {got})");
    }
}

// Boundary and edge case tests
RunTest(10.0, 9, 90.0, "Quantity 9 (< 10)");
RunTest(10.0, 10, 90.0, "Quantity 10 (>= 10, < 50)");
RunTest(10.0, 49, 441.0, "Quantity 49 (>= 10, < 50)");
RunTest(10.0, 50, 375.0, "Quantity 50 (>= 50)");
RunTest(10.0, 100, 750.0, "Quantity 100 (>= 50)");
RunTest(1.0, 0, 0.0, "Quantity 0");
RunTest(1.0, 1, 1.0, "Quantity 1");