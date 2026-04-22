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
			description: "quantity 9 (no discount)",
			price:       10.0,
			quantity:    9,
			expected:    90.0,
		},
		{
			description: "quantity 10 (10% discount)",
			price:       10.0,
			quantity:    10,
			expected:    90.0,
		},
		{
			description: "quantity 49 (10% discount)",
			price:       10.0,
			quantity:    49,
			expected:    441.0,
		},
		{
			description: "quantity 50 (25% discount)",
			price:       10.0,
			quantity:    50,
			expected:    375.0,
		},
		{
			description: "quantity 100 (25% discount)",
			price:       10.0,
			quantity:    100,
			expected:    750.0,
		},
	}

	for _, tc := range tests {
		result := calculateDiscount(tc.price, tc.quantity)
		if result == tc.expected {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %.4f, got: %.4f)\n", tc.description, tc.expected, result)
		}
	}
}