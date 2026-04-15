package main

import (
	"fmt"
)

// calculateDiscount calculates a bulk discount based on the quantity purchased.
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
	// Define test cases including boundary conditions: 9, 10, 49, 50
	tests := []testCase{
		{
			description: "No discount for quantity < 10 (Boundary: 9)",
			price:       100.0,
			quantity:    9,
			expected:    900.0,
		},
		{
			description: "10% discount for quantity >= 10 (Boundary: 10)",
			price:       100.0,
			quantity:    10,
			expected:    900.0, // 100 * 10 * 0.9
		},
		{
			description: "10% discount for quantity < 50 (Boundary: 49)",
			price:       100.0,
			quantity:    49,
			expected:    4410.0, // 100 * 49 * 0.9
		},
		{
			description: "25% discount for quantity >= 50 (Boundary: 50)",
			price:       100.0,
			quantity:    50,
			expected:    3750.0, // 100 * 50 * 0.75
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