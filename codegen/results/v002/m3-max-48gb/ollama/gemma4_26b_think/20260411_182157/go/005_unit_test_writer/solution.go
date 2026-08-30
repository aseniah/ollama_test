package main

import (
	"fmt"
	"math"
)

// calculateDiscount calculates the total price based on quantity-based discounts.
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

// testCase defines a structure for our unit tests
type testCase struct {
	description string
	price       float64
	quantity    int
	expected    float64
}

func main() {
	// Define the test suite including the required boundary conditions
	tests := []testCase{
		{
			description: "Quantity 9 (No discount)",
			price:       100.0,
			quantity:    9,
			expected:    900.0,
		},
		{
			description: "Quantity 10 (Boundary: 10% discount applied)",
			price:       100.0,
			quantity:    10,
			expected:    900.0,
		},
		{
			description: "Quantity 49 (Boundary: 10% discount applied)",
			price:       100.0,
			quantity:    49,
			expected:    4410.0,
		},
		{
			description: "Quantity 50 (Boundary: 25% discount applied)",
			price:       100.0,
			quantity:    50,
			expected:    3750.0,
		},
		{
			description: "Quantity 100 (25% discount applied)",
			price:       10.0,
			quantity:    100,
			expected:    750.0,
		},
	}

	// Iterate through each test case
	for _, tc := range tests {
		got := calculateDiscount(tc.price, tc.quantity)

		// Use a small epsilon for comparing floating point numbers to avoid precision issues
		const epsilon = 1e-9
		if math.Abs(got-tc.expected) < epsilon {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			// Format output to 2 decimal places for readability
			fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.description, tc.expected, got)
		}
	}
}