package main

import (
	"fmt"
)

// calculateDiscount calculates the total price based on quantity discounts:
// - Less than 10: No discount
// - 10 to 49: 10% discount
// - 50 or more: 25% discount
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

	// Define test cases focusing on boundary conditions: 9, 10, 49, 50
	tests := []testCase{
		{
			description: "No discount for quantity 9",
			price:       10.0,
			quantity:    9,
			expected:    90.0,
		},
		{
			description: "10% discount for quantity 10 (lower boundary of second tier)",
			price:       10.0,
			quantity:    10,
			expected:    90.0, // 10 * 10 * 0.9
		},
		{
			description: "10% discount for quantity 49 (upper boundary of second tier)",
			price:       10.0,
			quantity:    49,
			expected:    441.0, // 10 * 49 * 0.9
		},
		{
			description: "25% discount for quantity 50 (lower boundary of third tier)",
			price:       10.0,
			quantity:    50,
			expected:    375.0, // 10 * 50 * 0.75
		},
	}

	for _, tc := range tests {
		actual := calculateDiscount(tc.price, tc.quantity)
		if actual == tc.expected {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %f, got: %f)\n", tc.description, tc.expected, actual)
		}
	}
}