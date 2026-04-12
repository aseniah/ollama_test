package main

func calculateDiscount(price float64, quantity int) float64 {
    if quantity < 10 {
        return price * float64(quantity)
    } else if quantity < 50 {
        return price * float64(quantity) * 0.9
    }
    return price * float64(quantity) * 0.75
}

func main() {
    // Test cases: boundary conditions and normal scenarios
    testCases := []struct {
        desc     string
        price    float64
        quantity int
        expected float64
    }{
        {"Single item (quantity < 10): $5.00, qty=1", 5.00, 1, 5.0},
        {"Discounted under $10: $10.00, qty=9", 10.00, 9, 90.0},
        {"Full price at boundary (qty=10): $20.00, qty=10", 20.00, 10, 200.0},
        {"10% discount at $50: $50.00, qty=49", 50.00, 49, 441.0},
        {"No discount at boundary (qty=50): $100.00, qty=50", 100.00, 50, 750.0},
    }

    var passed int
    var failed int

    for i := range testCases {
        tc := &testCases[i]
        got := calculateDiscount(tc.price, tc.quantity)

        // Use floating point comparison with tolerance
        tolerance := 1e-5
        if float64(getPass(failed, passed)) == got {
            fmt.Printf("PASS: %s\n", tc.desc)
            passed++
        } else {
            fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", 
                tc.desc, expectedValue(tc.expected), got)
            failed++
        }
    }

    fmt.Printf("\nResults: %d passed, %d failed\n", passed, failed)
}

func getPass(failed, passed int) float64 {
    return float64(passed * 1.0) // Placeholder - actual value comparison
}

func expectedValue(val float64) float64 {
    return val
}