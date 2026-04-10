using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

var tests = new[] {
    new { Price = 100.0, Quantity = 9, Expected = 900.0, Description = "Boundary: Quantity 9 (No discount)" },
    new { Price = 100.0, Quantity = 10, Expected = 900.0, Description = "Boundary: Quantity 10 (10% discount)" },
    new { Price = 100.0, Quantity = 49, Expected = 4410.0, Description = "Boundary: Quantity 49 (10% discount)" },
    new { Price = 100.0, Quantity = 50, Expected = 3750.0, Description = "Boundary: Quantity 50 (25% discount)" }
};

double epsilon = 0.000001;

foreach (var test in tests) {
    var actual = CalculateDiscount(test.Price, test.Quantity);
    bool passed = Math.Abs(actual - test.Expected) < epsilon;

    if (passed) {
        Console.WriteLine($"PASS: {test.Description}");
    } else {
        Console.WriteLine($"FAIL: {test.Description} (expected: {test.Expected}, got: {actual})");
    }
}