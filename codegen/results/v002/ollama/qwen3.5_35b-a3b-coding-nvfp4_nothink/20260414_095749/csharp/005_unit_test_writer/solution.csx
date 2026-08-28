using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void RunTest(string description, double price, int quantity, double expected) {
    double result = CalculateDiscount(price, quantity);
    // Use a small epsilon for floating point comparison
    double epsilon = 0.0001;
    if (Math.Abs(result - expected) < epsilon) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {result})");
    }
}

// Test Case 1: quantity < 10 (quantity = 9)
RunTest("quantity = 9 (less than 10)", 100.0, 9, 900.0);

// Test Case 2: quantity = 10 (boundary: first tier -> second tier)
// 10 * 100 * 0.9 = 900.0
RunTest("quantity = 10 (boundary 10)", 100.0, 10, 900.0);

// Test Case 3: quantity = 49 (boundary: second tier -> third tier)
// 49 * 100 * 0.9 = 4410.0
RunTest("quantity = 49 (boundary 49)", 100.0, 49, 4410.0);

// Test Case 4: quantity = 50 (boundary: third tier)
// 50 * 100 * 0.75 = 3750.0
RunTest("quantity = 50 (boundary 50)", 100.0, 50, 3750.0);