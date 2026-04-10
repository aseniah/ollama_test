using System;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

static double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

var testCases = new List<(string Desc, double Price, int Qty, double Expected)> {
    ("No discount (1 item)", 100.0, 1, 100.0),
    ("Partial discount (9 items)", 100.0, 9, 900.0),
    ("Break-even point (10 items)", 100.0, 10, 900.0),
    ("Deep discount (49 items)", 100.0, 49, 3510.0),
    ("Discount starts at 50 (50 items)", 100.0, 50, 3750.0)
};

int passed = 0;

foreach (var tc in testCases) {
    var result = CalculateDiscount(tc.Price, tc.Qty);
    var tolerance = 0.0001; // float precision safety
    
    bool success = Math.Abs(result - tc.Expected) < tolerance;
    
    if (success) {
        Console.WriteLine($"PASS: {tc.Desc}");
        passed++;
    } else {
        Console.WriteLine($"FAIL: {tc.Desc} (expected: {tc.Expected}, got: {result});");
    }
}

Console.WriteLine($"\n--- Summary ---\nPassed: {passed}/{testCases.Count} tests");