using System;

public class DiscountCalculator {
    public static double CalculateDiscount(double price, int quantity) {
        if (quantity < 10) return price * quantity;
        else if (quantity < 50) return price * quantity * 0.9;
        else return price * quantity * 0.75;
    }

    public static void Main(string[] args) {
        TestCalculateDiscount(2.0, 9, 18.0);
        TestCalculateDiscount(3.0, 10, 30.0);
        TestCalculateDiscount(4.0, 49, 147.0);
        TestCalculateDiscount(5.0, 50, 187.5);
    }

    public static void TestCalculateDiscount(double price, int quantity, double expected) {
        double result = CalculateDiscount(price, quantity);
        if (result == expected) {
            Console.WriteLine($"PASS: Price {price}, Quantity {quantity}");
        } else {
            Console.WriteLine($"FAIL: Price {price}, Quantity {quantity} - Expected: {expected}, Got: {result}");
        }
    }
}