// Function provided in the requirements
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Define a structure for our test cases
// Using a tuple for simplicity in a script environment
var testCases = new (string Description, double Price, int Quantity, double Expected)[]
{
    ("Quantity 9 (No discount boundary)", 100.0, 9, 900.0),
    ("Quantity 10 (10% discount boundary)", 100.0, 10, 900.0),
    ("Quantity 49 (10% discount upper bound)", 100.0, 49, 4410.0),
    ("Quantity 50 (25% discount boundary)", 100.0, 50, 3750.0),
    ("Quantity 100 (Large bulk discount)", 10.0, 100, 750.0)
};

// Iterate through tests and validate results
foreach (var test in testCases)
{
    double actual = CalculateDiscount(test.Price, test.Quantity);
    
    // Using a small epsilon for floating point comparison
    bool passed = Math.Abs(actual - test.Expected) < 0.000001;

    if (passed)
    {
        Console.WriteLine($"PASS: {test.Description}");
    }
    else
    {
        // Format expected/got to avoid long floating point precision artifacts in output
        Console.WriteLine($"FAIL: {test.Description} (expected: {test.Expected:F2}, got: {actual:F2})");
    }
}