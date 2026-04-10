using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void TestCalculateDiscount() {
    var testCases = new[] {
        new { Price = 10, Quantity = 9, Expected = 90 },
        new { Price = 10, Quantity = 10, Expected = 90 },
        new { Price = 10, Quantity = 49, Expected = 441 },
        new { Price = 10, Quantity = 50, Expected = 375 }
    };

    foreach (var testCase in testCases) {
        double result = CalculateDiscount(testCase.Price, testCase.Quantity);
        if (result == testCase.Expected) {
            Console.WriteLine($"PASS: For price {testCase.Price} and quantity {testCase.Quantity}, expected {testCase.Expected} and got {result}");
        } else {
            Console.WriteLine($"FAIL: For price {testCase.Price} and quantity {testCase.Quantity}, expected {testCase.Expected} but got {result}");
        }
    }
}

TestCalculateDiscount();