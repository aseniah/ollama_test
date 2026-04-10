using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

public class BulkDiscountTest
{
    public static void Main()
    {
        TestDiscount(100, 9, 90.0);
        TestDiscount(100, 10, 100.0);
        TestDiscount(100, 49, 475.0);
        TestDiscount(100, 50, 500.0);
    }

    public static void TestDiscount(double basePrice, int quantity, double expectedDiscount)
    {
        double calculatedDiscount = CalculateDiscount(basePrice, quantity);
        if (calculatedDiscount == expectedDiscount)
        {
            Console.WriteLine($"PASS: Discount calculated correctly for base price {basePrice} and quantity {quantity}");
        }
        else
        {
            Console.WriteLine($"FAIL: Expected {expectedDiscount}, got {calculatedDiscount}");
        }
    }

    public static double CalculateDiscount(double price, int quantity)
    {
        if (quantity < 10) return price * quantity;
        else if (quantity < 50) return price * quantity * 0.9;
        else return price * quantity * 0.75;
    }
}