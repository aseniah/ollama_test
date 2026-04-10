package main

import (
	"fmt"
)

// calculateDiscount calculates the total price after applying bulk discounts.
// < 10: No discount
// 10 to 49: 10% discount
// 50+: 25% discount
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

// TestCase defines the structure for our unit tests
type TestCase struct {
	description string
	price       float64
	quantity    int
	expected    float64
}

func main() {
	// Define test cases covering boundary conditions
	tests := []TestCase{
		{
			description: "Boundary: quantity 9 (No discount)",
			price:       100.0,
			quantity:    9,
			expected:    900.0,
		},
		{
			description: "Boundary: quantity 10 (10% discount begins)",
			price:       100.0,
			quantity:    10,
			expected:    900.0, // 100 * 10 * 0.9
		},
		{
			description: "Boundary: quantity 49 (10% discount end)",
			price:       100.0,
			quantity:    49,
			expected:    4410.0, // 100 * 49 * 0.9
		},
		{
			description: "Boundary: quantity 50 (25% discount begins)",
			price:       100.0,
			quantity:    50,
			expected:    3750.0, // 100 * 50 * 0.75
		},
		{
			description: "Large quantity (25% discount)",
			price:       10.0,
			quantity:    100,
			expected:    750.0, // 10 * 100 * 0.75
		},
	}

	// Iterate through tests and validate results
	for _, tc := range tests {
		got := calculateDiscount(tc.price, tc.quantity)

		// Using a small epsilon for float comparison to avoid precision issues
		const epsilon = 0.000001
		diff := got - tc.expected
		if diff < 0 {
			diff = -diff
		}

		if diff <= epsilon {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.description, tc.expected, got)
		}
	}
}