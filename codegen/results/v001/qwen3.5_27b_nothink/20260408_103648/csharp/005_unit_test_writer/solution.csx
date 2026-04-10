#r "System.Text.Json"

using System;
using System.Text.Json;
using System.Text.Json.Nodes;

// The function to test
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Helper to define a test case
TestRecord[] tests = new TestRecord[] {
    new TestRecord { Description = "Quantity 9 (No discount)", Price = 100.0, Qty = 9, Expected = 900.0 },
    new TestRecord { Description = "Quantity 10 (10% discount)", Price = 100.0, Qty = 10, Expected = 900.0 },
    new TestRecord { Description = "Quantity 49 (10% discount)", Price = 100.0, Qty = 49, Expected = 4410.0 },
    new TestRecord { Description = "Quantity 50 (25% discount)", Price = 100.0, Qty = 50, Expected = 3750.0 },
    new TestRecord { Description = "Quantity 100 (25% discount)", Price = 25.0, Qty = 100, Expected = 1875.0 }
};

// Run tests
foreach (var test in tests) {
    double result = CalculateDiscount(test.Price, test.Qty);
    
    // Allow small floating point tolerance if necessary, though exact matches expected here
    bool passed = Math.Abs(result - test.Expected) < 0.0001;

    if (passed) {
        Console.WriteLine($"PASS: {test.Description}");
    } else {
        Console.WriteLine($"FAIL: {test.Description} (expected: {test.Expected}, got: {result})");
    }
}

record TestRecord {
    public string Description { get; init; } = "";
    public double Price { get; init; }
    public int Qty { get; init; }
    public double Expected { get; init; }
}