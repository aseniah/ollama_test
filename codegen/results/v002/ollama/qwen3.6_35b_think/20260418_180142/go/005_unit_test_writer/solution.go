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
		{"Quantity 9 (boundary < 10)", 100, 9, 900.0},
		{"Quantity 10 (boundary >= 10)", 100, 10, 900.0},
		{"Quantity 49 (boundary < 50)", 100, 49, 4410.0},
		{"Quantity 50 (boundary >= 50)", 100, 50, 3750.0},
	}

	for _, tc := range testCases {
		result := calculateDiscount(tc.price, tc.quantity)
		if math.Abs(tc.expected-result) < 1e-9 {
			fmt.Printf("PASS: %s\n", tc.desc)
		} else {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", tc.desc, tc.expected, result)
		}
	}
}