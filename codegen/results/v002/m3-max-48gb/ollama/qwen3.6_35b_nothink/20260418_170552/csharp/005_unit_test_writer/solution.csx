double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void CheckDiscount(string description, double price, int quantity, double expected) {
    double result = CalculateDiscount(price, quantity);
    // Use a small epsilon for floating point comparison
    double epsilon = 0.0001;
    if (Math.Abs(result - expected) < epsilon) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {result})");
    }
}

// Test case 1: quantity = 9 (below 10, no discount)
CheckDiscount("Quantity 9 - no discount", 100.0, 9, 900.0);

// Test case 2: quantity = 10 (boundary, 10% discount)
CheckDiscount("Quantity 10 - 10% discount", 100.0, 10, 900.0);

// Test case 3: quantity = 49 (boundary, 10% discount)
CheckDiscount("Quantity 49 - 10% discount", 100.0, 49, 4410.0);

// Test case 4: quantity = 50 (boundary, 25% discount)
CheckDiscount("Quantity 50 - 25% discount", 100.0, 50, 3750.0);

// Test case 5: quantity = 100 (above 50, 25% discount)
CheckDiscount("Quantity 100 - 25% discount", 100.0, 100, 7500.0);

// Test case 6: quantity = 1 (edge case, minimum quantity)
CheckDiscount("Quantity 1 - no discount", 100.0, 1, 100.0);