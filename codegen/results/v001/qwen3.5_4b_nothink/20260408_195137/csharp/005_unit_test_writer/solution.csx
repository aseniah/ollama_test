using System;

// Copy the function into the solution
static double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases
int passCount = 0;
int failCount = 0;

// Test case 1: quantity < 10 (boundary)
if (CalculateDiscount(10.0, 9) == 90.0) {
    Console.WriteLine("PASS: Price 10.0, Quantity 9");
    passCount++;
} else {
    Console.WriteLine($"FAIL: Price 10.0, Quantity 9 (expected: 90.0, got: {CalculateDiscount(10.0, 9)})");
    failCount++;
}

// Test case 2: quantity = 10 (boundary)
if (CalculateDiscount(10.0, 10) == 90.0) {
    Console.WriteLine("PASS: Price 10.0, Quantity 10");
    passCount++;
} else {
    Console.WriteLine($"FAIL: Price 10.0, Quantity 10 (expected: 90.0, got: {CalculateDiscount(10.0, 10)})");
    failCount++;
}

// Test case 3: quantity = 49 (boundary)
if (CalculateDiscount(10.0, 49) == 441.0) {
    Console.WriteLine("PASS: Price 10.0, Quantity 49");
    passCount++;
} else {
    Console.WriteLine($"FAIL: Price 10.0, Quantity 49 (expected: 441.0, got: {CalculateDiscount(10.0, 49)})");
    failCount++;
}

// Test case 4: quantity = 50 (boundary)
if (CalculateDiscount(10.0, 50) == 375.0) {
    Console.WriteLine("PASS: Price 10.0, Quantity 50");
    passCount++;
} else {
    Console.WriteLine($"FAIL: Price 10.0, Quantity 50 (expected: 375.0, got: {CalculateDiscount(10.0, 50)})");
    failCount++;
}

// Summary
Console.WriteLine($"\nTotal tests: {passCount + failCount}, Passed: {passCount}, Failed: {failCount}");