double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test case 1: quantity 9 (just below 10, no discount)
double expected1 = 10.0 * 9;
double got1 = CalculateDiscount(10.0, 9);
if (Math.Abs(expected1 - got1) < 0.001)
    Console.WriteLine("PASS: quantity 9, no discount");
else
    Console.WriteLine($"FAIL: quantity 9, no discount (expected: {expected1}, got: {got1})");

// Test case 2: quantity 10 (boundary, 10% discount)
double expected2 = 10.0 * 10 * 0.9;
double got2 = CalculateDiscount(10.0, 10);
if (Math.Abs(expected2 - got2) < 0.001)
    Console.WriteLine("PASS: quantity 10, 10% discount");
else
    Console.WriteLine($"FAIL: quantity 10, 10% discount (expected: {expected2}, got: {got2})");

// Test case 3: quantity 49 (just below 50, 10% discount)
double expected3 = 10.0 * 49 * 0.9;
double got3 = CalculateDiscount(10.0, 49);
if (Math.Abs(expected3 - got3) < 0.001)
    Console.WriteLine("PASS: quantity 49, 10% discount");
else
    Console.WriteLine($"FAIL: quantity 49, 10% discount (expected: {expected3}, got: {got3})");

// Test case 4: quantity 50 (boundary, 25% discount)
double expected4 = 10.0 * 50 * 0.75;
double got4 = CalculateDiscount(10.0, 50);
if (Math.Abs(expected4 - got4) < 0.001)
    Console.WriteLine("PASS: quantity 50, 25% discount");
else
    Console.WriteLine($"FAIL: quantity 50, 25% discount (expected: {expected4}, got: {got4})");