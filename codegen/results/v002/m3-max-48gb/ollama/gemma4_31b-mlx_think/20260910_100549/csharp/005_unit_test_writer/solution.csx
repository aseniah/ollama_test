double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void TestDiscount(string description, double price, int quantity, double expected) {
    double actual = CalculateDiscount(price, quantity);
    // Using a small epsilon for double comparison to avoid precision issues
    if (Math.Abs(actual - expected) < 0.000001) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
    }
}

// Test Case 1: Boundary just below the first discount threshold
TestDiscount("Quantity 9 (No discount)", 10.0, 9, 90.0);

// Test Case 2: Boundary exactly at the first discount threshold (10%)
TestDiscount("Quantity 10 (10% discount)", 10.0, 10, 90.0);

// Test Case 3: Boundary just below the second discount threshold (10%)
TestDiscount("Quantity 49 (10% discount)", 10.0, 49, 441.0);

// Test Case 4: Boundary exactly at the second discount threshold (25%)
TestDiscount("Quantity 50 (25% discount)", 10.0, 50, 375.0);

// Additional case: Higher quantity
TestDiscount("Quantity 100 (25% discount)", 10.0, 100, 750.0);