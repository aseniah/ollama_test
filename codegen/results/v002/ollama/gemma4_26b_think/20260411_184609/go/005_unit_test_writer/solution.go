package main

import (
	"fmt"
	"math"
)

// calculateDiscount calculates the bulk discount based on the quantity provided.
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

// TestCase defines the structure for our test suite
type TestCase struct {
	description string
	price       float64
	quantity    int
	expected    float64
}

func main() {
	// Define the test cases including the required boundary conditions
	tests := []TestCase{
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
			expected:    900.0,
		},
		{
			description: "Quantity 49 (10% discount boundary)",
			price:       100.0,
			quantity:    49,
			expected:    4410.0,
		},
		{
			description: "Quantity 50 (25% discount boundary)",
			price:       100.0,
			quantity:    50,
			expected:    3750.0,
		},
		{
			description: "Quantity 100 (25% discount)",
			price:       10.0,
			quantity:    100,
			expected:    750.0,
		},
	}

	// Small epsilon for floating point comparison
	const epsilon = 1e-9

	for _, tc := range tests {
		got := calculateDiscount(tc.price, tc.quantity)

		// Check if the difference is within a tiny threshold to handle float precision
		if math.Abs(got-tc.expected) < epsilon {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			// Formatting to 2 decimal places for readable error output
			fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.description, tc.expected, got)
		}
	}
}