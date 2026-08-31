double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

double price = 10.0;
double tolerance = 1e-9;

var tests = new (string desc, double price, int qty, double expected)[] {
    ("quantity=9 (below first threshold)", 10.0, 9, 10.0 * 9),
    ("quantity=10 (boundary: enters 10% discount)", 10.0, 10, 10.0 * 10 * 0.9),
    ("quantity=49 (boundary: still 10% discount)", 10.0, 49, 10.0 * 49 * 0.9),
    ("quantity=50 (boundary: enters 25% discount)", 10.0, 50, 10.0 * 50 * 0.75),
    ("quantity=100 (above all thresholds)", 10.0, 100, 10.0 * 100 * 0.75),
};

int passCount = 0;
int failCount = 0;

foreach (var (desc, p, qty, expected) in tests) {
    double got = CalculateDiscount(p, qty);
    if (Math.Abs(got - expected) < tolerance) {
        Console.WriteLine($"PASS: {desc}");
        passCount++;
    } else {
        Console.WriteLine($"FAIL: {desc} (expected: {expected}, got: {got})");
        failCount++;
    }
}

Console.WriteLine($"\nTotal: {passCount + failCount} | Passed: {passCount} | Failed: {failCount}");