double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void RunTest(string description, double price, int quantity, double expected) {
    double actual = CalculateDiscount(price, quantity);
    // Use a small tolerance for floating point comparison
    if (Math.Abs(actual - expected) < 0.001) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
    }
}

Console.WriteLine("Running CalculateDiscount tests...\n");

// Boundary test: quantity 9 (no discount)
RunTest("Quantity 9 - no discount", 10.0, 9, 90.0);

// Boundary test: quantity 10 (10% discount starts)
RunTest("Quantity 10 - 10% discount", 10.0, 10, 90.0);

// Boundary test: quantity 49 (10% discount ends)
RunTest("Quantity 49 - 10% discount", 10.0, 49, 441.0);

// Boundary test: quantity 50 (25% discount starts)
RunTest("Quantity 50 - 25% discount", 10.0, 50, 375.0);

// Additional test: quantity above 50
RunTest("Quantity 60 - 25% discount", 20.0, 60, 900.0);

// Additional test: different price values
RunTest("Different price - Quantity 8", 15.5, 8, 124.0);

Console.WriteLine("\nAll tests completed.");