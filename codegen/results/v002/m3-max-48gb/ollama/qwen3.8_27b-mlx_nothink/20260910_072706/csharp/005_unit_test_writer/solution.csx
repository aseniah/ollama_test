using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void RunTest(string description, double expected, double actual)
{
    // Use a tolerance for floating-point comparison
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

// Test case 1: quantity 9 — just below 10, no discount
// price = 10.0, quantity = 9 => 10.0 * 9 = 90.0
RunTest("Quantity 9 (no discount)", 90.0, CalculateDiscount(10.0, 9));

// Test case 2: quantity 10 — boundary, enters the 10% discount tier
// price = 10.0, quantity = 10 => 10.0 * 10 * 0.9 = 90.0
RunTest("Quantity 10 (10% discount boundary)", 90.0, CalculateDiscount(10.0, 10));

// Test case 3: quantity 49 — just below 50, still 10% discount
// price = 10.0, quantity = 49 => 10.0 * 49 * 0.9 = 441.0
RunTest("Quantity 49 (10% discount, just below 50)", 441.0, CalculateDiscount(10.0, 49));

// Test case 4: quantity 50 — boundary, enters the 25% discount tier
// price = 10.0, quantity = 50 => 10.0 * 50 * 0.75 = 375.0
RunTest("Quantity 50 (25% discount boundary)", 375.0, CalculateDiscount(10.0, 50));