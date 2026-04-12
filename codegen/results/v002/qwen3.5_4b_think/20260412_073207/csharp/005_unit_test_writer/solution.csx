using System;

static double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases
var tests = new[] {
    (price: 10, quantity: 9, expected: 90, desc: "quantity 9, price 10"),
    (price: 10, quantity: 10, expected: 90, desc: "quantity 10, price 10"),
    (price: 10, quantity: 49, expected: 441, desc: "quantity 49, price 10"),
    (price: 10, quantity: 50, expected: 375, desc: "quantity 50, price 10"),
};

foreach (var test in tests) {
    var got = CalculateDiscount(test.price, test.quantity);
    var expected = test.expected;
    if (got == expected) {
        Console.WriteLine($"PASS: {test.desc}");
    } else {
        Console.WriteLine($"FAIL: {test.desc} (expected: {expected}, got: {got})");
    }
}