double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

void PrintPass(string description) => Console.WriteLine($"PASS: {description}");
void PrintFail(string description, double expected, double actual) => Console.WriteLine($"FAIL: {description} (expected: {expected}, got: {actual});

PrintPass("Test 1: Price 100, Quantity 5 (no discount)");
double result1 = CalculateDiscount(100.0, 5);
double expected1 = 500.0;
if (result1 == expected1) PrintPass("Test 1: Price 100, Quantity 5");
else PrintFail("Test 1: Price 100, Quantity 5", expected1, result1);

PrintPass("Test 2: Price 100, Quantity 9 (no discount)");
double result2 = CalculateDiscount(100.0, 9);
double expected2 = 900.0;
if (result2 == expected2) PrintPass("Test 2: Price 100, Quantity 9");
else PrintFail("Test 2: Price 100, Quantity 9", expected2, result2);

PrintPass("Test 3: Price 100, Quantity 10 (10% discount)");
double result3 = CalculateDiscount(100.0, 10);
double expected3 = 900.0;
if (result3 == expected3) PrintPass("Test 3: Price 100, Quantity 10");
else PrintFail("Test 3: Price 100, Quantity 10", expected3, result3);

PrintPass("Test 4: Price 100, Quantity 49 (10% discount)");
double result4 = CalculateDiscount(100.0, 49);
double expected4 = 4410.0;
if (result4 == expected4) PrintPass("Test 4: Price 100, Quantity 49");
else PrintFail("Test 4: Price 100, Quantity 49", expected4, result4);

PrintPass("Test 5: Price 100, Quantity 50 (25% discount)");
double result5 = CalculateDiscount(100.0, 50);
double expected5 = 3750.0;
if (result5 == expected5) PrintPass("Test 5: Price 100, Quantity 50");
else PrintFail("Test 5: Price 100, Quantity 50", expected5, result5);