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
		description string
		price       float64
		quantity    int
		expected    float64
	}{
		{"Quantity 9 (no discount)", 100.0, 9, 900.0},
		{"Quantity 10 (10% discount)", 100.0, 10, 900.0},
		{"Quantity 49 (10% discount)", 100.0, 49, 4410.0},
		{"Quantity 50 (25% discount)", 100.0, 50, 3750.0},
	}

	for _, tc := range testCases {
		got := calculateDiscount(tc.price, tc.quantity)
		if math.Abs(got-tc.expected) < 0.001 {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.description, tc.expected, got)
		}
	}
}