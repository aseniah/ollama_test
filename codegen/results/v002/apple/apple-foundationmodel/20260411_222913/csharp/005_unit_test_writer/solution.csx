using System;
using System.Collections.Generic;
using System.Text.Json;

// Function to calculate bulk discount
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases
public class BulkDiscountTest
{
    public static void RunTests()
    {
        // Test case 1: quantity 9 (less than 10)
        if (CalculateDiscount(100, 9) == 900)
        {
            Console.WriteLine("PASS: Test 1 - Quantity 9: Expected 900, got " + CalculateDiscount(100, 9));
        }
        else
        {
            Console.WriteLine("FAIL: Test 1 - Quantity 9: Expected 900, got " + CalculateDiscount(100, 9));
        }

        // Test case 2: quantity 10 (less than 50)
        if (CalculateDiscount(100, 10) == 1000)
        {
            Console.WriteLine("PASS: Test 2 - Quantity 10: Expected 1000, got " + CalculateDiscount(100, 10));
        }
        else
        {
            Console.WriteLine("FAIL: Test 2 - Quantity 10: Expected 1000, got " + CalculateDiscount(100, 10));
        }

        // Test case 3: quantity 49 (less than 50)
        if (CalculateDiscount(100, 49) == 4900)
        {
            Console.WriteLine("PASS: Test 3 - Quantity 49: Expected 4900, got " + CalculateDiscount(100, 49));
        }
        else
        {
            Console.WriteLine("FAIL: Test 3 - Quantity 49: Expected 4900, got " + CalculateDiscount(100, 49));
        }

        // Test case 4: quantity 50 (less than 50)
        if (CalculateDiscount(100, 50) == 4900)
        {
            Console.WriteLine("PASS: Test 4 - Quantity 50: Expected 4900, got " + CalculateDiscount(100, 50));
        }
        else
        {
            Console.WriteLine("FAIL: Test 4 - Quantity 50: Expected 4900, got " + CalculateDiscount(100, 50));
        }
    }
}

// Entry point
BulkDiscountTest.RunTests();