using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void Test(string description, double price, int quantity, double expected) {
    double actual = CalculateDiscount(price, quantity);
    if (Math.Abs(actual - expected) < 1e-9) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
    }
}

// Test case 1: quantity = 9 (boundary - just below first threshold)
// price = 10, quantity = 9 => 10 * 9 = 90
Test("Quantity 9 (just below 10)", 10, 9, 90);

// Test case 2: quantity = 10 (boundary - first threshold)
// price = 10, quantity = 10 => 10 * 10 * 0.9 = 90
Test("Quantity 10 (first threshold)", 10, 10, 90);

// Test case 3: quantity = 49 (boundary - just below second threshold)
// price = 10, quantity = 49 => 10 * 49 * 0.9 = 441
Test("Quantity 49 (just below 50)", 10, 49, 441);

// Test case 4: quantity = 50 (boundary - second threshold)
// price = 10, quantity = 50 => 10 * 50 * 0.75 = 375
Test("Quantity 50 (second threshold)", 10, 50, 375);

// Additional test case 5: quantity = 1 (simple case)
// price = 10, quantity = 1 => 10 * 1 = 10
Test("Quantity 1 (simple case)", 10, 1, 10);

// Additional test case 6: quantity = 100 (large quantity)
// price = 10, quantity = 100 => 10 * 100 * 0.75 = 750
Test("Quantity 100 (large quantity)", 10, 100, 750);