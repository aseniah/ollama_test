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
	tests := []struct {
		desc     string
		price    float64
		quantity int
		expected float64
	}{
		{"quantity 9", 100.0, 9, 900.0},
		{"quantity 10", 100.0, 10, 900.0},
		{"quantity 49", 100.0, 49, 441.0},
		{"quantity 50", 100.0, 50, 375.0},
	}

	for _, tc := range tests {
		result := calculateDiscount(tc.price, tc.quantity)
		if result == tc.expected {
			fmt.Printf("PASS: %s\n", tc.desc)
		} else {
			fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.desc, tc.expected, result)
		}
	}
}