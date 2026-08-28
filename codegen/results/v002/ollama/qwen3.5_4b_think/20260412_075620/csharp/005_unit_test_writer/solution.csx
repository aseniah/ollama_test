#r "System"

// Copy the CalculateDiscount function as specified
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test 1: quantity = 9 (boundary < 10)
int price = 10.0
int qty9 = 9
double expected9 = price * qty9
double actual9 = CalculateDiscount(price, qty9)
if (Math.Abs(actual9 - expected9) < 0.001) {
    Console.WriteLine("PASS: quantity 9 test")
} else {
    Console.WriteLine("FAIL: quantity 9 test (expected: {expected9}, got: {actual9})")
}

// Test 2: quantity = 10 (boundary >= 10 and < 50)
int qty10 = 10
double expected10 = price * qty10 * 0.9
double actual10 = CalculateDiscount(price, qty10)
if (Math.Abs(actual10 - expected10) < 0.001) {
    Console.WriteLine("PASS: quantity 10 test")
} else {
    Console.WriteLine("FAIL: quantity 10 test (expected: {expected10}, got: {actual10})")
}

// Test 3: quantity = 49 (boundary < 50)
int qty49 = 49
double expected49 = price * qty49 * 0.9
double actual49 = CalculateDiscount(price, qty49)
if (Math.Abs(actual49 - expected49) < 0.001) {
    Console.WriteLine("PASS: quantity 49 test")
} else {
    Console.WriteLine("FAIL: quantity 49 test (expected: {expected49}, got: {actual49})")
}

// Test 4: quantity = 50 (boundary >= 50)
int qty50 = 50
double expected50 = price * qty50 * 0.75
double actual50 = CalculateDiscount(price, qty50)
if (Math.Abs(actual50 - expected50) < 0.001) {
    Console.WriteLine("PASS: quantity 50 test")
} else {
    Console.WriteLine("FAIL: quantity 50 test (expected: {expected50}, got: {actual50})")
}