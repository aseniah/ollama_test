// Function to be tested
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Helper method to run and report test cases
void RunTest(string description, double price, int quantity, double expected)
{
    double actual = CalculateDiscount(price, quantity);
    
    // Using a small epsilon to account for floating-point precision errors
    const double epsilon = 0.000001;
    bool isPass = Math.Abs(actual - expected) < epsilon;

    if (isPass)
    {
        Console.WriteLine($"PASS: {description}");
    }
    else
    {
        // Formatting expected and actual to 2 decimal places for readable output
        Console.WriteLine($"FAIL: {description} (expected: {expected:F2}, got: {actual:F2})");
    }
}

// Test Suite
Console.WriteLine("Running Discount Calculation Tests...\n");

// Boundary condition: quantity 9 (No discount)
// Expected: 100 * 9 = 900
RunTest("Boundary Test - Quantity 9 (No discount)", 100.0, 9, 900.0);

// Boundary condition: quantity 10 (10% discount threshold)
// Expected: 100 * 10 * 0.9 = 900
RunTest("Boundary Test - Quantity 10 (10% discount)", 100.0, 10, 900.0);

// Boundary condition: quantity 49 (Just below 25% discount threshold)
// Expected: 100 * 49 * 0.9 = 4410
RunTest("Boundary Test - Quantity 49 (10% discount)", 100.0, 49, 4410.0);

// Boundary condition: quantity 50 (25% discount threshold)
// Expected: 100 * 50 * 0.75 = 3750
RunTest("Boundary Test - Quantity 50 (25% discount)", 100.0, 50, 3750.0);

// Extra case: High quantity
// Expected: 10 * 100 * 0.75 = 750
RunTest("General Test - High quantity (25% discount)", 10.0, 100, 750.0);

// Extra case: Small quantity
// Expected: 5.5 * 2 = 11
RunTest("General Test - Very small quantity (No discount)", 5.5, 2, 11.0);

Console.WriteLine("\nTests Completed.");