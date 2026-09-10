package main

import (
	"fmt"
)

// calculateDiscount calculates a bulk discount based on the quantity ordered.
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

func main() {
	// Define a test case structure
	type testCase struct {
		description string
		price       float64
		quantity    int
		expected    float64
	}

	// Define the test cases, including the requested boundary conditions
	tests := []testCase{
		{
			description: "Quantity just below 10 (no discount)",
			price:       10.0,
			quantity:    9,
			expected:    90.0,
		},
		{
			description: "Quantity exactly 10 (10% discount)",
			price:       10.0,
			quantity:    10,
			expected:    90.0, // 10 * 10 * 0.9 = 90
		},
		{
			description: "Quantity just below 50 (10% discount)",
			price:       10.0,
			quantity:    49,
			expected:    441.0, // 10 * 49 * 0.9 = 441
		},
		{
			description: "Quantity exactly 50 (25% discount)",
			price:       10.0,
			quantity:    50,
			expected:    375.0, // 10 * 50 * 0.75 = 375
		},
	}

	for _, tc := range tests {
		got := calculateDiscount(tc.price, tc.quantity)
		if got == tc.expected {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %f, got: %f)\n", tc.description, tc.expected, got)
		}
	}
}