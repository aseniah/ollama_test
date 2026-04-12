using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

static void Check(string desc, double expected, double got) {
    if (Math.Abs(expected - got) < 1e-9) {
        Console.WriteLine($"PASS: {desc}");
    } else {
        Console.WriteLine($"FAIL: {desc} (expected: {expected}, got: {got})");
    }
}

double price = 100.0;

Check("Quantity 9 (price 100)", 900.0, CalculateDiscount(price, 9));
Check("Quantity 10 (price 100)", 900.0, CalculateDiscount(price, 10));
Check("Quantity 49 (price 100)", 4410.0, CalculateDiscount(price, 49));
Check("Quantity 50 (price 100)", 3750.0, CalculateDiscount(price, 50));