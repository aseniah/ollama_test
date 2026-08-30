using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void RunTest(double price, int quantity, string description) {
    double expected = price * quantity * (quantity < 10 ? 1.0 : quantity < 50 ? 0.9 : 0.75);
    double got = CalculateDiscount(price, quantity);
    double tolerance = 1e-9;
    if (Math.Abs(expected - got) < tolerance) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {got})");
    }
}

// Test cases covering boundary conditions
RunTest(10.0, 9, "quantity 9 (below 10, no discount)");
RunTest(10.0, 10, "quantity 10 (at 10, 10% discount)");
RunTest(20.0, 49, "quantity 49 (below 50, 10% discount)");
RunTest(20.0, 50, "quantity 50 (at 50, 25% discount)");
RunTest(5.0, 100, "quantity 100 (above 50, 25% discount)");