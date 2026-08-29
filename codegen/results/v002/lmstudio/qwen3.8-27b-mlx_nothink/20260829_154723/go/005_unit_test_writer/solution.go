package main

import "fmt"

func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

func formatFloat(f float64) string {
	// Use a reasonable format to avoid tiny floating point differences
	// We'll round to a reasonable number of decimal places for comparison display
	return fmt.Sprintf("%.6f", f)
}

func main() {
	type testCase struct {
		description string
		price       float64
		quantity    int
		expected    float64
	}

	testCases := []testCase{
		{
			description: "quantity 9 (below 10, no discount)",
			price:       10.0,
			quantity:    9,
			expected:    10.0 * 9,
		},
		{
			description: "quantity 10 (boundary, 10% discount)",
			price:       10.0,
			quantity:    10,
			expected:    10.0 * 10 * 0.9,
		},
		{
			description: "quantity 49 (boundary, 10% discount)",
			price:       10.0,
			quantity:    49,
			expected:    10.0 * 49 * 0.9,
		},
		{
			description: "quantity 50 (boundary, 25% discount)",
			price:       10.0,
			quantity:    50,
			expected:    10.0 * 50 * 0.75,
		},
		{
			description: "quantity 1 (no discount)",
			price:       5.5,
			quantity:    1,
			expected:    5.5 * 1,
		},
		{
			description: "quantity 100 (25% discount)",
			price:       3.2,
			quantity:    100,
			expected:    3.2 * 100 * 0.75,
		},
	}

	for _, tc := range testCases {
		got := calculateDiscount(tc.price, tc.quantity)
		// Compare with a small epsilon to handle floating point
		diff := got - tc.expected
		if diff < 0 {
			diff = -diff
		}
		if diff < 1e-9 {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %s, got: %s)\n", tc.description, formatFloat(tc.expected), formatFloat(got))
		}
	}
}