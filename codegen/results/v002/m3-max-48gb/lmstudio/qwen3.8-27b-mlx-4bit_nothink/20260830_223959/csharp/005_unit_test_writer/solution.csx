using System;

// Define the function to test
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Define test cases
// Format: (price, quantity, expectedResult, description)
var tests = new List<(double price, int quantity, double expected, string description)>
{
    (10.0, 9, 90.0, "Quantity 9 (below 10, no discount)"),
    (10.0, 10, 90.0, "Quantity 10 (boundary 10, 10% discount)"),
    (10.0, 49, 441.0, "Quantity 49 (below 50, 10% discount)"),
    (10.0, 50, 375.0, "Quantity 50 (boundary 50, 25% discount)"),
    (100.0, 100, 7500.0, "Quantity 100 (above 50, 25% discount)"),
    (0.0, 10, 0.0, "Price 0 (edge case: zero price)"),
    (5.5, 2, 11.0, "Non-integer price (basic calculation)")
};

// Run tests
bool allPassed = true;
foreach (var test in tests)
{
    double result = CalculateDiscount(test.price, test.quantity);
    
    // Use a tolerance for floating point comparisons
    if (Math.Abs(result - test.expected) < 0.0001)
    {
        Console.WriteLine($"PASS: {test.description}");
    }
    else
    {
        allPassed = false;
        Console.WriteLine($"FAIL: {test.description} (expected: {test.expected}, got: {result})");
    }
}

if (allPassed)
{
    Console.WriteLine("All tests passed.");
}