using System;
using System.Text.Json;
using System.Text.Json.Nodes;

// Define the discount calculation function
double CalculateDiscount(double price, int quantity)
{
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases
void TestDiscountCalculations()
{
    // Test case 1: quantity 9
    double price = 100.0;
    int quantity = 9;
    double expected = 900.0; // price * quantity
    double result = CalculateDiscount(price, quantity);

    if (result == expected)
    {
        Console.WriteLine("PASS: Test case 1 (quantity 9)");
    }
    else
    {
        Console.WriteLine($"FAIL: Test case 1 (quantity 9) (expected: {expected}, got: {result})");
    }

    // Test case 2: quantity 10
    price = 150.0;
    quantity = 10;
    expected = 1500.0; // price * quantity
    result = CalculateDiscount(price, quantity);

    if (result == expected)
    {
        Console.WriteLine("PASS: Test case 2 (quantity 10)");
    }
    else
    {
        Console.WriteLine($"FAIL: Test case 2 (quantity 10) (expected: {expected}, got: {result})");
    }

    // Test case 3: quantity 49
    price = 200.0;
    quantity = 49;
    expected = 1950.0; // price * quantity * 0.9
    result = CalculateDiscount(price, quantity);

    if (result == expected)
    {
        Console.WriteLine("PASS: Test case 3 (quantity 49)");
    }
    else
    {
        Console.WriteLine($"FAIL: Test case 3 (quantity 49) (expected: {expected}, got: {result})");
    }

    // Test case 4: quantity 50
    price = 250.0;
    quantity = 50;
    expected = 12500.0; // price * quantity * 0.75
    result = CalculateDiscount(price, quantity);

    if (result == expected)
    {
        Console.WriteLine("PASS: Test case 4 (quantity 50)");
    }
    else
    {
        Console.WriteLine($"FAIL: Test case 4 (quantity 50) (expected: {expected}, got: {result})");
    }
}

// Execute the test cases
TestDiscountCalculations();