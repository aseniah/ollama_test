using System;

static double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

static void RunTests() {
    // Test Case 1: quantity < 10
    double expected1 = 100.0 * 9;
    double actual1 = CalculateDiscount(100.0, 9);
    if (Math.Abs(actual1 - expected1) < 0.0001) {
        Console.WriteLine("PASS: quantity=9 price=100");
    } else {
        Console.WriteLine($"FAIL: quantity=9 price=100 (expected: {expected1}, got: {actual1})");
    }

    // Test Case 2: 10 <= quantity < 50
    double expected2 = 100.0 * 10 * 0.9;
    double actual2 = CalculateDiscount(100.0, 10);
    if (Math.Abs(actual2 - expected2) < 0.0001) {
        Console.WriteLine("PASS: quantity=10 price=100");
    } else {
        Console.WriteLine($"FAIL: quantity=10 price=100 (expected: {expected2}, got: {actual2})");
    }

    // Test Case 3: quantity < 50 (boundary)
    double expected3 = 100.0 * 49 * 0.9;
    double actual3 = CalculateDiscount(100.0, 49);
    if (Math.Abs(actual3 - expected3) < 0.0001) {
        Console.WriteLine("PASS: quantity=49 price=100");
    } else {
        Console.WriteLine($"FAIL: quantity=49 price=100 (expected: {expected3}, got: {actual3})");
    }

    // Test Case 4: quantity >= 50 (boundary)
    double expected4 = 100.0 * 50 * 0.75;
    double actual4 = CalculateDiscount(100.0, 50);
    if (Math.Abs(actual4 - expected4) < 0.0001) {
        Console.WriteLine("PASS: quantity=50 price=100");
    } else {
        Console.WriteLine($"FAIL: quantity=50 price=100 (expected: {expected4}, got: {actual4})");
    }
}

RunTests();