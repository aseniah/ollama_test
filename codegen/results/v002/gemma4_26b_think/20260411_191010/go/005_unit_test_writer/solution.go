package main

import (
	"fmt"
)

// calculateDiscount calculates a bulk discount based on quantity.
// quantity < 10: No discount.
// 10 <= quantity < 50: 10% discount.
// quantity >= 50: 25% discount.
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

// testCase defines the structure for our unit tests.
type testCase struct {
	description string
	price       float66
	quantity    int
	expected    float64
}

func main() {
	// Defining test cases including the required boundary conditions.
	tests := []testCase{
		{
			description: "quantity 9 (no discount boundary)",
			price:       100.0,
			quantity:    9,
			expected:    900.0,
		},
		{
			description: "quantity 10 (10% discount boundary)",
			price:       100.0,
			quantity:    10,
			expected:    900.0,
		},
		{
			description: "quantity 49 (10% discount boundary)",
			price:       100.0,
			quantity:    49,
			expected:    4410.0,
		},
		{
			description: "quantity 50 (25% discount boundary)",
			price:       100.0,
			quantity:    50,
			expected:    3750.0,
		},
		{
			description: "large quantity (25% discount)",
			price:       10.0,
			quantity:    100,
			expected:    750.0,
		},
	}

	// Run the tests.
	for _, tc := range tests {
		got := calculateDiscount(tc.price, tc.quantity)

		// In real-world float comparisons, we might use a small epsilon.
		// For these specific calculation values, direct comparison works.
		if got == tc.expected {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			// Using %g to format the float clearly in the output.
			fmt.Printf("FAIL: %s (expected: %g, got: %g)\n", tc.description, tc.expected, got)
		}
	}
}