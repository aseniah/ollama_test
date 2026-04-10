package main

import "fmt"

// calculateDiscount calculates the total price with bulk discounts.
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
		// Test Case 1: Boundary condition (below 10) - No discount
		{
			description: "Quantity 9 (no discount)",
			price:       10.0,
			quantity:    9,
			expected:    90.0, // 10 * 9
		},
		// Test Case 2: Boundary condition (at 10) - 10% discount
		{
			description: "Quantity 10 (10% discount)",
			price:       10.0,
			quantity:    10,
			expected:    90.0, // 10 * 10 * 0.9
		},
		// Test Case 3: Boundary condition (below 50) - 10% discount
		{
			description: "Quantity 49 (10% discount)",
			price:       20.0,
			quantity:    49,
			expected:    882.0, // 20 * 49 * 0.9
		},
		// Test Case 4: Boundary condition (at 50) - 25% discount
		{
			description: "Quantity 50 (25% discount)",
			price:       10.0,
			quantity:    50,
			expected:    375.0, // 10 * 50 * 0.75
		},
	}

	for _, tc := range tests {
		got := calculateDiscount(tc.price, tc.quantity)

		// Use a small epsilon for float comparison to handle potential precision issues
		const epsilon = 0.001
		diff := got - tc.expected
		if diff < 0 {
			diff = -diff
		}

		if diff <= epsilon {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", tc.description, tc.expected, got)
		}
	}
}