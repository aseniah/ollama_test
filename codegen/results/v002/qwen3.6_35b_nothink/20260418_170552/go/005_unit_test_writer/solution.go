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
		description string
		price       float64
		quantity    int
		expected    float64
	}{
		{
			description: "quantity < 10 (quantity=9)",
			price:       100.0,
			quantity:    9,
			expected:    900.0,
		},
		{
			description: "quantity = 10 (boundary between <10 and 10-49)",
			price:       100.0,
			quantity:    10,
			expected:    900.0,
		},
		{
			description: "quantity = 49 (boundary of upper range 10-49)",
			price:       100.0,
			quantity:    49,
			expected:    4410.0,
		},
		{
			description: "quantity = 50 (boundary of lower range 50+)",
			price:       100.0,
			quantity:    50,
			expected:    3750.0,
		},
	}

	for _, tc := range tests {
		result := calculateDiscount(tc.price, tc.quantity)
		// Use a small epsilon for float comparison
		epsilon := 1e-9
		if result < tc.expected-epsilon || result > tc.expected+epsilon {
			fmt.Printf("FAIL: %s (expected: %.4f, got: %.4f)\n", tc.description, tc.expected, result)
		} else {
			fmt.Printf("PASS: %s\n", tc.description)
		}
	}
}