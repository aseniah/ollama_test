using System;
using System.Text;

static double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

static void RunTest(string description, double price, int quantity, double expected) {
    double got = CalculateDiscount(price, quantity);
    double tol = Math.Max(1e-6, Math.Abs(expected) * 1e-6);
    if (Math.Abs(got - expected) <= tol) {
        Console.WriteLine("PASS: " + description);
    } else {
        var sb = new StringBuilder();
        sb.Append("FAIL: ").Append(description).Append(" (expected: ");
        sb.Append(expected.ToString("0.00", System.Globalization.CultureInfo.InvariantCulture)).Append(", got: ");
        sb.Append(got.ToString("0.00", System.Globalization.CultureInfo.InvariantCulture)).Append(")");
        Console.WriteLine(sb.ToString());
    }
}

// At least 4 tests including boundaries: 9, 10, 49, 50
RunTest("quantity 9 (no discount)", 10.0, 9, 90.0);
RunTest("quantity 10 (10% discount)", 10.0, 10, 90.0);
RunTest("quantity 49 (10% discount)", 10.0, 49, 441.0);
RunTest("quantity 50 (25% discount)", 10.0, 50, 375.0);
RunTest("quantity 1 (no discount)", 5.5, 1, 5.5);
RunTest("quantity 100 (25% discount)", 2.25, 100, 168.75);