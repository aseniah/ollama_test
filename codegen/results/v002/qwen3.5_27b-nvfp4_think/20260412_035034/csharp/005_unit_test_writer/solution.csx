using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

Action<string, double, int, double> RunTest = (description, price, quantity, expected) => {
    double actual = CalculateDiscount(price, quantity);
    
    // Use a small epsilon for floating point comparison safety
    if (Math.Abs(actual - expected) < 0.0001) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
    }
};

// Test Case 1: Quantity below 10 (No discount)
RunTest("Quantity 9 boundary", 10.0, 9, 90.0);

// Test Case 2: Quantity exactly 10 (Discount applies)
RunTest("Quantity 10 boundary", 10.0, 10, 90.0);

// Test Case 3: Quantity below 50 (90% price)
RunTest("Quantity 49 boundary", 10.0, 49, 441.0);

// Test Case 4: Quantity exactly 50 (75% price)
RunTest("Quantity 50 boundary", 10.0, 50, 375.0);