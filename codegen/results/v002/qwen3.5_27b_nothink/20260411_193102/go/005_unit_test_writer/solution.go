package main

import (
	"fmt"
	"math"
)

// calculateDiscount calculates the total price based on quantity and unit price.
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

func main() {
	// Define test cases
	// Description format: "Qty: <quantity>, Price: <price>"
	testCases := []struct {
		description string
		price       float64
		quantity    int
		expected    float64
	}{
		{
			description: "Qty: 9 (No Discount)",
			price:       10.0,
			quantity:    9,
			expected:    90.0, // 10 * 9
		},
		{
			description: "Qty: 10 (10% Discount Boundary)",
			price:       10.0,
			quantity:    10,
			expected:    90.0, // 10 * 10 * 0.9
		},
		{
			description: "Qty: 49 (10% Discount Upper Limit)",
			price:       10.0,
			quantity:    49,
			expected:    441.0, // 10 * 49 * 0.9
		},
		{
			description: "Qty: 50 (25% Discount Boundary)",
			price:       10.0,
			quantity:    50,
			expected:    375.0, // 10 * 50 * 0.75
		},
	}

	// Run tests
	for _, tc := range testCases {
		got := calculateDiscount(tc.price, tc.quantity)

		// Use a small epsilon for float comparison to handle potential precision issues
		if math.Abs(got-tc.expected) < 0.0001 {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.description, tc.expected, got)
		}
	}
}