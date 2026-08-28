using System;
using System.Collections.Generic;

// The function to be tested
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test case structure
public class TestCase {
    public string Description { get; set; }
    public double Price { get; set; }
    public int Quantity { get; set; }
    public double Expected { get; set; }

    public TestCase(string desc, double price, int qty, double expected) {
        Description = desc;
        Price = price;
        Quantity = qty;
        Expected = expected;
    }
}

// Define test cases
var tests = new List<TestCase> {
    new TestCase("Quantity just below 10 (no discount)", 10.0, 9, 90.0),
    new TestCase("Quantity exactly 10 (10% discount)", 10.0, 10, 90.0),
    new TestCase("Quantity just below 50 (10% discount)", 10.0, 49, 441.0),
    new TestCase("Quantity exactly 50 (25% discount)", 10.0, 50, 375.0),
    new TestCase("Quantity well above 50 (25% discount)", 10.0, 100, 750.0)
};

// Execute tests
foreach (var test in tests) {
    double actual = CalculateDiscount(test.Price, test.Quantity);
    
    // Using a small epsilon for floating point comparison
    if (Math.Abs(actual - test.Expected) < 0.0001) {
        Console.WriteLine($"PASS: {test.Description}");
    } else {
        Console.WriteLine($"FAIL: {test.Description} (expected: {test.Expected}, got: {actual})");
    }
}