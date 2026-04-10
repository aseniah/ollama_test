using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void TestCalculateDiscount() {
    var testCases = new List<(double price, int quantity, double expected)> {
        (10.0, 9, 90.0),   // Below boundary
        (10.0, 10, 100.0), // Boundary
        (10.0, 49, 441.0), // Below higher boundary
        (10.0, 50, 375.0)  // Higher boundary
    };

    foreach (var testCase in testCases) {
        double result = CalculateDiscount(testCase.price, testCase.quantity);
        if (result == testCase.expected) {
            Console.WriteLine($"PASS: CalculateDiscount({testCase.price}, {testCase.quantity}) returned the expected value of {result}");
        } else {
            Console.WriteLine($"FAIL: CalculateDiscount({testCase.price}, {testCase.quantity}) returned {result} instead of the expected value of {testCase.expected}");
        }
    }
}

TestCalculateDiscount();