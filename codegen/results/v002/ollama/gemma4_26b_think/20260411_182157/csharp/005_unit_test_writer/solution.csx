using System;

// The function to be tested
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Define a structure for our test cases
var testCases = new[]
{
    new TestCase("Quantity 9 (Boundary: No discount)", 10.0, 9, 90.0),
    new TestCase("Quantity 10 (Boundary: 10% discount start)", 10.0, 10, 90.0),
    new TestCase("Quantity 49 (Boundary: 10% discount end)", 10.0, 49, 441.0),
    new TestCase("Quantity 50 (Boundary: 25% discount start)", 10.0, 50, 375.0),
    new TestCase("Quantity 100 (Large quantity)", 100.0, 100, 7500.0),
    new TestCase("Quantity 1 (Single item)", 50.0, 1, 50.0)
};

// Execute tests
foreach (var test in testCases)
{
    double actual = CalculateDiscount(test.Price, test.Quantity);
    
    // Using a small epsilon for floating-point comparison
    bool isPassed = Math.Abs(actual - test.Expected) < 0.000001;

    if (isPassed)
    {
        Console.WriteLine($"PASS: {test.Description}");
    }
    else
    {
        // Formatted to match requirement: (expected: X, got: Y)
        Console.WriteLine($"FAIL: {test.Description} (expected: {test.Expected}, got: {actual})");
    }
}

// Helper class to organize test data
record TestCase(string Description, double Price, int Quantity, double Expected);