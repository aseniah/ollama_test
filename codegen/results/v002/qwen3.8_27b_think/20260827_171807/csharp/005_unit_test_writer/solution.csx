double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

const double tolerance = 0.0001;

int passed = 0;
int failed = 0;

void RunTest(string description, double expected, double actual) {
    if (Math.Abs(expected - actual) < tolerance) {
        Console.WriteLine($"PASS: {description}");
        passed++;
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
        failed++;
    }
}

// Test 1: quantity 9 (boundary below 10, no discount)
RunTest("quantity=9, price=10 (no discount)", 10.0 * 9, CalculateDiscount(10, 9));

// Test 2: quantity 10 (boundary at 10, 10% discount)
RunTest("quantity=10, price=10 (10% discount)", 10.0 * 10 * 0.9, CalculateDiscount(10, 10));

// Test 3: quantity 49 (boundary below 50, 10% discount)
RunTest("quantity=49, price=10 (10% discount)", 10.0 * 49 * 0.9, CalculateDiscount(10, 49));

// Test 4: quantity 50 (boundary at 50, 25% discount)
RunTest("quantity=50, price=10 (25% discount)", 10.0 * 50 * 0.75, CalculateDiscount(10, 50));

// Test 5: larger quantity in 25% tier
RunTest("quantity=100, price=5.50 (25% discount)", 5.50 * 100 * 0.75, CalculateDiscount(5.50, 100));

Console.WriteLine($"\n{passed + failed} tests: {passed} passed, {failed} failed");