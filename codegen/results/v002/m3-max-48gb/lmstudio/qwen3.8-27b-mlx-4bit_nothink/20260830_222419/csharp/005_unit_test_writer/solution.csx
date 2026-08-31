using System;

public class Program
{
    static double CalculateDiscount(double price, int quantity)
    {
        if (quantity < 10) return price * quantity;
        else if (quantity < 50) return price * quantity * 0.9;
        else return price * quantity * 0.75;
    }

    static void Main()
    {
        int failures = 0;
        int total = 0;

        (double price, int quantity, double expected, string description)[] testCases =
        {
            (10.0, 9,  90.0,  "quantity=9 (below first threshold, no discount)"),
            (10.0, 10, 90.0,  "quantity=10 (boundary, 10% discount)"),
            (10.0, 49, 441.0, "quantity=49 (just below second threshold, 10% discount)"),
            (10.0, 50, 375.0, "quantity=50 (boundary, 25% discount)"),
            (10.0, 51, 382.5, "quantity=51 (above second threshold, 25% discount)"),
            (20.0, 8,  160.0, "price=20, quantity=8 (below first threshold)"),
            (20.0, 25, 450.0, "price=20, quantity=25 (10% discount)"),
            (20.0, 100, 1500.0,"price=20, quantity=100 (25% discount)"),
        };

        foreach (var (price, quantity, expected, description) in testCases)
        {
            total++;
            double actual = CalculateDiscount(price, quantity);
            // Use tolerance for floating point comparison
            if (Math.Abs(actual - expected) < 1e-9)
            {
                Console.WriteLine($"PASS: {description}");
            }
            else
            {
                failures++;
                Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
            }
        }

        Console.WriteLine($"\n{total - failures}/{total} tests passed.");
        if (failures > 0)
            Environment.Exit(1);
    }
}