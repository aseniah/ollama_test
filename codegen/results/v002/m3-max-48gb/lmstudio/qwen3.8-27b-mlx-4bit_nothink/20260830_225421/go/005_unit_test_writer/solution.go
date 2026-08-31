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

func approxEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
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
			description: "quantity 9, no discount",
			price:       10.0,
			quantity:    9,
			expected:    10.0 * 9, // 90
		},
		{
			description: "quantity 10, 10% discount",
			price:       10.0,
			quantity:    10,
			expected:    10.0 * 10 * 0.9, // 90
		},
		{
			description: "quantity 49, 10% discount",
			price:       10.0,
			quantity:    49,
			expected:    10.0 * 49 * 0.9, // 441
		},
		{
			description: "quantity 50, 25% discount",
			price:       10.0,
			quantity:    50,
			expected:    10.0 * 50 * 0.75, // 375
		},
	}

	for _, tc := range testCases {
		got := calculateDiscount(tc.price, tc.quantity)
		if approxEqual(got, tc.expected) {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %f, got: %f)\n", tc.description, tc.expected, got)
		}
	}
}