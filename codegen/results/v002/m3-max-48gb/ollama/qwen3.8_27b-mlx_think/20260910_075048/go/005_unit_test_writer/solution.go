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
		{"boundary: quantity 9 (no discount)", 10.0, 9, 10.0 * 9},
		{"boundary: quantity 10 (10% discount)", 10.0, 10, 10.0 * 10 * 0.9},
		{"boundary: quantity 49 (10% discount)", 10.0, 49, 10.0 * 49 * 0.9},
		{"boundary: quantity 50 (25% discount)", 10.0, 50, 10.0 * 50 * 0.75},
	}

	for _, tc := range tests {
		got := calculateDiscount(tc.price, tc.quantity)
		if math.Abs(got-tc.expected) < 1e-9 {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.description, tc.expected, got)
		}
	}
}