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
	epsilon := 1e-9

	tests := []struct {
		description string
		price       float64
		quantity    int
		expected    float64
	}{
		{"quantity 9 (below 10, no discount)", 10.0, 9, 10.0 * 9},
		{"quantity 10 (boundary, 10% discount)", 10.0, 10, 10.0 * 10 * 0.9},
		{"quantity 49 (below 50, 10% discount)", 10.0, 49, 10.0 * 49 * 0.9},
		{"quantity 50 (boundary, 25% discount)", 10.0, 50, 10.0 * 50 * 0.75},
		{"quantity 100 (well above 50, 25% discount)", 5.0, 100, 5.0 * 100 * 0.75},
	}

	for _, tc := range tests {
		got := calculateDiscount(tc.price, tc.quantity)
		if math.Abs(got-tc.expected) < epsilon {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", tc.description, tc.expected, got)
		}
	}
}