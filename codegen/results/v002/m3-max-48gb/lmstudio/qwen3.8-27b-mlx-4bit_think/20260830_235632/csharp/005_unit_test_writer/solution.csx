using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

const double Epsilon = 1e-9;

bool Check(double actual, double expected, string description)
{
    if (Math.Abs(actual - expected) < Epsilon)
    {
        Console.WriteLine($"PASS: {description}");
        return true;
    }
    else
    {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
        return false;
    }
}

// Test cases with price = 10.0
double price = 10.0;

Check(CalculateDiscount(price, 9),  90.0,   "quantity=9 (just below 10, no discount)");
Check(CalculateDiscount(price, 10), 90.0,   "quantity=10 (boundary, 10% discount)");
Check(CalculateDiscount(price, 49), 441.0,  "quantity=49 (just below 50, 10% discount)");
Check(CalculateDiscount(price, 50), 375.0,  "quantity=50 (boundary, 25% discount)");
Check(CalculateDiscount(price, 100), 750.0, "quantity=100 (25% discount)");
Check(CalculateDiscount(price, 1),   10.0,  "quantity=1 (no discount)");