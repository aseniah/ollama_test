using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void RunTest(double price, int quantity, double expected, string description) {
    double result = CalculateDiscount(price, quantity);
    if (Math.Abs(result - expected) < 1e-9) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {result})");
    }
}

RunTest(10, 9, 90.0, "Quantity 9 (no discount)");
RunTest(10, 10, 90.0, "Quantity 10 (10% discount)");
RunTest(10, 49, 441.0, "Quantity 49 (10% discount)");
RunTest(10, 50, 375.0, "Quantity 50 (25% discount)");