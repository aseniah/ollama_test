```c#
using System;
using System.Collections.Generic;
using System.Text.Json;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

var testCases = new List<(double price, int quantity, string description)> {
    (10.0, 5, "Test 1: price=10.0, quantity=5 (quantity < 10)"),
    (20.0, 9, "Test 2: price=20.0, quantity=9 (boundary: quantity = 9)"),
    (30.0, 10, "Test 3: price=30.0, quantity=10 (boundary: quantity = 10)"),
    (40.0, 49, "Test 4: price=40.0, quantity=49 (boundary: quantity = 49)"),
    (50.0, 50, "Test 5: price=50.0, quantity=50 (boundary: quantity = 50)"),
};

var results = new List<(string status, string details)>();
var passed = 0;
var failed = 0;

foreach (var test in testCases) {
    var actual = CalculateDiscount(test.price, test.quantity);
    var expected = (test.quantity < 10 ? test.price * test.quantity : test.quantity < 50 ? test.price * test.quantity * 0.9 : test.price * test.quantity * 0.75);
    
    bool isPass = Math.Abs(actual - expected) < 0.0001;
    var status = isPass ? "PASS" : "FAIL";
    var details = $"{test.description} (expected: {expected}, got: {actual})";
    
    results.Add((status, details));
    
    Console.WriteLine($"{status}: {test.description}");
    
    if (isPass) passed++;
    else failed++;
}

Console.WriteLine();
Console.WriteLine($"Results Summary: {passed} passed, {failed} failed");
```