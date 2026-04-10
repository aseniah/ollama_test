using System;
using System.Text.Json;
using System.Text.Json.Nodes;

// Define the discount calculation function
double CalculateDiscount(double price, int quantity)
{
    if (quantity < 10)
    {
        return price * quantity;
    }
    else if (quantity < 50)
    {
        return price * quantity * 0.9;
    }
    else
    {
        return price * quantity * 0.75;
    }
}

// Test cases
double[] testPrices = { 100.0, 150.0, 200.0, 250.0 };

Console.WriteLine("Testing CalculateDiscount function...");

// Test case 1: quantity = 9
double expected = 9 * 100.0;
double result = CalculateDiscount(100.0, 9);
if (result == expected)
{
    Console.WriteLine("PASS: Test case 1 - Quantity 9");
}
else
{
    Console.WriteLine("FAIL: Test case 1 - Quantity 9 (expected: " + expected + ", got: " + result + ")");
}

// Test case 2: quantity = 10
double expected = 10 * 100.0;
double result = CalculateDiscount(100.0, 10);
if (result == expected)
{
    Console.WriteLine("PASS: Test case 2 - Quantity 10");
}
else
{
    Console.WriteLine("FAIL: Test case 2 - Quantity 10 (expected: " + expected + ", got: " + result + ")");
}

// Test case 3: quantity = 49
double expected = 49 * 100.0 * 0.9;
double result = CalculateDiscount(100.0, 49);
if (result == expected)
{
    Console.WriteLine("PASS: Test case 3 - Quantity 49");
}
else
{
    Console.WriteLine("FAIL: Test case 3 - Quantity 49 (expected: " + expected + ", got: " + result + ")");
}

// Test case 4: quantity = 50
double expected = 50 * 100.0 * 0.75;
double result = CalculateDiscount(100.0, 50);
if (result == expected)
{
    Console.WriteLine("PASS: Test case 4 - Quantity 50");
}
else
{
    Console.WriteLine("FAIL: Test case 4 - Quantity 50 (expected: " + expected + ", got: " + result + ")");
}