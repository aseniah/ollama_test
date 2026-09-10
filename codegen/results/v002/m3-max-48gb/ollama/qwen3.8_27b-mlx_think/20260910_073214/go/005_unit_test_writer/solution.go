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
	tests := []struct {
		description string
		price       float64
		quantity    int
		expected    float64
	}{
		{"No discount for quantity below 10 (boundary: 9)", 10.0, 9, 10.0 * 9},
		{"10% discount starts at quantity 10 (boundary: 10)", 10.0, 10, 10.0 * 10 * 0.9},
		{"10% discount for quantity 49 (boundary: 49)", 10.0, 49, 10.0 * 49 * 0.9},
		{"25% discount starts at quantity 50 (boundary: 50)", 10.0, 50, 10.0 * 50 * 0.75},
		{"No discount for quantity 1", 5.5, 1, 5.5 * 1},
		{"25% discount for large quantity (quantity 100)", 20.0, 100, 20.0 * 100 * 0.75},
	}

	const epsilon = 1e-9

	for _, tc := range tests {
		got := calculateDiscount(tc.price, tc.quantity)
		if math.Abs(got-tc.expected) < epsilon {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %f, got: %f)\n", tc.description, tc.expected, got)
		}
	}
}