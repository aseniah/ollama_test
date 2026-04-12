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
	// Define test cases: price, quantity, expected result, description
	testCases := []struct {
		price     float64
		quantity  int
		expected  float64
		description string
	}{
		{10.0, 9, 90.0, "No discount (quantity 9)"},
		{10.0, 10, 90.0, "10% discount (quantity 10)"},
		{10.0, 49, 441.0, "10% discount (quantity 49)"},
		{10.0, 50, 375.0, "25% discount (quantity 50)"},
	}

	for _, tc := range testCases {
		got := calculateDiscount(tc.price, tc.quantity)

		// Floating point comparison with a small epsilon for safety
		epsilon := 0.0001
		if abs(got-tc.expected) < epsilon {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", tc.description, tc.expected, got)
		}
	}
}

// Helper function to get absolute value
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}