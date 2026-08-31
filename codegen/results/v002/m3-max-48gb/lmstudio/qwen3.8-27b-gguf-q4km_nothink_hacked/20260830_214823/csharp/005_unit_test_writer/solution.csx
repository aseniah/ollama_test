using System;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test helper
bool RunTest(string description, double price, int quantity, double expected) {
    double actual = CalculateDiscount(price, quantity);
    double tolerance = 1e-9;
    bool pass = Math.Abs(actual - expected) < tolerance;
    if (pass) {
        Console.WriteLine($"PASS: {description}");
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual})");
    }
    return pass;
}

int passed = 0;
int total = 0;

// Test case 1: quantity < 10 (no discount)
total++;
if (RunTest("quantity 5, price 10.0 (no discount)", 10.0, 5, 50.0)) passed++;

// Test case 2: boundary quantity = 9 (just below 10, no discount)
total++;
if (RunTest("quantity 9, price 10.0 (boundary, no discount)", 10.0, 9, 90.0)) passed++;

// Test case 3: boundary quantity = 10 (10% discount applies)
total++;
if (RunTest("quantity 10, price 10.0 (boundary, 10% discount)", 10.0, 10, 90.0)) passed++;

// Test case 4: quantity in middle range (10% discount)
total++;
if (RunTest("quantity 25, price 10.0 (10% discount)", 10.0, 25, 225.0)) passed++;

// Test case 5: boundary quantity = 49 (just below 50, 10% discount)
total++;
if (RunTest("quantity 49, price 10.0 (boundary, 10% discount)", 10.0, 49, 441.0)) passed++;

// Test case 6: boundary quantity = 50 (25% discount applies)
total++;
if (RunTest("quantity 50, price 10.0 (boundary, 25% discount)", 10.0, 50, 375.0)) passed++;

// Test case 7: quantity > 50 (25% discount)
total++;
if (RunTest("quantity 100, price 10.0 (25% discount)", 10.0, 100, 750.0)) passed++;

Console.WriteLine($"\nResults: {passed}/{total} tests passed");