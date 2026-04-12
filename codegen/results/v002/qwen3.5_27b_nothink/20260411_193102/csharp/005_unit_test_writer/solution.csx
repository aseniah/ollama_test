double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

var testCases = new List<(int quantity, double price, double expected, string description)>
{
    (9, 10.0, 90.0, "Boundary: quantity 9 (no discount)"),
    (10, 10.0, 90.0, "Boundary: quantity 10 (10% discount)"),
    (49, 10.0, 441.0, "Boundary: quantity 49 (10% discount)"),
    (50, 10.0, 375.0, "Boundary: quantity 50 (25% discount)"),
    (100, 20.0, 1500.0, "High quantity (25% discount)")
};

foreach (var test in testCases) {
    double result = CalculateDiscount(test.price, test.quantity);
    double tolerance = 0.0001; // To handle potential floating point inaccuracies

    if (Math.Abs(result - test.expected) <= tolerance) {
        Console.WriteLine($"PASS: {test.description}");
    }
    else {
        Console.WriteLine($"FAIL: {test.description} (expected: {test.expected}, got: {result})");
    }
}