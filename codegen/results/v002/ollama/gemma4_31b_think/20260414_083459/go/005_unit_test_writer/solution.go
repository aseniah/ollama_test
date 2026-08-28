package main

import (
	"fmt"
	"math"
)

// calculateDiscount calculates the total cost based on quantity thresholds.
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

type testCase struct {
	description string
	price       float64
	quantity    int
	expected    float64
}

func main() {
	// Define test cases covering the requested boundary conditions
	tests := []testCase{
		{
			description: "Boundary: just below 10 items (no discount)",
			price:       100.0,
			quantity:    9,
			expected:    900.0,
		},
		{
			description: "Boundary: exactly 10 items (10% discount)",
			price:       100.0,
			quantity:    10,
			expected:    900.0, // 100 * 10 * 0.9
		},
		{
			description: "Boundary: just below 50 items (10% discount)",
			price:       100.0,
			quantity:    49,
			expected:    4410.0, // 100 * 49 * 0.9
		},
		{
			description: "Boundary: exactly 50 items (25% discount)",
			price:       100.0,
			quantity:    50,
			expected:    3750.0, // 100 * 50 * 0.75
		},
	}

	for _, tc := range tests {
		got := calculateDiscount(tc.price, tc.quantity)

		// Using a small epsilon for floating point comparison to avoid precision issues
		if math.Abs(got-tc.expected) < 0.0001 {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %f, got: %f)\n", tc.description, tc.expected, got)
		}
	}
}