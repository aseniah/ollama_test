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

type testCase struct {
	description string
	price       float64
	quantity    int
	expected    float64
}

func main() {
	testCases := []testCase{
		{
			description: "quantity 9 (below 10, no discount)",
			price:       10.0,
			quantity:    9,
			expected:    90.0,
		},
		{
			description: "quantity 10 (boundary: 10% discount)",
			price:       10.0,
			quantity:    10,
			expected:    90.0,
		},
		{
			description: "quantity 49 (just below 50, 10% discount)",
			price:       10.0,
			quantity:    49,
			expected:    441.0,
		},
		{
			description: "quantity 50 (boundary: 25% discount)",
			price:       10.0,
			quantity:    50,
			expected:    375.0,
		},
	}

	for _, tc := range testCases {
		got := calculateDiscount(tc.price, tc.quantity)

		if math.Abs(got-tc.expected) < 1e-6 {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %.6f, got: %.6f)\n", tc.description, tc.expected, got)
		}
	}
}