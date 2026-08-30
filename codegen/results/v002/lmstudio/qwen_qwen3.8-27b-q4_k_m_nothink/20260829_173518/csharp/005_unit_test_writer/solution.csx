double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test case 1: quantity < 10 (boundary: 9)
double price1 = 100.0;
int qty1 = 9;
double expected1 = price1 * qty1; // 900.0
double actual1 = CalculateDiscount(price1, qty1);
if (Math.Abs(expected1 - actual1) < 1e-9)
    Console.WriteLine("PASS: quantity 9 (no discount)");
else
    Console.WriteLine($"FAIL: quantity 9 (no discount) (expected: {expected1}, got: {actual1})");

// Test case 2: quantity = 10 (boundary: start of 10-49 range)
double price2 = 100.0;
int qty2 = 10;
double expected2 = price2 * qty2 * 0.9; // 900.0
double actual2 = CalculateDiscount(price2, qty2);
if (Math.Abs(expected2 - actual2) < 1e-9)
    Console.WriteLine("PASS: quantity 10 (10% discount)");
else
    Console.WriteLine($"FAIL: quantity 10 (10% discount) (expected: {expected2}, got: {actual2})");

// Test case 3: quantity = 49 (boundary: end of 10-49 range)
double price3 = 100.0;
int qty3 = 49;
double expected3 = price3 * qty3 * 0.9; // 4410.0
double actual3 = CalculateDiscount(price3, qty3);
if (Math.Abs(expected3 - actual3) < 1e-9)
    Console.WriteLine("PASS: quantity 49 (10% discount)");
else
    Console.WriteLine($"FAIL: quantity 49 (10% discount) (expected: {expected3}, got: {actual3})");

// Test case 4: quantity = 50 (boundary: start of 50+ range)
double price4 = 100.0;
int qty4 = 50;
double expected4 = price4 * qty4 * 0.75; // 3750.0
double actual4 = CalculateDiscount(price4, qty4);
if (Math.Abs(expected4 - actual4) < 1e-9)
    Console.WriteLine("PASS: quantity 50 (25% discount)");
else
    Console.WriteLine($"FAIL: quantity 50 (25% discount) (expected: {expected4}, got: {actual4})");