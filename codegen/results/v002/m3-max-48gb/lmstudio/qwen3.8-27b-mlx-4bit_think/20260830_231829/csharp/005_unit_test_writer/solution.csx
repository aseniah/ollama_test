double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

bool TestDiscount(string description, double price, int quantity, double expected) {
    double got = CalculateDiscount(price, quantity);
    const double epsilon = 1e-9;
    if (Math.Abs(got - expected) < epsilon) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {got})");
    }
    return Math.Abs(got - expected) < epsilon;
}

int passed = 0;
int failed = 0;

void RunTest(string desc, double price, int qty, double expected) {
    if (TestDiscount(desc, price, qty, expected)) passed++;
    else failed++;
}

RunTest("quantity=9, no discount (boundary below 10)", 10.0, 9, 90.0);
RunTest("quantity=10, 10% discount (boundary at 10)", 10.0, 10, 90.0);
RunTest("quantity=49, 10% discount (boundary below 50)", 10.0, 49, 441.0);
RunTest("quantity=50, 25% discount (boundary at 50)", 10.0, 50, 375.0);

Console.WriteLine($"\n{passed + failed} tests: {passed} passed, {failed} failed");