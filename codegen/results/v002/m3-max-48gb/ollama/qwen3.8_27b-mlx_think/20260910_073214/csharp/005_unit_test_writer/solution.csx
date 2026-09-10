double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Helper to run a test with floating-point tolerance
void RunTest(string description, double price, int quantity, double expected)
{
    double got = CalculateDiscount(price, quantity);
    double tolerance = 0.0001;
    if (Math.Abs(got - expected) < tolerance)
    {
        Console.WriteLine($"PASS: {description}");
    }
    else
    {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {got})");
    }
}

// Test case 1: quantity below 10 (no discount)
RunTest("quantity=9, no discount", 10.0, 9, 90.0);

// Test case 2: boundary at quantity=10 (10% discount kicks in)
RunTest("quantity=10, 10% discount boundary", 10.0, 10, 100.0 * 0.9);

// Test case 3: boundary just below 50 (still 10% discount)
RunTest("quantity=49, 10% discount upper boundary", 10.0, 49, 490.0 * 0.9);

// Test case 4: boundary at quantity=50 (25% discount kicks in)
RunTest("quantity=50, 25% discount boundary", 10.0, 50, 500.0 * 0.75);

// Test case 5: quantity above 50 (25% discount)
RunTest("quantity=100, 25% discount", 5.0, 100, 500.0 * 0.75);

// Test case 6: quantity=1 (minimum, no discount)
RunTest("quantity=1, no discount", 3.5, 1, 3.5);