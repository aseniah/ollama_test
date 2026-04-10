// CalculateDiscount function
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases
var testCases = new List<(double price, int quantity, double expected, string description)>
{
    (10.0, 9, 90.0, "Quantity 9 (no discount)"),
    (10.0, 10, 90.0, "Quantity 10 (10% discount)"),
    (10.0, 49, 441.0, "Quantity 49 (10% discount)"),
    (10.0, 50, 375.0, "Quantity 50 (25% discount)")
};

// Run tests
foreach (var test in testCases)
{
    double result = CalculateDiscount(test.price, test.quantity);
    
    if (Math.Abs(result - test.expected) < 0.001)
    {
        Console.WriteLine($"PASS: {test.description}");
    }
    else
    {
        Console.WriteLine($"FAIL: {test.description} (expected: {test.expected}, got: {result})");
    }
}