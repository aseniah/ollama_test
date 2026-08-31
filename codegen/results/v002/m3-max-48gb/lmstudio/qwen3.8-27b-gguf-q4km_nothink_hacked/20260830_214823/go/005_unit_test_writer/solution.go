package main

import (
	"fmt"
	"math"
)

func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

func main() {
	type testCase struct {
		description string
		price       float64
		quantity    int
		expected    float64
	}

	testCases := []testCase{
		{
			description: "quantity 9 (below discount threshold)",
			price:       10.0,
			quantity:    9,
			expected:    90.0,
		},
		{
			description: "quantity 10 (at 10% discount threshold)",
			price:       10.0,
			quantity:    10,
			expected:    10.0 * 10 * 0.9,
		},
		{
			description: "quantity 49 (just below 25% discount threshold)",
			price:       10.0,
			quantity:    49,
			expected:    10.0 * 49 * 0.9,
		},
		{
			description: "quantity 50 (at 25% discount threshold)",
			price:       10.0,
			quantity:    50,
			expected:    10.0 * 50 * 0.75,
		},
	}

	for _, tc := range testCases {
		got := calculateDiscount(tc.price, tc.quantity)
		if math.Abs(got-tc.expected) < 1e-9 {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", tc.description, tc.expected, got)
		}
	}
}