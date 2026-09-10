package main

import (
	"fmt"
)

// calculateDiscount calculates the total price based on quantity thresholds.
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

// testCase defines the structure for our test data
type testCase struct {
	description string
	price       float64
	quantity    int
	expected    float64
}

func main() {
	// Define test cases including the requested boundary conditions
	tests := []testCase{
		{
			description: "Quantity 9 (no discount)",
			price:       10.0,
			quantity:    9,
			expected:    90.0,
		},
		{
			description: "Quantity 10 (10% discount boundary)",
			price:       10.0,
			quantity:    10,
			expected:    90.0, // 10 * 10 * 0.9 = 90
		},
		{
			description: "Quantity 49 (10% discount boundary)",
			price:       10.0,
			quantity:    49,
			expected:    441.0, // 10 * 49 * 0.9 = 441
		},
		{
			description: "Quantity 50 (25% discount boundary)",
			price:       10.0,
			quantity:    50,
			expected:    375.0, // 10 * 50 * 0.75 = 375
		},
		{
			description: "Quantity 100 (25% discount)",
			price:       2.0,
			quantity:    100,
			expected:    150.0, // 2 * 100 * 0.75 = 150
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