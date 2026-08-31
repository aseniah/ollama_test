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

func floatEqual(a, b float64) bool {
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
			expected:    90.0,
		},
		{
			description: "quantity 10, 10% discount",
			price:       10.0,
			quantity:    10,
			expected:    90.0,
		},
		{
			description: "quantity 49, 10% discount",
			price:       10.0,
			quantity:    49,
			expected:    490.0 * 0.9,
		},
		{
			description: "quantity 50, 25% discount",
			price:       10.0,
			quantity:    50,
			expected:    500.0 * 0.75,
		},
		{
			description: "quantity 1, no discount",
			price:       5.0,
			quantity:    1,
			expected:    5.0,
		},
		{
			description: "quantity 100, 25% discount",
			price:       2.5,
			quantity:    100,
			expected:    250.0 * 0.75,
		},
	}

	for _, tc := range testCases {
		got := calculateDiscount(tc.price, tc.quantity)
		if floatEqual(got, tc.expected) {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", tc.description, tc.expected, got)
		}
	}
}