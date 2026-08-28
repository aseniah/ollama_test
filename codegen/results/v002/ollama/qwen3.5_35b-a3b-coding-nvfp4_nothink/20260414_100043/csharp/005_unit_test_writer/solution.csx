using System;

// Copy the function into the solution
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Helper to compare doubles with a small epsilon for floating point tolerance
bool AreEqual(double a, double b, double epsilon = 0.0001) {
    return Math.Abs(a - b) < epsilon;
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
RunTest("Quantity < 10 (quantity=9)", 100.0, 9, 900.0);
RunTest("Quantity >= 10 and < 50 (quantity=10)", 100.0, 10, 900.0); // 100 * 10 * 0.9
RunTest("Quantity >= 10 and < 50 (quantity=49)", 100.0, 49, 4410.0); // 100 * 49 * 0.9
RunTest("Quantity >= 50 (quantity=50)", 100.0, 50, 3750.0); // 100 * 50 * 0.75