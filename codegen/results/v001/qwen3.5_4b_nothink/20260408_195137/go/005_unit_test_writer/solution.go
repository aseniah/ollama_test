package main

import (
	"fmt"
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
		{"Quantity 9 (no discount)", 10.0, 9, 90.0},
		{"Quantity 10 (90% discount)", 10.0, 10, 90.0},
		{"Quantity 49 (90% discount)", 10.0, 49, 441.0},
		{"Quantity 50 (75% discount)", 10.0, 50, 375.0},
	}

	for _, tc := range testCases {
		result := calculateDiscount(tc.price, tc.quantity)

		if fmt.Sprintf("%.1f", result) == fmt.Sprintf("%.1f", tc.expected) {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %.1f, got: %.1f)\n", tc.description, tc.expected, result)
		}
	}
}