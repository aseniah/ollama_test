package main

import (
	"fmt"
	"math"
)

// calculateDiscount calculates the total price with bulk discount
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
		{"price=10.0, quantity=9", 10.0, 9, 90.0},
		{"price=10.0, quantity=10", 10.0, 10, 90.0},
		{"price=10.0, quantity=49", 10.0, 49, 441.0},
		{"price=10.0, quantity=50", 10.0, 50, 375.0},
	}

	for _, tc := range testCases {
		result := calculateDiscount(tc.price, tc.quantity)
		// Allow a small floating-point error (epsilon)
		epsilon := 1e-6
		if math.Abs(result-tc.expected) < epsilon {
			fmt.Printf("PASS: %s\n", tc.desc)
		} else {
			fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.desc, tc.expected, result)
		}
	}
}