package main

import (
	"fmt"
	"math"
)

func calculateDiscount(price float64, quantity int) float64 {
    if quantity < 10 {
        return price * float64(quantity)
    } else if quantity < 50 {
        return price * float64(quantity) * 0.9
    }
    return price * float64(quantity) * 0.75
}

func main() {
    price := 100.0

    // Define test cases covering boundary conditions
    testCases := []struct {
        desc   string
        qty    int
        expected float64
    }{
        {"Quantity 9 (full price)", 9, 900.0},
        {"Quantity 10 (90% discount)", 10, 900.0},
        {"Quantity 49 (90% discount)", 49, 4410.0},
        {"Quantity 50 (75% discount)", 50, 3750.0},
    }

    for _, tc := range testCases {
        result := calculateDiscount(price, tc.qty)
        diff := math.Abs(result - tc.expected)
        
        if diff < 1e-6 {
            fmt.Printf("PASS: %s\n", tc.desc)
        } else {
            fmt.Printf("FAIL: %s (expected: %.1f, got: %.1f)\n", tc.desc, tc.expected, result)
        }
    }
}