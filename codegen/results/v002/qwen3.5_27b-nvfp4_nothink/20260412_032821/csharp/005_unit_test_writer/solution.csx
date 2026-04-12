// CalculateDiscount function definition
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test runner
var tests = new List<(string description, double price, int quantity, double expected)>
{
    ("Boundary: Quantity 9 (No Discount)", 100.0, 9, 900.0),
    ("Boundary: Quantity 10 (10% Discount)", 100.0, 10, 900.0),
    ("Boundary: Quantity 49 (10% Discount)", 100.0, 49, 4410.0),
    ("Boundary: Quantity 50 (25% Discount)", 100.0, 50, 3750.0)
};

foreach (var test in tests) {
    var result = CalculateDiscount(test.price, test.quantity);
    
    // Use a small epsilon for floating point comparison safety, 
    // though exact match is expected here given the math.
    bool passed = Math.Abs(result - test.expected) < 0.0001;

    if (passed) {
        Console.WriteLine($"PASS: {test.description}");
    } else {
        Console.WriteLine($"FAIL: {test.description} (expected: {test.expected}, got: {result})");
    }
}