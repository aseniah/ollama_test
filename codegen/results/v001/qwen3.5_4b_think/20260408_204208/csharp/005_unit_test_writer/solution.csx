using System;

// The function provided
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Configuration for all tests
double price = 100.0;

// Test Case 1: Quantity 9 (boundary condition, no discount)
int qty1 = 9;
double expected1 = price * qty1;
double got1 = CalculateDiscount(price, qty1);
string desc1 = "Quantity 9 (no discount)";
if (Math.Abs(expected1 - got1) < 0.0001) {
    Console.WriteLine($"PASS: {desc1}");
} else {
    Console.WriteLine($"FAIL: {desc1} (expected: {expected1}, got: {got1})");
}

// Test Case 2: Quantity 10 (boundary condition, 90% price)
int qty2 = 10;
double expected2 = price * qty2 * 0.9;
double got2 = CalculateDiscount(price, qty2);
string desc2 = "Quantity 10 (90% discount)";
if (Math.Abs(expected2 - got2) < 0.0001) {
    Console.WriteLine($"PASS: {desc2}");
} else {
    Console.WriteLine($"FAIL: {desc2} (expected: {expected2}, got: {got2})");
}

// Test Case 3: Quantity 49 (boundary condition, 90% price)
int qty3 = 49;
double expected3 = price * qty3 * 0.9;
double got3 = CalculateDiscount(price, qty3);
string desc3 = "Quantity 49 (90% discount)";
if (Math.Abs(expected3 - got3) < 0.0001) {
    Console.WriteLine($"PASS: {desc3}");
} else {
    Console.WriteLine($"FAIL: {desc3} (expected: {expected3}, got: {got3})");
}

// Test Case 4: Quantity 50 (boundary condition, 75% price)
int qty4 = 50;
double expected4 = price * qty4 * 0.75;
double got4 = CalculateDiscount(price, qty4);
string desc4 = "Quantity 50 (75% discount)";
if (Math.Abs(expected4 - got4) < 0.0001) {
    Console.WriteLine($"PASS: {desc4}");
} else {
    Console.WriteLine($"FAIL: {desc4} (expected: {expected4}, got: {got4})");
}