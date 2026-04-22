using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void Test(string description, double price, int quantity, double expected) {
    double actual = CalculateDiscount(price, quantity);
    if (Math.Abs(actual - expected) < 1e-9) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
    }
}

// Boundary condition: quantity just below first threshold (9)
Test("Price 10, quantity 9 (10% discount not applied)", 10.0, 9, 90.0);

// Boundary condition: quantity at first threshold (10)
Test("Price 10, quantity 10 (10% discount applied)", 10.0, 10, 90.0);

// Boundary condition: quantity just below second threshold (49)
Test("Price 10, quantity 49 (10% discount)", 10.0, 49, 441.0);

// Boundary condition: quantity at second threshold (50)
Test("Price 10, quantity 50 (25% discount)", 10.0, 50, 375.0);

// Additional test: large quantity
Test("Price 5, quantity 100 (25% discount)", 5.0, 100, 375.0);