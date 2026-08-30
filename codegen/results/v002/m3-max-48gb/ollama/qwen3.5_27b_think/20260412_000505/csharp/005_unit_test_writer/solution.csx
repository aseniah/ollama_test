using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

var testCases = new[] {
    new { price = 10.0, quantity = 9, expected = 90.0, description = "Quantity 9 (below 10, no discount)" },
    new { price = 10.0, quantity = 10, expected = 90.0, description = "Quantity 10 (10-49, 10% discount)" },
    new { price = 10.0, quantity = 49, expected = 441.0, description = "Quantity 49 (10-49, 10% discount)" },
    new { price = 10.0, quantity = 50, expected = 375.0, description = "Quantity 50 (50+, 25% discount)" },
};

foreach (var tc in testCases) {
    double result = CalculateDiscount(tc.price, tc.quantity);
    bool passed = Math.Abs(result - tc.expected) < 0.001;
    
    if (passed) {
        Console.WriteLine($"PASS: {tc.description}");
    } else {
        Console.WriteLine($"FAIL: {tc.description} (expected: {tc.expected}, got: {result})");
    }
}