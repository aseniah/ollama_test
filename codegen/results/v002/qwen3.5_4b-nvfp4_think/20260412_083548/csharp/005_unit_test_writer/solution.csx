using System;
using System.Collections.Generic;
using System.Text.Json.Nodes;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

var testCases = new List() {
    new TestCase("Single item", 10.0, 9, 90),
    new TestCase("Boundary low (quantity=10)", 10.0, 10, 90),
    new TestCase("Boundary mid (quantity=49)", 10.0, 49, 441),
    new TestCase("Maximum discount (quantity=50)", 10.0, 50, 375)
};

foreach (var tc in testCases) {
    var result = CalculateDiscount(tc.price, tc.quantity);
    var passed = Math.Abs(result - tc.expected) < 0.001;
    
    Console.WriteLine(
        passed 
            ? $"PASS: {tc.description} (expected: {tc.expected}, got: {result})" 
            : $"FAIL: {tc.description} (expected: {tc.expected}, got: {result})"
    );
}

public struct TestCase(string description, double price, int quantity, double expected) { }