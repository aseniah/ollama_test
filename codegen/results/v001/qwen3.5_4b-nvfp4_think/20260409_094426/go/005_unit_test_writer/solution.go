package main

import "fmt"

func calculateDiscount(price float64, quantity int) float64 {
    if quantity < 10 {
        return price * float64(quantity)
    } else if quantity < 50 {
        return price * float64(quantity) * 0.9
    }
    return price * float64(quantity) * 0.75
}

func main() {
    // Test cases covering boundary conditions: 9, 10, 49, 50
    testCases := []struct {
        quantity int
        price    float64
        desc     string
    }{
        {9, 100.0, "quantity=9 (below 10 threshold)"},
        {10, 100.0, "quantity=10 (boundary >= 10, < 50)"},
        {49, 100.0, "quantity=49 (boundary < 50)"},
        {50, 100.0, "quantity=50 (>= 50 threshold)"},
    }

    fmt.Println("Testing discount calculation logic...")

    for _, tc := range testCases {
        result := calculateDiscount(tc.price, tc.quantity)

        // Manually calculate expected value based on the function's logic
        expected := 0.0
        if tc.quantity < 10 {
            expected = tc.price * float64(tc.quantity)
        } else if tc.quantity < 50 {
            expected = tc.price * float64(tc.quantity) * 0.9
        } else {
            expected = tc.price * float64(tc.quantity) * 0.75
        }

        // Compare results with a small epsilon to handle floating point precision issues
        if (result - expected) >= -1e-9 && (expected - result) >= -1e-9 {
            fmt.Printf("PASS: %s\n", tc.desc)
        } else {
            fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.desc, expected, result)
        }
    }
}