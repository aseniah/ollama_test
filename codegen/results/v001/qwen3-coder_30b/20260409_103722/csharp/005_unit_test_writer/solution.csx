double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test case 1: quantity < 10 (should return full price)
double result1 = CalculateDiscount(10.0, 5);
if (result1 == 50.0) {
    Console.WriteLine("PASS: quantity 5 (less than 10)");
} else {
    Console.WriteLine($"FAIL: quantity 5 (expected: 50.0, got: {result1})");
}

// Test case 2: quantity = 9 (boundary condition)
double result2 = CalculateDiscount(10.0, 9);
if (result2 == 90.0) {
    Console.WriteLine("PASS: quantity 9 (boundary condition)");
} else {
    Console.WriteLine($"FAIL: quantity 9 (expected: 90.0, got: {result2})");
}

// Test case 3: quantity = 10 (boundary condition)
double result3 = CalculateDiscount(10.0, 10);
if (result3 == 90.0) {
    Console.WriteLine("PASS: quantity 10 (boundary condition)");
} else {
    Console.WriteLine($"FAIL: quantity 10 (expected: 90.0, got: {result3})");
}

// Test case 4: quantity = 49 (boundary condition)
double result4 = CalculateDiscount(10.0, 49);
if (result4 == 441.0) {
    Console.WriteLine("PASS: quantity 49 (boundary condition)");
} else {
    Console.WriteLine($"FAIL: quantity 49 (expected: 441.0, got: {result4})");
}

// Test case 5: quantity = 50 (boundary condition)
double result5 = CalculateDiscount(10.0, 50);
if (result5 == 375.0) {
    Console.WriteLine("PASS: quantity 50 (boundary condition)");
} else {
    Console.WriteLine($"FAIL: quantity 50 (expected: 375.0, got: {result5})");
}