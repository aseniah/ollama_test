double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void TestDiscountFunction() {
    var testCases = new System.Collections.Generic.List<(string description, double expectedPrice, int quantity, double expectedResult)> {
        ("Quantity < 10 (boundary)", 10.0, 9, 90.0),
        ("Quantity >= 10, < 50 (lower boundary)", 10.0, 10, 90.0),
        ("Quantity >= 49, < 50", 10.0, 49, 44.10),
        ("Quantity >= 50 (boundary)", 10.0, 50, 37.50),
        ("Quantity > 50", 20.0, 60, 90.0)
    };

    foreach (var testCase in testCases) {
        var price = testCase.price;
        var quantity = testCase.quantity;
        var expectedResult = testCase.expectedResult;
        var description = testCase.description;

        var actualResult = CalculateDiscount(price, quantity);
        var difference = Math.Abs(actualResult - expectedResult);
        
        if (difference <= 0.001) {
            Console.WriteLine($"PASS: {description}");
        } else {
            Console.WriteLine($"FAIL: {description} (expected: {expectedResult}, got: {actualResult})");
        }
    }
}

TestDiscountFunction();