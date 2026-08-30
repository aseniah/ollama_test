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

RunTest(100, 9, 900, "Quantity 9 (less than 10)");
RunTest(100, 10, 900, "Quantity 10 (boundary)");
RunTest(100, 49, 4410, "Quantity 49 (less than 50)");
RunTest(100, 50, 3750, "Quantity 50 (boundary)");