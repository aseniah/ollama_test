package main

import (
	"fmt"
)

// calculateDiscount calculates the total price based on quantity discounts:
// < 10: No discount
// 10 - 49: 10% discount
// 50+: 25% discount
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
			description: "No discount for quantity under 10 (boundary 9)",
			price:       100.0,
			quantity:    9,
			expected:    900.0,
		},
		{
			description: "10% discount for quantity 10 (boundary 10)",
			price:       100.0,
			quantity:    10,
			expected:    900.0, // 100 * 10 * 0.9
		},
		{
			description: "10% discount for quantity under 50 (boundary 49)",
			price:       100.0,
			quantity:    49,
			expected:    4410.0, // 100 * 49 * 0.9
		},
		{
			description: "25% discount for quantity 50 (boundary 50)",
			price:       100.0,
			quantity:    50,
			expected:    3750.0, // 100 * 50 * 0.75
		},
		{
			description: "25% discount for quantity over 50",
			price:       100.0,
			quantity:    100,
			expected:    7500.0, // 100 * 100 * 0.75
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