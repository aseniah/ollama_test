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
	const epsilon = 1e-9

	type testCase struct {
		description string
		price       float64
		quantity    int
		expected    float64
	}

	tests := []testCase{
		{
			description: "quantity 9 (below 10, no discount)",
			price:       100.0,
			quantity:    9,
			expected:    100.0 * 9, // 900.0
		},
		{
			description: "quantity 10 (boundary: exactly 10, 10% discount)",
			price:       100.0,
			quantity:    10,
			expected:    100.0 * 10 * 0.9, // 900.0
		},
		{
			description: "quantity 49 (below 50, 10% discount)",
			price:       100.0,
			quantity:    49,
			expected:    100.0 * 49 * 0.9, // 441.0
		},
		{
			description: "quantity 50 (boundary: exactly 50, 25% discount)",
			price:       100.0,
			quantity:    50,
			expected:    100.0 * 50 * 0.75, // 3750.0
		},
	}

	for _, tc := range tests {
		got := calculateDiscount(tc.price, tc.quantity)
		if math.Abs(got-tc.expected) < epsilon {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", tc.description, tc.expected, got)
		}
	}
}