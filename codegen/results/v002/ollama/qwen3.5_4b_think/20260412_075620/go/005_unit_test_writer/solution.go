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
    testCases := []struct {
		quantity   int
		price      float64
		expected   float64
		description string
	}{
		{5, 10.0, 50.0, "5 units, price $10 (less than 10)"},
		{9, 10.0, 90.0, "9 units, price $10 (boundary under 10)"},
		{10, 10.0, 90.0, "10 units, price $10 (boundary at 10)"},
		{49, 10.0, 441.0, "49 units, price $10 (boundary just under 50)"},
		{50, 10.0, 375.0, "50 units, price $10 (boundary at 50)"},
		{25, 10.0, 225.0, "25 units, price $10 (in 10-49 range)"},
		{100, 10.0, 750.0, "100 units, price $10 (above 50)"},
		{1, 50.0, 50.0, "1 unit, price $50 (quantity = 1)"},
		{7, 20.0, 140.0, "7 units, price $20"},
	}

    passed := 0
    failed := 0

    for _, tc := range testCases {
        result := calculateDiscount(tc.price, tc.quantity)
        expected := tc.expected

        // Use small epsilon for float comparison
        diff := math.Abs(result - expected)
        if diff > 0.0001 {
            failed++
            fmt.Printf("FAIL: %s (expected: %.0f, got: %.0f)\n", tc.description, expected, result)
        } else {
            passed++
            fmt.Printf("PASS: %s\n", tc.description)
        }
    }

    fmt.Printf("\nTest Summary: %d passed, %d failed\n", passed, failed)
}