using System;
using System.Collections.Generic;
using System.Text.Json;

namespace BulkDiscountTest
{
    class Program
    {
        static void Main(string[] args)
        {
            var tests = new List<(double expected, double quantity, string description)>
            {
                (100, 10, "Quantity 10"),
                (120, 11, "Quantity 11 (boundary)"),
                (240, 49, "Quantity 49"),
                (375, 50, "Quantity 50")
            };

            foreach (var test in tests)
            {
                double result = CalculateDiscount(test.quantity, test.quantity);
                if (Math.Approximately(result, test.expected, 1e-6))
                {
                    Console.WriteLine($"PASS: {test.description}");
                }
                else
                {
                    Console.WriteLine($"FAIL: {test.description} (expected: {test.expected}, got: {result})");
                }
            }
        }

        static double CalculateDiscount(double price, int quantity)
        {
            if (quantity < 10) return price * quantity;
            else if (quantity < 50) return price * quantity * 0.9;
            else return price * quantity * 0.75;
        }
    }
}