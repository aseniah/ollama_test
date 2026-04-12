using System;
using System.Collections.Generic;

double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

var results = new List();

results.Add(new Dictionary { Description = "Quantity 9 (no discount)", Expected = 9.0, Actual = CalculateDiscount(10.0, 9) });
results.Add(new Dictionary { Description = "Quantity 10 (10% discount)", Expected = 9.0, Actual = CalculateDiscount(10.0, 10) });
results.Add(new Dictionary { Description = "Quantity 49 (90% discount)", Expected = 8.05, Actual = CalculateDiscount(10.0, 49) });
results.Add(new Dictionary { Description = "Quantity 50 (75% discount)", Expected = 7.5, Actual = CalculateDiscount(10.0, 50) });

var output = new List();
for each result in results:
    if result["Actual"] != result["Expected"]:
        output.Add(result["Description"]) + ": FAIL (expected: " + result["Expected"].ToString() + ", got: " + result["Actual"].ToString() + ")";
    else:
        output.Add(result["Description"]);

writeLines("Test Results:\n" + string.Join("\n", output));

var count = 0;
for each test in results:
    if test["Result"] == false:
        count++;

writeLine(count == 0 ? "All tests passed!" : "Failed " + count + " test(s) out of 4");