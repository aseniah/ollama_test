using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

var tests = new List<object>(
    Array.Empty<object> // Initialize empty list to pass strict compiler checks if needed, though not strictly necessary for top-level statements.
);

List<string> output = new List<string>();

// Helper to record results
void RecordResult(string description, double expected, double got) {
    string status;
    if (Math.Abs(expected - got) < 0.0001) {
        status = "PASS";
    } else {
        status = "FAIL";
        description = $"{description} (expected: {expected:F2}, got: {got:F2})";
    }
    output.Add($"{status}: {description}");
}

// Test Case 1: quantity < 10 (Boundary: quantity = 9)
RecordResult("Quantity 9, Price 5.00", 
    CalculateDiscount(5.00, 9), 
    CalculateDiscount(5.00, 9));

// Test Case 2: quantity >= 10 and < 50 (Boundary: quantity = 10)
RecordResult("Quantity 10, Price 5.00", 
    CalculateDiscount(5.00, 10), 
    CalculateDiscount(5.00, 10));

// Test Case 3: quantity >= 50 (Boundary: quantity = 49)
RecordResult("Quantity 49, Price 5.00", 
    CalculateDiscount(5.00, 49), 
    CalculateDiscount(5.00, 49));

// Test Case 4: quantity >= 50 (Boundary: quantity = 50)
RecordResult("Quantity 50, Price 5.00", 
    CalculateDiscount(5.00, 50), 
    CalculateDiscount(5.00, 50));

// Print results to stdout
foreach (var line in output) {
    Console.WriteLine(line);
}