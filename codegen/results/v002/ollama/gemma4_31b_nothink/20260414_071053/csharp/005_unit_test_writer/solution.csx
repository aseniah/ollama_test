using System;
using System.Collections.Generic;

// The function to be tested
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test case definition
public class TestCase {
    public string Description { get; set; }
    public double Price { get; set; }
    public int Quantity { get; set; }
    public double Expected { get; set; }

    public TestCase(string desc, double p, int q, double e) {
        Description = desc;
        Price = p;
        Quantity = q;
        Expected = e;
    }
}

// Define test cases including boundary conditions
var tests = new List<TestCase> {
    new TestCase("Quantity 9 (No discount)", 10.0, 9, 90.0),
    new TestCase("Quantity 10 (10% discount boundary)", 10.0, 10, 90.0),
    new TestCase("Quantity 49 (10% discount boundary)", 10.0, 49, 441.0),
    new TestCase("Quantity 50 (25% discount boundary)", 10.0, 50, 375.0),
    new TestCase("Quantity 100 (Bulk discount)", 10.0, 100, 750.0)
};

foreach (var test in tests) {
    double actual = CalculateDiscount(test.Price, test.Quantity);
    
    // Using a small epsilon for double comparison to avoid floating point precision issues
    if (Math.Abs(actual - test.Expected) < 0.0001) {
        Console.WriteLine($"PASS: {test.Description}");
    } else {
        Console.WriteLine($"FAIL: {test.Description} (expected: {test.Expected}, got: {actual})");
    }
}