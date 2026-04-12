using System;
using System.Collections.Generic;

double CalculateDiscount(double price, int quantity)
{
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

class Program
{
    static void Main()
    {
        // Test cases
        TestCalculateDiscount(
            new double[] { 100, 50 },
            new int[] { 10, 50, 49 },
            "Test case 1: Price 100, Quantity 10"
        );
        TestCalculateDiscount(
            new double[] { 200, 100 },
            new int[] { 9, 10, 49 },
            "Test case 2: Price 200, Quantity 9, 10, 49"
        );
        TestCalculateDiscount(
            new double[] { 150, 75 },
            new int[] { 50, 49 },
            "Test case 3: Price 150, Quantity 50, 49"
        );
        TestCalculateDiscount(
            new double[] { 300, 100 },
            new int[] { 50, 49 },
            "Test case 4: Price 300, Quantity 50, 49"
        );
    }

    static void TestCalculateDiscount(
        double[] prices, 
        int[] quantities, 
        string description)
    {
        foreach (int quantity in quantities)
        {
            double expected = CalculateDiscount(prices[0], quantity);
            double result = prices[0] * quantity;

            if (result == expected)
            {
                Console.WriteLine($"PASS: {description}");
            }
            else
            {
                Console.WriteLine($"FAIL: {description}: Expected {expected}, got {result}");
            }
        }
    }
}