// Define the function to be tested
double CalculateDiscount(double price, int quantity) {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases definition
var testCases = new List<(string description, double expectedPrice, int expectedQuantity, double expectedResult)> {
    ("Single item (no discount)", 10.0, 1, 10.0),       // quantity < 10
    ("Bulk of 9 items (no discount)", 10.0, 9, 90.0),   // quantity < 10
    ("Bulk of 10 items (10% off)", 10.0, 10, 90.0),     // 10 <= quantity < 50
    ("Bulk of 49 items (10% off)", 10.0, 49, 441.0),    // 10 <= quantity < 50
    ("Large bulk (25% off)", 10.0, 50, 375.0)           // quantity >= 50
};

var results = new List<(string status, string description, double expected, double got)>();

foreach (var tc in testCases) {
    double result = CalculateDiscount(tc.expectedPrice, tc.expectedQuantity);
    
    if (result == tc.expected) {
        results.Add(( "PASS", $"{tc.description}", tc.expected, result ));
    } else {
        results.Add(( "FAIL", $"{tc.description}", tc.expected, result ));
    }
}

// Output results
foreach (var r in results) {
    string msg = $"{r.status}: {r.description}";
    if (r.status == "FAIL") {
        msg += $" (expected: {r.expected:F2}, got: {r.got:F2})";
    }
    Console.WriteLine(msg);
}

Console.WriteLine("\nTotal tests passed: " + results.Count(c => c.status == "PASS").ToString());