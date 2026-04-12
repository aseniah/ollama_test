double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

int testsRun = 0;
int testsPassed = 0;

void RunTest(string description, double price, int quantity, double expected) {
    testsRun++;
    double result = CalculateDiscount(price, quantity);
    // Using a small epsilon for floating point comparison
    double epsilon = 0.0001;
    
    if (Math.Abs(result - expected) < epsilon) {
        Console.WriteLine($"PASS: {description}");
        testsPassed++;
    } else {
        Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {result})");
    }
}

// Test Case 1: Boundary condition (Quantity < 10)
RunTest("No discount for quantity 9", 10.0, 9, 90.0);

// Test Case 2: Boundary condition (Quantity >= 10, < 50)
RunTest("10% discount for quantity 10", 10.0, 10, 90.0);

// Test Case 3: Boundary condition (Quantity >= 10, < 50, upper edge)
RunTest("10% discount for quantity 49", 10.0, 49, 441.0);

// Test Case 4: Boundary condition (Quantity >= 50)
RunTest("25% discount for quantity 50", 10.0, 50, 375.0);

// Optional: Additional test for higher quantity
RunTest("25% discount for quantity 100", 5.0, 100, 375.0);

Console.WriteLine($"\nTests completed: {testsRun}, Passed: {testsPassed}, Failed: {testsRun - testsPassed}");