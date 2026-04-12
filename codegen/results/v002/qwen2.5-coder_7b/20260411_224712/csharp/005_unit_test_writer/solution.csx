using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void TestCalculateDiscount(double price, int quantity, double expected) {
    double result = CalculateDiscount(price, quantity);
    if (result == expected) {
        Console.WriteLine($"PASS: CalculateDiscount({price}, {quantity}) = {expected}");
    } else {
        Console.WriteLine($"FAIL: CalculateDiscount({price}, {quantity}) = {result} (expected: {expected})");
    }
}

int main() {
    // Test cases
    TestCalculateDiscount(10.0, 9, 90.0);   // Boundary condition: quantity < 10
    TestCalculateDiscount(10.0, 10, 90.0);  // Boundary condition: quantity = 10
    TestCalculateDiscount(10.0, 49, 441.0); // Quantity between 10 and 50
    TestCalculateDiscount(10.0, 50, 375.0); // Boundary condition: quantity = 50
}