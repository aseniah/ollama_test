using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void RunTest(string description, double expected, double actual)
{
    double tolerance = 1e-9;
    if (Math.Abs(expected - actual) < tolerance)
    {
        Console.WriteLine($"PASS: {description}");
    }
    else
    {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
    }
}

// Test case 1: Quantity 9 (just below 10, no discount)
// price = 100, quantity = 9 -> 100 * 9 = 900
RunTest("Quantity 9, no discount", 100 * 9, CalculateDiscount(100, 9));

// Test case 2: Quantity 10 (boundary, 10% discount)
// price = 100, quantity = 10 -> 100 * 10 * 0.9 = 900
RunTest("Quantity 10, 10% discount boundary", 100 * 10 * 0.9, CalculateDiscount(100, 10));

// Test case 3: Quantity 49 (just below 50, 10% discount)
// price = 100, quantity = 49 -> 100 * 49 * 0.9 = 4410
RunTest("Quantity 49, 10% discount", 100 * 49 * 0.9, CalculateDiscount(100, 49));

// Test case 4: Quantity 50 (boundary, 25% discount)
// price = 100, quantity = 50 -> 100 * 50 * 0.75 = 3750
RunTest("Quantity 50, 25% discount boundary", 100 * 50 * 0.75, CalculateDiscount(100, 50));