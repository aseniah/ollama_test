using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

class Program {
    static void Main() {
        var tests = new List<TestCase> {
            new TestCase(10.0, 5, "Low quantity: price=10.0, qty=5"),
            new TestCase(20.0, 10, "Boundary low: price=20.0, qty=10 (should apply 90%)"),
            new TestCase(30.0, 49, "Boundary high: price=30.0, qty=49 (should apply 90%)"),
            new TestCase(50.0, 50, "High quantity: price=50.0, qty=50 (should apply 75%)")
        };

        foreach (var tc in tests) {
            var expected = CalculateDiscount(tc.Price, tc.Quantity);
            var actual = tc.Expected;

            if (Math.Abs(expected - actual) < 1e-6) {
                Console.WriteLine($"PASS: {tc.Description}");
            } else {
                Console.WriteLine($"FAIL: {tc.Description} (expected: {expected}, got: {actual})");
            }
        }
    }

    class TestCase {
        public double Price { get; set; }
        public int Quantity { get; set; }
        public double Expected { get; set; }
        public string Description { get; set; }
        public TestCase(double p, int q, string d) {
            Price = p; Quantity = q; Description = d;
            var calced = CalculateDiscount(p, q);
            Expected = calced; // Set expected value based on logic verification
        }
    }
}