using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void TestCase(string description, double price, int quantity, double expected) {
    double got = CalculateDiscount(price, quantity);
    double epsilon = 1e-9;
    if (Math.Abs(got - expected) < epsilon) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {got})");
    }
}

// Test case 1: quantity < 10 (quantity = 5)
TestCase("quantity 5 (below 10)", 10.0, 5, 10.0 * 5);

// Test case 2: boundary condition quantity = 9 (just below 10)
TestCase("quantity 9 (boundary below 10)", 10.0, 9, 10.0 * 9);

// Test case 3: boundary condition quantity = 10 (start of 10-49 range)
TestCase("quantity 10 (boundary at 10)", 10.0, 10, 10.0 * 10 * 0.9);

// Test case 4: boundary condition quantity = 49 (just below 50)
TestCase("quantity 49 (boundary below 50)", 10.0, 49, 10.0 * 49 * 0.9);

// Test case 5: boundary condition quantity = 50 (start of 50+ range)
TestCase("quantity 50 (boundary at 50)", 10.0, 50, 10.0 * 50 * 0.75);

// Test case 6: quantity > 50 (e.g., quantity = 100)
TestCase("quantity 100 (above 50)", 10.0, 100, 10.0 * 100 * 0.75);