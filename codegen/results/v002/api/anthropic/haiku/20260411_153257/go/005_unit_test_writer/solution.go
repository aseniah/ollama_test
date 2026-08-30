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
	type testCase struct {
		description string
		price       float64
		quantity    int
		expected    float64
	}

	tests := []testCase{
		{"quantity 9 (no discount)", 10.0, 9, 90.0},
		{"quantity 10 (10% discount boundary)", 10.0, 10, 90.0},
		{"quantity 49 (10% discount upper boundary)", 10.0, 49, 441.0},
		{"quantity 50 (25% discount boundary)", 10.0, 50, 375.0},
		{"quantity 1 (single item)", 5.0, 1, 5.0},
		{"quantity 100 (large order)", 20.0, 100, 1500.0},
	}

	for _, tc := range tests {
		got := calculateDiscount(tc.price, tc.quantity)
		if math.Abs(got-tc.expected) < 1e-9 {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %g, got: %g)\n", tc.description, tc.expected, got)
		}
	}
}
