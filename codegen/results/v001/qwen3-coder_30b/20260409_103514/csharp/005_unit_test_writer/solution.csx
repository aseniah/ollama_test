double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test case 1: quantity = 9 (boundary condition - less than 10)
double result1 = CalculateDiscount(100.0, 9);
double expected1 = 900.0;
if (result1 == expected1) {
    Console.WriteLine("PASS: quantity 9 (no discount)");
} else {
    Console.WriteLine($"FAIL: quantity 9 (no discount) (expected: {expected1}, got: {result1})");
}

// Test case 2: quantity = 10 (boundary condition - first discount tier)
double result2 = CalculateDiscount(100.0, 10);
double expected2 = 900.0;
if (result2 == expected2) {
    Console.WriteLine("PASS: quantity 10 (10% discount)");
} else {
    Console.WriteLine($"FAIL: quantity 10 (10% discount) (expected: {expected2}, got: {result2})");
}

// Test case 3: quantity = 49 (boundary condition - second discount tier)
double result3 = CalculateDiscount(100.0, 49);
double expected3 = 4410.0;
if (result3 == expected3) {
    Console.WriteLine("PASS: quantity 49 (10% discount)");
} else {
    Console.WriteLine($"FAIL: quantity 49 (10% discount) (expected: {expected3}, got: {result3})");
}

// Test case 4: quantity = 50 (boundary condition - second discount tier boundary)
double result4 = CalculateDiscount(100.0, 50);
double expected4 = 3750.0;
if (result4 == expected4) {
    Console.WriteLine("PASS: quantity 50 (25% discount)");
} else {
    Console.WriteLine($"FAIL: quantity 50 (25% discount) (expected: {expected4}, got: {result4})");
}