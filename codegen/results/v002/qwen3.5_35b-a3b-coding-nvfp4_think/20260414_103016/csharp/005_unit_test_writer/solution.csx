using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void RunTest(string description, double price, int quantity, double expected) {
    double actual = CalculateDiscount(price, quantity);
    const double epsilon = 0.0001;
    if (Math.Abs(actual - expected) < epsilon) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
    }
}

// Test case 1: quantity = 9 (no discount)
RunTest("quantity = 9 (no discount)", 100.0, 9, 900.0);

// Test case 2: quantity = 10 (9% discount)
RunTest("quantity = 10 (9% discount)", 100.0, 10, 900.0);

// Test case 3: quantity = 49 (9% discount)
RunTest("quantity = 49 (9% discount)", 100.0, 49, 4410.0);

// Test case 4: quantity = 50 (25% discount)
RunTest("quantity = 50 (25% discount)", 100.0, 50, 3750.0);