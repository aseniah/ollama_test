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

type testCase struct {
	description string
	price       float64
	quantity    int
	expected    float64
}

func main() {
	// Define test cases covering boundary conditions
	tests := []testCase{
		{
			description: "Quantity just below first threshold (9)",
			price:       10.0,
			quantity:    9,
			expected:    90.0, // 10 * 9
		},
		{
			description: "Quantity exactly at first threshold (10)",
			price:       10.0,
			quantity:    10,
			expected:    90.0, // 10 * 10 * 0.9
		},
		{
			description: "Quantity just below second threshold (49)",
			price:       10.0,
			quantity:    49,
			expected:    441.0, // 10 * 49 * 0.9
		},
		{
			description: "Quantity exactly at second threshold (50)",
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