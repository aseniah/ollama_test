using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void RunTest(string description, double price, int quantity, double expected) {
    double actual = CalculateDiscount(price, quantity);
    // Use a small tolerance for floating-point comparison
    if (Math.Abs(actual - expected) < 0.000001) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
    }
}

// Test case 1: quantity = 9 (below 10, no discount)
// price = 10, quantity = 9 => 10 * 9 = 90
RunTest("quantity 9 (no discount)", 10, 9, 90.0);

// Test case 2: quantity = 10 (10% discount)
// price = 10, quantity = 10 => 10 * 10 * 0.9 = 90
RunTest("quantity 10 (10% discount)", 10, 10, 90.0);

// Test case 3: quantity = 49 (10% discount)
// price = 10, quantity = 49 => 10 * 49 * 0.9 = 441
RunTest("quantity 49 (10% discount)", 10, 49, 441.0);

// Test case 4: quantity = 50 (25% discount)
// price = 10, quantity = 50 => 10 * 50 * 0.75 = 375
RunTest("quantity 50 (25% discount)", 10, 50, 375.0);

// Test case 5: quantity = 5 (below 10, no discount)
// price = 5, quantity = 5 => 5 * 5 = 25
RunTest("quantity 5 (no discount)", 5, 5, 25.0);

// Test case 6: quantity = 100 (25% discount)
// price = 20, quantity = 100 => 20 * 100 * 0.75 = 1500
RunTest("quantity 100 (25% discount)", 20, 100, 1500.0);