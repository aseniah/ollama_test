double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test case 1: quantity = 9 (just below first boundary, no discount)
double result1 = CalculateDiscount(10.0, 9);
if (Math.Abs(result1 - 90.0) < 0.001)
    Console.WriteLine("PASS: quantity 9, price 10.0 (no discount)");
else
    Console.WriteLine($"FAIL: quantity 9, price 10.0 (expected: 90.0, got: {result1})");

// Test case 2: quantity = 10 (first boundary, 10% discount)
double result2 = CalculateDiscount(10.0, 10);
if (Math.Abs(result2 - 90.0) < 0.001)
    Console.WriteLine("PASS: quantity 10, price 10.0 (10% discount)");
else
    Console.WriteLine($"FAIL: quantity 10, price 10.0 (expected: 90.0, got: {result2})");

// Test case 3: quantity = 49 (just below second boundary, 10% discount)
double result3 = CalculateDiscount(10.0, 49);
if (Math.Abs(result3 - 441.0) < 0.001)
    Console.WriteLine("PASS: quantity 49, price 10.0 (10% discount)");
else
    Console.WriteLine($"FAIL: quantity 49, price 10.0 (expected: 441.0, got: {result3})");

// Test case 4: quantity = 50 (second boundary, 25% discount)
double result4 = CalculateDiscount(10.0, 50);
if (Math.Abs(result4 - 375.0) < 0.001)
    Console.WriteLine("PASS: quantity 50, price 10.0 (25% discount)");
else
    Console.WriteLine($"FAIL: quantity 50, price 10.0 (expected: 375.0, got: {result4})");

// Test case 5: quantity = 100 (well into third tier, 25% discount)
double result5 = CalculateDiscount(20.0, 100);
if (Math.Abs(result5 - 1500.0) < 0.001)
    Console.WriteLine("PASS: quantity 100, price 20.0 (25% discount)");
else
    Console.WriteLine($"FAIL: quantity 100, price 20.0 (expected: 1500.0, got: {result5})");

// Test case 6: quantity = 1 (minimum, no discount)
double result6 = CalculateDiscount(5.0, 1);
if (Math.Abs(result6 - 5.0) < 0.001)
    Console.WriteLine("PASS: quantity 1, price 5.0 (no discount)");
else
    Console.WriteLine($"FAIL: quantity 1, price 5.0 (expected: 5.0, got: {result6})");