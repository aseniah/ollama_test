using System;

// Copy the function into the solution
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Helper to compare doubles with a small epsilon for floating point equality
bool AreEqual(double a, double b) {
    return Math.Abs(a - b) < 1e-9;
}

void RunTest(string description, double price, int quantity, double expected) {
    double actual = CalculateDiscount(price, quantity);
    if (AreEqual(actual, expected)) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
    }
}

// Test cases including boundary conditions
// Case 1: quantity < 10 (quantity = 5)
RunTest("Quantity 5 (no discount)", 10.0, 5, 50.0);

// Case 2: quantity = 9 (boundary, just under 10)
RunTest("Quantity 9 (no discount)", 10.0, 9, 90.0);

// Case 3: quantity = 10 (boundary, enters 10% discount)
RunTest("Quantity 10 (10% discount)", 10.0, 10, 90.0); // 10 * 10 * 0.9 = 90.0

// Case 4: quantity = 49 (boundary, just under 50)
RunTest("Quantity 49 (10% discount)", 10.0, 49, 441.0); // 49 * 10 * 0.9 = 441.0

// Case 5: quantity = 50 (boundary, enters 25% discount)
RunTest("Quantity 50 (25% discount)", 10.0, 50, 375.0); // 50 * 10 * 0.75 = 375.0

// Case 6: quantity > 50 (quantity = 100)
RunTest("Quantity 100 (25% discount)", 10.0, 100, 750.0); // 100 * 10 * 0.75 = 750.0