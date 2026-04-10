#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

// Function definition provided in the prompt
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test Cases: Boundary conditions (9, 10, 49, 50) and one standard case
List<(string description, double price, int quantity, double expected)> testCases = new()
{
    ("Boundary: Quantity just below 10", 10.0, 9, 90.0),       // 10 * 9 = 90
    ("Boundary: Quantity exactly 10 (triggers 10% off)", 10.0, 10, 90.0), // 10 * 10 * 0.9 = 90
    ("Boundary: Quantity just below 50", 10.0, 49, 441.0),     // 10 * 49 * 0.9 = 441
    ("Boundary: Quantity exactly 50 (triggers 25% off)", 10.0, 50, 375.0), // 10 * 50 * 0.75 = 375
    ("Standard: Large quantity over 50", 20.0, 60, 900.0)      // 20 * 60 * 0.75 = 900
};

int passCount = 0;
int failCount = 0;

Console.WriteLine("Starting Unit Tests for CalculateDiscount...\n");

foreach (var test in testCases)
{
    double result = CalculateDiscount(test.price, test.quantity);
    
    // Allow for minor floating point discrepancies using a small epsilon
    bool isPass = Math.Abs(result - test.expected) < 0.0001;

    if (isPass)
    {
        Console.WriteLine($"PASS: {test.description}");
        passCount++;
    }
    else
    {
        Console.WriteLine($"FAIL: {test.description} (expected: {test.expected}, got: {result})");
        failCount++;
    }
}

Console.WriteLine($"\nResults: {passCount} passed, {failCount} failed.");

// Optional: Return a JSON summary for automation tools if desired
var summary = new Dictionary<string, string>
{
    ["status"] = failCount > 0 ? "failure" : "success",
    ["passed"] = passCount.ToString(),
    ["failed"] = failCount.ToString()
};

Console.WriteLine("\nJSON Summary:");
Console.WriteLine(System.Text.Json.JsonSerializer.Serialize(summary));