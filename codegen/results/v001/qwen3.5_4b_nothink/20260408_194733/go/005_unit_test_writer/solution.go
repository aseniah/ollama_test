package main

import (
	"fmt"
	"math"
)

// calculateDiscount calculates a bulk discount
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

func main() {
	testCases := []struct {
		desc     string
		price    float64
		quantity int
		expected float64
	}{
		{desc: "single item price", price: 10.0, quantity: 1, expected: 10.0},
		{desc: "boundary below 10", price: 10.0, quantity: 9, expected: 90.0},
		{desc: "boundary at 10", price: 10.0, quantity: 10, expected: 90.0},
		{desc: "boundary below 50", price: 10.0, quantity: 49, expected: 441.0},
		{desc: "boundary at 50", price: 10.0, quantity: 50, expected: 375.0},
		{desc: "large quantity", price: 10.0, quantity: 100, expected: 750.0},
	}

	for _, tc := range testCases {
		result := calculateDiscount(tc.price, tc.quantity)
		// Use a small epsilon for float comparison
		epsilon := 1e-6
		isMatch := math.Abs(result-tc.expected) < epsilon

		if isMatch {
			fmt.Printf("PASS: %s\n", tc.desc)
		} else {
			fmt.Printf("FAIL: %s (expected: %.1f, got: %.1f)\n", tc.desc, tc.expected, result)
		}
	}
}