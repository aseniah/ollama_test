using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Define test cases: (description, price, quantity, expectedResult)
var testCases = new (string description, double price, int quantity, double expected)[]
{
    ("Quantity below 10 (9 items)", 100.0, 9, 900.0),
    ("Quantity exactly 10 (Lower boundary for 10%)", 100.0, 10, 900.0),
    ("Quantity below 50 (49 items)", 100.0, 49, 4410.0),
    ("Quantity exactly 50 (Lower boundary for 25%)", 100.0, 50, 3750.0),
    ("Quantity well above 50 (100 items)", 100.0, 100, 7500.0)
};

foreach (var test in testCases)
{
    double actual = CalculateDiscount(test.price, test.quantity);
    
    // Using a small epsilon for double comparison to avoid precision issues
    if (Math.Abs(actual - test.expected) < 0.0001)
    {
        Console.WriteLine($"PASS: {test.description}");
    }
    else
    {
        Console.WriteLine($"FAIL: {test.description} (expected: {test.expected}, got: {actual})");
    }
}