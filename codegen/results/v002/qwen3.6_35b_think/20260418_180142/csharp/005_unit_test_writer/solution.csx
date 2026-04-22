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

double price = 100.0;
RunTest(price, 9, price * 9, "quantity < 10 (boundary 9)");
RunTest(price, 10, price * 10 * 0.9, "quantity = 10 (boundary 10)");
RunTest(price, 49, price * 49 * 0.9, "quantity = 49 (boundary 49)");
RunTest(price, 50, price * 50 * 0.75, "quantity = 50 (boundary 50)");