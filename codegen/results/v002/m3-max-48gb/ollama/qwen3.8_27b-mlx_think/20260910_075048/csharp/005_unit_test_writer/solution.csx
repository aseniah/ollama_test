using System;
using System.Text;

double CalculateDiscount(double price, int quantity)
{
    if (quantity < 10)
        return price * quantity;
    else if (quantity < 50)
        return price * quantity * 0.9;
    else
        return price * quantity * 0.75;
}

int passed = 0;
int failed = 0;

void RunTest(string description, double expected, double actual)
{
    // Use a small tolerance for floating-point comparison
    if (Math.Abs(expected - actual) < 0.0001)
    {
        Console.WriteLine($"PASS: {description}");
        passed++;
    }
    else
    {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
        failed++;
    }
}

// Test case 1: quantity = 9 (just below the 10-unit threshold, no discount)
RunTest("quantity 9 (no discount, below 10 threshold)",
    10.0 * 9,
    CalculateDiscount(10.0, 9));

// Test case 2: quantity = 10 (boundary: exactly 10, gets 10% discount)
RunTest("quantity 10 (10% discount, at 10-unit boundary)",
    10.0 * 10 * 0.9,
    CalculateDiscount(10.0, 10));

// Test case 3: quantity = 49 (just below 50, still 10% discount)
RunTest("quantity 49 (10% discount, just below 50 threshold)",
    10.0 * 49 * 0.9,
    CalculateDiscount(10.0, 49));

// Test case 4: quantity = 50 (boundary: exactly 50, gets 25% discount)
RunTest("quantity 50 (25% discount, at 50-unit boundary)",
    10.0 * 50 * 0.75,
    CalculateDiscount(10.0, 50));

// Test case 5: quantity = 1 (minimum, no discount)
RunTest("quantity 1 (single item, no discount)",
    5.5 * 1,
    CalculateDiscount(5.5, 1));

// Test case 6: quantity = 100 (well above 50, 25% discount)
RunTest("quantity 100 (25% discount, well above 50)",
    20.0 * 100 * 0.75,
    CalculateDiscount(20.0, 100));

Console.WriteLine($"\n--- Results: {passed} passed, {failed} failed, {passed + failed} total ---");