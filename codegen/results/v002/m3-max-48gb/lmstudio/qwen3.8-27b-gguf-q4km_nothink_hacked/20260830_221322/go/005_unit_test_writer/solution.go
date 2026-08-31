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
		price    float64
		quantity int
		expected float64
		desc     string
	}

	testCases := []testCase{
		{price: 10.0, quantity: 9, expected: 10.0 * 9, desc: "quantity 9 (below 10 discount)"},
		{price: 10.0, quantity: 10, expected: 10.0 * 10 * 0.9, desc: "quantity 10 (exactly 10, 10% discount)"},
		{price: 10.0, quantity: 49, expected: 10.0 * 49 * 0.9, desc: "quantity 49 (just below 50, 10% discount)"},
		{price: 10.0, quantity: 50, expected: 10.0 * 50 * 0.75, desc: "quantity 50 (exactly 50, 25% discount)"},
	}

	for _, tc := range testCases {
		got := calculateDiscount(tc.price, tc.quantity)
		if math.Abs(got-tc.expected) < 1e-9 {
			fmt.Printf("PASS: %s\n", tc.desc)
		} else {
			fmt.Printf("FAIL: %s (expected: %f, got: %f)\n", tc.desc, tc.expected, got)
		}
	}
}