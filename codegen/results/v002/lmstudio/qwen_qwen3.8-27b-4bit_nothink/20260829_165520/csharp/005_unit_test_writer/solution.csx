double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Helper to run a single test
bool RunTest(double price, int quantity, double expected, string description)
{
    double actual = CalculateDiscount(price, quantity);
    const double epsilon = 1e-9;
    bool pass = Math.Abs(actual - expected) < epsilon;
    if (pass)
    {
        Console.WriteLine($"PASS: {description}");
    }
    else
    {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
    }
    return pass;
}

int failures = 0;

// Boundary: quantity = 9 (just below 10 -> no discount)
if (!RunTest(100.0, 9, 100.0 * 9, "quantity=9 is below the 10 threshold, no discount applied")) failures++;

// Boundary: quantity = 10 (enters the 10..49 range -> 10% discount)
if (!RunTest(100.0, 10, 100.0 * 10 * 0.9, "quantity=10 enters the 10-49 range, 10% discount applied")) failures++;

// Boundary: quantity = 49 (still in the 10..49 range -> 10% discount)
if (!RunTest(100.0, 49, 100.0 * 49 * 0.9, "quantity=49 is at the top of the 10-49 range, 10% discount applied")) failures++;

// Boundary: quantity = 50 (enters the >=50 range -> 25% discount)
if (!RunTest(100.0, 50, 100.0 * 50 * 0.75, "quantity=50 enters the >=50 range, 25% discount applied")) failures++;

// Additional: quantity = 1 with fractional price
if (!RunTest(19.99, 1, 19.99, "quantity=1 with fractional price, no discount")) failures++;

// Additional: large quantity with non-round price
if (!RunTest(24.50, 75, 24.50 * 75 * 0.75, "quantity=75 with non-round price, 25% discount applied")) failures++;

Console.WriteLine();
if (failures == 0)
    Console.WriteLine("All tests passed.");
else
    Console.WriteLine($"{failures} test(s) failed.");