// Define the function to be tested
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Helper method to run a test and print the result
void RunTest(string description, double price, int quantity, double expected) {
    double result = CalculateDiscount(price, quantity);
    // Use a small epsilon for floating point comparison to avoid precision errors
    if (Math.Abs(result - expected) < 0.0001) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {result})");
    }
}

// Test Case 1: Boundary condition - quantity just below the first threshold (9)
// Expected: No discount. Price * Quantity.
RunTest("Quantity 9 (No Discount)", 100.0, 9, 900.0);

// Test Case 2: Boundary condition - quantity at the first threshold (10)
// Expected: 10% discount (price * quantity * 0.9).
RunTest("Quantity 10 (10% Discount)", 100.0, 10, 900.0);

// Test Case 3: Boundary condition - quantity just below the second threshold (49)
// Expected: 10% discount.
RunTest("Quantity 49 (10% Discount)", 100.0, 49, 4410.0);

// Test Case 4: Boundary condition - quantity at the second threshold (50)
// Expected: 25% discount (price * quantity * 0.75).
RunTest("Quantity 50 (25% Discount)", 100.0, 50, 3750.0);

// Additional Test Case: General case above the highest threshold
RunTest("Quantity 60 (25% Discount)", 50.0, 60, 2250.0);