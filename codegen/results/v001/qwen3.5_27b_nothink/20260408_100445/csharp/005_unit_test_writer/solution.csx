using System;
using System.Text.Json;
using System.Text.Json.Nodes;

// Copy of the provided function
double CalculateDiscount(double price, int quantity) 
{
    if (quantity < 10) 
        return price * quantity;
    else if (quantity < 50) 
        return price * quantity * 0.9;
    else 
        return price * quantity * 0.75;
}

// Helper class to define test cases
struct TestCase
{
    public string Description;
    public double Price;
    public int Quantity;
    public double Expected;
}

// Define test cases covering boundary conditions
var testCases = new TestCase[]
{
    new TestCase { Description = "Quantity 9 (No Discount)", Price = 100.0, Quantity = 9, Expected = 900.0 },
    new TestCase { Description = "Quantity 10 (10% Discount Boundary)", Price = 100.0, Quantity = 10, Expected = 900.0 },
    new TestCase { Description = "Quantity 49 (10% Discount)", Price = 100.0, Quantity = 49, Expected = 4410.0 },
    new TestCase { Description = "Quantity 50 (25% Discount Boundary)", Price = 100.0, Quantity = 50, Expected = 3750.0 },
    new TestCase { Description = "Quantity 100 (25% Discount)", Price = 100.0, Quantity = 100, Expected = 7500.0 }
};

// Run tests
foreach (var testCase in testCases)
{
    double result = CalculateDiscount(testCase.Price, testCase.Quantity);
    
    // Use a small epsilon for floating point comparison
    bool passed = Math.Abs(result - testCase.Expected) < 0.0001;

    if (passed)
    {
        Console.WriteLine($"PASS: {testCase.Description}");
    }
    else
    {
        Console.WriteLine($"FAIL: {testCase.Description} (expected: {testCase.Expected}, got: {result})");
    }
}