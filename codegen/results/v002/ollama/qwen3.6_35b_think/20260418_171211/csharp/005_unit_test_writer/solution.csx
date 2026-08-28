using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void RunTest(double price, int quantity, double expected, string description) {
    double actual = CalculateDiscount(price, quantity);
    if (Math.Abs(actual - expected) < 1e-9) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
    }
}

RunTest(10.0, 9, 90.0, "Quantity 9");
RunTest(10.0, 10, 90.0, "Quantity 10");
RunTest(10.0, 49, 441.0, "Quantity 49");
RunTest(10.0, 50, 375.0, "Quantity 50");
RunTest(5.0, 25, 225.0, "Quantity 25 (mid-tier test)");