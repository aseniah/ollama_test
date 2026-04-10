using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void TestCalculateDiscount() {
    var testCases = new[] {
        new { Price = 10.0, Quantity = 9, Expected = 90.0 },
        new { Price = 10.0, Quantity = 10, Expected = 100.0 },
        new { Price = 10.0, Quantity = 49, Expected = 441.0 },
        new { Price = 10.0, Quantity = 50, Expected = 375.0 }
    };

    foreach (var testCase in testCases) {
        var result = CalculateDiscount(testCase.Price, testCase.Quantity);
        if (Math.Abs(result - testCase.Expected) < double.Epsilon) {
            Console.WriteLine($"PASS: Price {testCase.Price}, Quantity {testCase.Quantity}");
        } else {
            Console.WriteLine($"FAIL: Price {testCase.Price}, Quantity {testCase.Quantity} (expected: {testCase.Expected}, got: {result})");
        }
    }
}

TestCalculateDiscount();