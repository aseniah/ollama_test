package main

import (
	"fmt"
)

// calculateDiscount calculates the total price based on quantity-based bulk discounts.
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

// TestCase represents a single test scenario
type TestCase struct {
	description string
	price       float64
	quantity    int
	expected    float64
}

func main() {
	// Define the test cases including boundary conditions
	tests := []TestCase{
		{
			description: "Small quantity (below 10)",
			price:       10.0,
			quantity:    9,
			expected:    90.0,
		},
		{
			description: "Boundary case (exactly 10) - 10% discount starts",
			price:       10.0,
			quantity:    10,
			expected:    90.0, // 10 * 10 * 0.9
		},
		{
			description: "Mid-range quantity (below 50)",
			price:       10.0,
			quantity:    49,
			expected:    441.0, // 10 * 49 * 0.9
		},
		{
			description: "Boundary case (exactly 50) - 25% discount starts",
			price:       10.0,
			quantity:    50,
			expected:    375.0, // 10 * 50 * 0.75
		},
		{
			description: "Large quantity (above 50)",
			price:       100.0,
			quantity:    100,
			expected:    7500.0, // 100 * 100 * 0.75
		},
	}

	// Run the tests
	for _, tc := range tests {
		got := calculateDiscount(tc.price, tc.quantity)

		// Using a small epsilon for float comparison to handle precision issues
		const epsilon = 0.000001
		diff := got - tc.expected
		if diff < 0 {
			diff = -diff
		}

		if diff <= epsilon {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.description, tc.expected, got)
		}
	}
}