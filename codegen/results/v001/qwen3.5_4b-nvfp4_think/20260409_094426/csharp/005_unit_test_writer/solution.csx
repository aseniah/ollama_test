using System;
using System.Collections.Generic;
using System.Text.Json.Nodes;

public static class DiscountCalculator
{
    public static double CalculateDiscount(double price, int quantity) {
        if (quantity < 10) return price * quantity;
        else if (quantity < 50) return price * quantity * 0.9;
        else return price * quantity * 0.75;
    }

    // Run tests
    public static void RunTests() {
        IList<string> Args = new List<string>(Environment.GetCommandLineArgs()); 
        // We ignore args as we have specific test requirements defined in instructions
        
        Console.WriteLine("--- Test Case 1: quantity=9 (Boundary), price=50.0 ---");
        var result1 = DiscountCalculator.CalculateDiscount(50.0, 9);
        if (Math.Abs(result1 - (50.0 * 9)) < 0.0001) {
            Console.WriteLine($"PASS: quantity=9 units test");
        } else {
            Console.WriteLine($"FAIL: quantity=9 units test (expected: 450, got: {result1})");
        }

        Console.WriteLine("--- Test Case 2: quantity=10 (Boundary), price=50.0 ---");
        var result2 = DiscountCalculator.CalculateDiscount(50.0, 10);
        if (Math.Abs(result2 - (50.0 * 10 * 0.9)) < 0.0001) {
            Console.WriteLine($"PASS: quantity=10 units test");
        } else {
            Console.WriteLine($"FAIL: quantity=10 units test (expected: 450, got: {result2})");
        }

        Console.WriteLine("--- Test Case 3: quantity=49 (Boundary), price=50.0 ---");
        var result3 = DiscountCalculator.CalculateDiscount(50.0, 49);
        if (Math.Abs(result3 - (50.0 * 49 * 0.9)) < 0.0001) {
            Console.WriteLine($"PASS: quantity=49 units test");
        } else {
            Console.WriteLine($"FAIL: quantity=49 units test (expected: 454.5, got: {result3})");
        }

        Console.WriteLine("--- Test Case 4: quantity=50 (Boundary), price=50.0 ---");
        var result4 = DiscountCalculator.CalculateDiscount(50.0, 50);
        if (Math.Abs(result4 - (50.0 * 50 * 0.75)) < 0.0001) {
            Console.WriteLine($"PASS: quantity=50 units test");
        } else {
            Console.WriteLine($"FAIL: quantity=50 units test (expected: 3750, got: {result4})");
        }

        Console.WriteLine("--- Test Summary ---");
    }
}

DiscountCalculator.RunTests();