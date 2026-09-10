package main

import (
	"fmt"
)

// calculateDiscount calculates the total price based on the quantity ordered.
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

type testCase struct {
	description string
	price       float64
	quantity    int
	expected    float64
}

func main() {
	tests := []testCase{
		{
			description: "No discount for quantity < 10 (Boundary 9)",
			price:       10.0,
			quantity:    9,
			expected:    90.0,
		},
		{
			description: "10% discount for quantity >= 10 (Boundary 10)",
			price:       10.0,
			quantity:    10,
			expected:    90.0, // 10 * 10 * 0.9
		},
		{
			description: "10% discount for quantity < 50 (Boundary 49)",
			price:       10.0,
			quantity:    49,
			expected:    441.0, // 10 * 49 * 0.9
		},
		{
			description: "25% discount for quantity >= 50 (Boundary 50)",
			price:       10.0,
			quantity:    50,
			expected:    375.0, // 10 * 50 * 0.75
		},
		{
			description: "High volume discount",
			price:       100.0,
			quantity:    100,
			expected:    7500.0, // 100 * 100 * 0.75
		},
	}

	for _, tc := range tests {
		actual := calculateDiscount(tc.price, tc.quantity)
		if actual == tc.expected {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.description, tc.expected, actual)
		}
	}
}