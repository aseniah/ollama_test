using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

const double Epsilon = 1e-6;

void RunTest(double price, int quantity, double expected, string description)
{
    double actual = CalculateDiscount(price, quantity);
    if (Math.Abs(actual - expected) < Epsilon)
    {
        Console.WriteLine($"PASS: {description}");
    }
    else
    {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
    }
}

// Test cases covering boundary conditions
RunTest(10.0, 9,  10.0 * 9,  "Price 10, quantity 9 (below 10, no discount)");
RunTest(10.0, 10, 10.0 * 10 * 0.9, "Price 10, quantity 10 (boundary, 10% discount)");
RunTest(10.0, 49, 10.0 * 49 * 0.9, "Price 10, quantity 49 (below 50, 10% discount)");
RunTest(10.0, 50, 10.0 * 50 * 0.75, "Price 10, quantity 50 (boundary, 25% discount)");