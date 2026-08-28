using System;

// Function definition
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases
// Helper to check if doubles are approximately equal
bool DoubleEquals(double a, double b, double epsilon = 1e-9) {
    return Math.Abs(a - b) < epsilon;
}

// Test 1: quantity 9 (boundary, no discount)
{
    double price = 10.0;
    int quantity = 9;
    double expected = price * quantity;
    double got = CalculateDiscount(price, quantity);
    if (DoubleEquals(expected, got)) {
        Console.WriteLine("PASS: quantity 9 (no discount)");
    } else {
        Console.WriteLine($"FAIL: quantity 9 (no discount) (expected: {expected}, got: {got})");
    }
}

// Test 2: quantity 10 (boundary, 10% discount)
{
    double price = 10.0;
    int quantity = 10;
    double expected = price * quantity * 0.9;
    double got = CalculateDiscount(price, quantity);
    if (DoubleEquals(expected, got)) {
        Console.WriteLine("PASS: quantity 10 (10% discount)");
    } else {
        Console.WriteLine($"FAIL: quantity 10 (10% discount) (expected: {expected}, got: {got})");
    }
}

// Test 3: quantity 49 (boundary, 10% discount)
{
    double price = 10.0;
    int quantity = 49;
    double expected = price * quantity * 0.9;
    double got = CalculateDiscount(price, quantity);
    if (DoubleEquals(expected, got)) {
        Console.WriteLine("PASS: quantity 49 (10% discount)");
    } else {
        Console.WriteLine($"FAIL: quantity 49 (10% discount) (expected: {expected}, got: {got})");
    }
}

// Test 4: quantity 50 (boundary, 25% discount)
{
    double price = 10.0;
    int quantity = 50;
    double expected = price * quantity * 0.75;
    double got = CalculateDiscount(price, quantity);
    if (DoubleEquals(expected, got)) {
        Console.WriteLine("PASS: quantity 50 (25% discount)");
    } else {
        Console.WriteLine($"FAIL: quantity 50 (25% discount) (expected: {expected}, got: {got})");
    }
}

// Test 5: quantity 100 (25% discount)
{
    double price = 10.0;
    int quantity = 100;
    double expected = price * quantity * 0.75;
    double got = CalculateDiscount(price, quantity);
    if (DoubleEquals(expected, got)) {
        Console.WriteLine("PASS: quantity 100 (25% discount)");
    } else {
        Console.WriteLine($"FAIL: quantity 100 (25% discount) (expected: {expected}, got: {got})");
    }
}