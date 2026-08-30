using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void RunTest(string description, double price, int quantity, double expected) {
    double got = CalculateDiscount(price, quantity);
    if (Math.Abs(got - expected) < 1e-9) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {got})");
    }
}

// Test case 1: quantity = 9 (boundary: just below 10, first tier)
// price = 10, quantity = 9 => 10 * 9 = 90
RunTest("quantity 9, price 10", 10.0, 9, 90.0);

// Test case 2: quantity = 10 (boundary: exactly 10, second tier)
// price = 10, quantity = 10 => 10 * 10 * 0.9 = 90
RunTest("quantity 10, price 10", 10.0, 10, 90.0);

// Test case 3: quantity = 49 (boundary: just below 50, second tier)
// price = 10, quantity = 49 => 10 * 49 * 0.9 = 441
RunTest("quantity 49, price 10", 10.0, 49, 441.0);

// Test case 4: quantity = 50 (boundary: exactly 50, third tier)
// price = 10, quantity = 50 => 10 * 50 * 0.75 = 375
RunTest("quantity 50, price 10", 10.0, 50, 375.0);