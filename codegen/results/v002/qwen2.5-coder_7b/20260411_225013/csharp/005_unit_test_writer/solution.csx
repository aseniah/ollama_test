using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void TestCalculateDiscount(double price, int quantity, double expected) {
    double result = CalculateDiscount(price, quantity);
    if (Math.Abs(result - expected) < 0.01) {
        Console.WriteLine($"PASS: Quantity {quantity} with price {price} gives discount {result}");
    } else {
        Console.WriteLine($"FAIL: Quantity {quantity} with price {price} expected {expected} but got {result}");
    }
}

TestCalculateDiscount(10, 9, 90);      // Boundary condition
TestCalculateDiscount(10, 10, 100);     // Exactly at boundary
TestCalculateDiscount(10, 49, 441);     // Just below the next discount level
TestCalculateDiscount(10, 50, 375);     // Exactly at the next discount level