using System;

class DiscountLogic {
    public static double CalculateDiscount(double price, int quantity) {
        if (quantity < 10) return price * quantity;
        else if (quantity < 50) return price * quantity * 0.9;
        else return price * quantity * 0.75;
    }
}

double Price = 100.0;
string[] Describe = new string[] { "Quantity 9: No Discount", "Quantity 10: 10% Off", "Quantity 49: 10% Off", "Quantity 50: 25% Off" };
int[] Quantities = new int[] { 9, 10, 49, 50 };

for (int i = 0; i < Quantities.Length; i++) {
    double expectedCalculated;
    if (Quantities[i] < 10) {
        expectedCalculated = Price * Quantities[i];
    } else if (Quantities[i] < 50) {
        expectedCalculated = Price * Quantities[i] * 0.9;
    } else {
        expectedCalculated = Price * Quantities[i] * 0.75;
    }

    double got = DiscountLogic.CalculateDiscount(Price, Quantities[i]);
    
    string msg = "";
    if (Math.Abs(expectedCalculated - got) < 0.0001) {
        msg = "PASS: " + Describe[i];
    } else {
        msg = $"FAIL: {Describe[i]} (expected: {expectedCalculated:F2}, got: {got:F2})";
    }
    Console.WriteLine(msg);
}