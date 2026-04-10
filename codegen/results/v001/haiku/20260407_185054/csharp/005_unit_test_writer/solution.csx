double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

var testCases = new (double price, int quantity, double expected)[] {
    (10.0, 9, 90.0),
    (10.0, 10, 90.0),
    (10.0, 49, 441.0),
    (10.0, 50, 375.0)
};

foreach (var (price, quantity, expected) in testCases) {
    double result = CalculateDiscount(price, quantity);
    if (Math.Abs(result - expected) < 0.001) {
        WriteLine($"PASS: price={price}, quantity={quantity}, result={result}");
    } else {
        WriteLine($"FAIL: price={price}, quantity={quantity} (expected: {expected}, got: {result})");
    }
}
