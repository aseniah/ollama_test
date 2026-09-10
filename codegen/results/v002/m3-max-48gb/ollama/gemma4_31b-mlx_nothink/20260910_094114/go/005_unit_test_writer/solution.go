package main

import (
	"fmt"
)

// calculateDiscount calculates the total price based on quantity tiers.
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

// testCase defines the input and expected output for a test
type testCase struct {
	description string
	price       float64
	quantity    int
	expected    float64
}

func main() {
	// Define test cases covering boundaries: 9, 10, 49, 50
	tests := []testCase{
		{
			description: "No discount for quantity < 10 (9 units)",
			price:       100.0,
			quantity:    9,
			expected:    900.0,
		},
		{
			description: "10% discount for quantity 10 (lower boundary of 2nd tier)",
			price:       100.0,
			quantity:    10,
			expected:    900.0, // 10 * 100 * 0.9
		},
		{
			description: "10% discount for quantity 49 (upper boundary of 2nd tier)",
			price:       100.0,
			quantity:    49,
			expected:    4410.0, // 49 * 100 * 0.9
		},
		{
			description: "25% discount for quantity 50 (lower boundary of 3rd tier)",
			price:       100.0,
			quantity:    50,
			expected:    3750.0, // 50 * 100 * 0.75
		},
	}

	for _, tc := range tests {
		result := calculateDiscount(tc.price, tc.quantity)
		if result == tc.expected {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.description, tc.expected, result)
		}
	}
}