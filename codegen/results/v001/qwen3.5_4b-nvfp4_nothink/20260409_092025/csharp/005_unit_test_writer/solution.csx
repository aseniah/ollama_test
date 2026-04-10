using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

// Copy of the provided function
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

int passCount = 0;
int failCount = 0;

// Test Case 1: quantity < 10 (boundary condition: 9)
var tc1 = new { desc = "quantity=9, price=100", expectedPrice = 900.0 };
var result1 = CalculateDiscount(tc1.price, tc1.quantity);
if (Math.Abs(result1 - tc1.expectedPrice) < 0.0001) {
    Console.WriteLine($"PASS: {tc1.desc}");
    passCount++;
} else {
    Console.WriteLine($"FAIL: {tc1.desc} (expected: {tc1.expectedPrice}, got: {result1})");
    failCount++;
}

// Test Case 2: quantity == 10 (boundary condition)
var tc2 = new { desc = "quantity=10, price=25.5", expectedPrice = 255.0 }; // 25.5 * 10 * 0.9
var result2 = CalculateDiscount(tc2.price, tc2.quantity);
if (Math.Abs(result2 - tc2.expectedPrice) < 0.0001) {
    Console.WriteLine($"PASS: {tc2.desc}");
    passCount++;
} else {
    Console.WriteLine($"FAIL: {tc2.desc} (expected: {tc2.expectedPrice}, got: {result2})");
    failCount++;
}

// Test Case 3: quantity < 50 (boundary condition: 49)
var tc3 = new { desc = "quantity=49, price=20", expectedPrice = 882.0 }; // 20 * 49 * 0.9
var result3 = CalculateDiscount(tc3.price, tc3.quantity);
if (Math.Abs(result3 - tc3.expectedPrice) < 0.0001) {
    Console.WriteLine($"PASS: {tc3.desc}");
    passCount++;
} else {
    Console.WriteLine($"FAIL: {tc3.desc} (expected: {tc3.expectedPrice}, got: {result3})");
    failCount++;
}

// Test Case 4: quantity == 50 (boundary condition, switch to 0.75)
var tc4 = new { desc = "quantity=50, price=15", expectedPrice = 600.0 }; // 15 * 50 * 0.75
var result4 = CalculateDiscount(tc4.price, tc4.quantity);
if (Math.Abs(result4 - tc4.expectedPrice) < 0.0001) {
    Console.WriteLine($"PASS: {tc4.desc}");
    passCount++;
} else {
    Console.WriteLine($"FAIL: {tc4.desc} (expected: {tc4.expectedPrice}, got: {result4})");
    failCount++;
}

// Additional Test Case 5: quantity > 50 (regular bulk discount)
var tc5 = new { desc = "quantity=100, price=30", expectedPrice = 2250.0 }; // 30 * 100 * 0.75
var result5 = CalculateDiscount(tc5.price, tc5.quantity);
if (Math.Abs(result5 - tc5.expectedPrice) < 0.0001) {
    Console.WriteLine($"PASS: {tc5.desc}");
    passCount++;
} else {
    Console.WriteLine($"FAIL: {tc5.desc} (expected: {tc5.expectedPrice}, got: {result5})");
    failCount++;
}

// Output Summary JSON
var summary = new { totalTestCases = 5, passed = passCount, failed = failCount };
Console.WriteLine(JsonNode.Parse("{" + JsonSerializer.Serialize(summary) + "}"););