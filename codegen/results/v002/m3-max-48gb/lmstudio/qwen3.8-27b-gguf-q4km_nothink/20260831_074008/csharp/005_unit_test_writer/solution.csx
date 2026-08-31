using System;

// Define the CalculateDiscount function
static double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases
Test(100.0, 9, 900.0, "quantity 9 (below 10 threshold)");
Test(100.0, 10, 900.0, "quantity 10 (boundary: 10-49 threshold)");
Test(100.0, 49, 4410.0, "quantity 49 (boundary: just below 50)");
Test(100.0, 50, 3750.0, "quantity 50 (boundary: 50+ threshold)");
Test(25.5, 15, 344.25, "price 25.5, quantity 15 (10-49 range)");

void Test(double price, int quantity, double expected, string description)
{
    double got = CalculateDiscount(price, quantity);
    // Use a small tolerance for floating point comparison
    const double tolerance = 0.0001;
    if (Math.Abs(got - expected) < tolerance)
    {
        Console.WriteLine($"PASS: {description}");
    }
    else
    {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {got})");
    }
}