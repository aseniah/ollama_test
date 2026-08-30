using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void RunTest(string description, double price, int quantity, double expected) {
    double got = CalculateDiscount(price, quantity);
    double diff = Math.Abs(got - expected);
    if (diff < 1e-9) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {got})");
    }
}

// Test case 1: quantity below the first threshold (no discount)
RunTest("quantity 9 (no discount)", 10.0, 9, 90.0);

// Test case 2: boundary at quantity 10 (10% discount applies)
RunTest("quantity 10 (10% discount boundary)", 10.0, 10, 90.0);

// Test case 3: quantity just below the second threshold (10% discount)
RunTest("quantity 49 (10% discount, just below 50)", 10.0, 49, 441.0);

// Test case 4: boundary at quantity 50 (25% discount applies)
RunTest("quantity 50 (25% discount boundary)", 10.0, 50, 375.0);

// Test case 5: quantity well above the second threshold
RunTest("quantity 100 (25% discount)", 20.0, 100, 1500.0);