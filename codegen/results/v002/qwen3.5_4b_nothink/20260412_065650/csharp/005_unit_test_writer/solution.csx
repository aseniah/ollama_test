using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

var tests = new List<Dict<string, object>> {
    { "T1", { "price", 10.0, "quantity", 5, "expected", 50.0, "description", "quantity 5 (no discount)" } },
    { "T2", { "price", 10.0, "quantity", 9, "expected", 90.0, "description", "quantity 9 (no discount)" } },
    { "T3", { "price", 10.0, "quantity", 10, "expected", 90.0, "description", "quantity 10 (20% discount)" } },
    { "T4", { "price", 10.0, "quantity", 49, "expected", 441.0, "description", "quantity 49 (20% discount)" } },
    { "T5", { "price", 10.0, "quantity", 50, "expected", 375.0, "description", "quantity 50 (25% discount)" } },
};

bool RunTest(string name, double price, int quantity, double expected, string description) {
    var result = CalculateDiscount(price, quantity);
    string pass = result.Equals(expected);
    var msg = pass ? $"PASS: {description}" : $"FAIL: {description} (expected: {expected}, got: {result})";
    Console.WriteLine(msg);
    return pass;
}

var results = new List<string>();
foreach (var test in tests) {
    var name = (string)test["T1"];
    var price = (double)test["price"];
    var quantity = (int)test["quantity"];
    var expected = (double)test["expected"];
    var description = (string)test["description"];
    var passed = RunTest(name, price, quantity, expected, description);
    results.Add(passed ? "PASS" : "FAIL");
}

Console.WriteLine("Summary: " + results.Count + " test(s), passed: " + results.Count(results.Count(r) => r == "PASS"));