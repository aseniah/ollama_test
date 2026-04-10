package main

import "fmt"

// calculateDiscount calculates the total price after applying bulk discounts.
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

func main() {
	// Define test cases: {description, price, quantity, expectedResult}
	tests := []struct {
		description string
		price       float64
		quantity    int
		expected    float64
	}{
		{
			description: "Quantity 9 (No discount)",
			price:       10.0,
			quantity:    9,
			expected:    10.0 * 9.0, // 90.0
		},
		{
			description: "Quantity 10 (10% discount boundary)",
			price:       10.0,
			quantity:    10,
			expected:    10.0 * 10.0 * 0.9, // 90.0
		},
		{
			description: "Quantity 49 (10% discount boundary)",
			price:       10.0,
			quantity:    49,
			expected:    10.0 * 49.0 * 0.9, // 441.0
		},
		{
			description: "Quantity 50 (25% discount boundary)",
			price:       10.0,
			quantity:    50,
			expected:    10.0 * 50.0 * 0.75, // 375.0
		},
	}

	allPassed := true

	for _, tc := range tests {
		got := calculateDiscount(tc.price, tc.quantity)

		// Allow for small floating point precision differences
		const epsilon = 0.0001
		diff := got - tc.expected
		if diff < 0 {
			diff = -diff
		}

		if diff > epsilon {
			fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.description, tc.expected, got)
			allPassed = false
		} else {
			fmt.Printf("PASS: %s\n", tc.description)
		}
	}

	if allPassed {
		// Optional: Explicitly indicate all tests passed if needed, 
		// though the requirement just asked to print per test case.
	}
}