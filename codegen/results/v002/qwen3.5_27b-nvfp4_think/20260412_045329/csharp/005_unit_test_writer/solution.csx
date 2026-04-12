double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test case helper
void RunTest(string description, double price, int quantity, double expected) {
    var result = CalculateDiscount(price, quantity);
    // Use a small tolerance for floating point comparison
    if (Math.Abs(result - expected) < 0.01) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {result})");
    }
}

// Test cases covering boundary conditions
RunTest("Quantity below discount threshold", 100, 9, 900);
RunTest("Quantity at 10% discount threshold", 100, 10, 900);
RunTest("Quantity below 25% discount threshold", 100, 49, 4410);
RunTest("Quantity at 25% discount threshold", 100, 50, 3750);

// Additional test cases for extra coverage
RunTest("Small quantity with different price", 50, 5, 250);
RunTest("Large quantity with 25% discount", 200, 100, 15000);