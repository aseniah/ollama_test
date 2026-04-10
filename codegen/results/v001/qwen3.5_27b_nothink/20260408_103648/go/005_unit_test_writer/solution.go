package main

import (
	"fmt"
	"math"
)

// calculateDiscount calculates the bulk discount based on price and quantity.
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

// testCase defines a single test case for the discount calculator.
type testCase struct {
	description string
	price       float64
	quantity    int
	expected    float64
}

func main() {
	// Define test cases covering boundary conditions and standard ranges
	tests := []testCase{
		{
			description: "No discount for quantity 9 (under 10)",
			price:       100.0,
			quantity:    9,
			expected:    100.0 * 9.0, // 900.0
		},
		{
			description: "10% discount for quantity 10 (boundary)",
			price:       100.0,
			quantity:    10,
			expected:    100.0 * 10.0 * 0.9, // 900.0
		},
		{
			description: "10% discount for quantity 49 (upper bound of 10-49)",
			price:       100.0,
			quantity:    49,
			expected:    100.0 * 49.0 * 0.9, // 4410.0
		},
		{
			description: "25% discount for quantity 50 (boundary)",
			price:       100.0,
			quantity:    50,
			expected:    100.0 * 50.0 * 0.75, // 3750.0
		},
	}

	// Run tests
	for _, tc := range tests {
		got := calculateDiscount(tc.price, tc.quantity)
		
		// Use a small epsilon for floating point comparison to handle precision issues
		if math.Abs(got-tc.expected) < 1e-9 {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", tc.description, tc.expected, got)
		}
	}
}