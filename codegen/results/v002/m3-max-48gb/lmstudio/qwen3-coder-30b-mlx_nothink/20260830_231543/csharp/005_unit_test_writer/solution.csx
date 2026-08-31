double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void TestDiscount(double price, int quantity, double expected, string description) {
    double result = CalculateDiscount(price, quantity);
    if (Math.Abs(result - expected) < 0.0001) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {result})");
    }
}

// Test cases
TestDiscount(10.0, 5, 50.0, "quantity 5 (less than 10)");
TestDiscount(10.0, 9, 90.0, "quantity 9 (boundary case, less than 10)");
TestDiscount(10.0, 10, 90.0, "quantity 10 (boundary case, 10-49 discount)");
TestDiscount(10.0, 49, 441.0, "quantity 49 (boundary case, 10-49 discount)");
TestDiscount(10.0, 50, 375.0, "quantity 50 (boundary case, 50+ discount)");