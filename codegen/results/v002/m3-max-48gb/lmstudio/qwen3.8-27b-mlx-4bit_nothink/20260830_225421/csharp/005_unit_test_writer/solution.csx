```c#
using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

double CalculateExpected(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

bool Test(double price, int quantity, string description) {
    double expected = CalculateExpected(price, quantity);
    double got = CalculateDiscount(price, quantity);
    bool pass = Math.Abs(expected - got) < 1e-9;
    if (pass) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {got})");
    }
    return pass;
}

Test(10.0, 9, "Quantity 9 (below 10): 10 * 9 = 90");
Test(10.0, 10, "Quantity 10 (boundary 10-49): 10 * 10 * 0.9 = 90");
Test(10.0, 49, "Quantity 49 (boundary below 50): 10 * 49 * 0.9 = 441");
Test(10.0, 50, "Quantity 50 (boundary >= 50): 10 * 50 * 0.75 = 375");
Test(20.0, 1, "Quantity 1: 20 * 1 = 20");
Test(20.0, 25, "Quantity 25: 20 * 25 * 0.9 = 450");
Test(20.0, 100, "Quantity 100: 20 * 100 * 0.75 = 1500");
```