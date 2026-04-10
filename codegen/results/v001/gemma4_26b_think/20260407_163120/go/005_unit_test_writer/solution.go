package main

import (
	"fmt"
)

// calculateDiscount calculates the total price based on quantity-based bulk discounts.
func calculateDiscount(price float6loat64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

// testCase represents a single unit test definition
type testCase struct {
	description string
	price       float64
	quantity    int
	expected    float64
}

func main() {
	// Defining the test cases including the requested boundary conditions: 9, 10, 49, 50
	tests := []testCase{
		{
			description: "Quantity 9 (No discount)",
			price:       100.0,
			quantity:    9,
			expected:    900.0,
		},
		{
			description: "Quantity 10 (10% discount boundary)",
			price:       100.0,
			quantity:    10,
			expected:    900.0, // 100 * 10 * 0.9
		},
		{
			description: "Quantity 49 (10% discount boundary)",
			price:       100.0,
			quantity:    49,
			expected:    4410.0, // 100 * 49 * 0.9
		},
		{
			description: "Quantity 50 (25% discount boundary)",
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

	for _, tc := range tests {
		got := calculateDiscount(tc.price, tc.quantity)

		// Using a small epsilon for floating point comparison to avoid precision issues
		const epsilon = 1e-9
		diff := got - tc.expected
		if diff < 0 {
			diff = -diff
		}

		if diff < epsilon {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %g, got: %g)\n", tc.description, tc.expected, got)
		}
	}
}