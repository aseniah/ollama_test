double CalculateDiscount(double price, int quantity)
{
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

int passCount = 0;
int failCount = 0;

void Check(string description, double expected, double actual)
{
    const double tolerance = 1e-9;
    if (Math.Abs(expected - actual) < tolerance)
    {
        Console.WriteLine($"PASS: {description}");
        passCount++;
    }
    else
    {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
        failCount++;
    }
}

// Test case 1: quantity below 10 (no discount) — boundary at 9
Check("price=10.0, quantity=9 (no discount)", 90.0, CalculateDiscount(10.0, 9));

// Test case 2: quantity at 10 (10% discount kicks in) — boundary at 10
Check("price=10.0, quantity=10 (10% discount)", 90.0, CalculateDiscount(10.0, 10));

// Test case 3: quantity below 50 (10% discount) — boundary at 49
Check("price=10.0, quantity=49 (10% discount)", 441.0, CalculateDiscount(10.0, 49));

// Test case 4: quantity at 50 (25% discount kicks in) — boundary at 50
Check("price=10.0, quantity=50 (25% discount)", 375.0, CalculateDiscount(10.0, 50));

// Test case 5: large quantity (25% discount)
Check("price=20.0, quantity=100 (25% discount)", 1500.0, CalculateDiscount(20.0, 100));

// Test case 6: quantity = 1 (no discount)
Check("price=5.0, quantity=1 (no discount)", 5.0, CalculateDiscount(5.0, 1));

Console.WriteLine();
Console.WriteLine($"Results: {passCount} passed, {failCount} failed out of {passCount + failCount} tests.");