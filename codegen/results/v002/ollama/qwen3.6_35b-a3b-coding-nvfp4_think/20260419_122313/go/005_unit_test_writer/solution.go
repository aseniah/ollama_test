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
	testCases := []struct {
		desc     string
		price    float64
		quantity int
		expected float64
	}{
		{"Boundary quantity 9", 100.0, 9, 900.0},
		{"Boundary quantity 10", 100.0, 10, 900.0},
		{"Boundary quantity 49", 100.0, 49, 4410.0},
		{"Boundary quantity 50", 100.0, 50, 3750.0},
	}

	for _, tc := range testCases {
		got := calculateDiscount(tc.price, tc.quantity)
		// Use a small epsilon for safe floating-point comparison
		if math.Abs(got-tc.expected) < 1e-9 {
			fmt.Printf("PASS: %s\n", tc.desc)
		} else {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", tc.desc, tc.expected, got)
		}
	}
}