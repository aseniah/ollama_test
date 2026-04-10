package main

import (
	"fmt"
	"math"
)

// calculateDiscount calculates a bulk discount based on the quantity provided.
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

// testCase defines the structure for our automated tests
type testCase struct {
	description string
	price       float64
	quantity    int
	expected    float64
}

func main() {
	// Define the test cases including the required boundary conditions: 9, 10, 49, 50
	tests := []testCase{
		{
			description: "Quantity 9 (no discount)",
			price:       100.0,
			quantity:    9,
			expected:    900.0,
		},
		{
			description: "Quantity 10 (10% discount)",
			price:       100.0,
			quantity:    10,
			expected:    900.0,
		},
		{
			description: "Quantity 49 (10% discount)",
			price:       100.0,
			quantity:    49,
			expected:    4410.0,
		},
		{
			description: "Quantity 50 (25% discount)",
			price:       100.0,
			quantity:    50,
			expected:    3750.0,
		},
	}

	// Iterate through each test case
	for _, tc := range tests {
		got := calculateDiscount(tc.price, tc.quantity)

		// Use a small epsilon to account for floating-point precision errors
		const epsilon = 1e-9
		if math.Abs(got-tc.expected) < epsilon {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			// Using %v to format the output as requested
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", tc.description, tc.expected, got)
		}
	}
}