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
	testCases := []struct {
		description string
		price       float64
		quantity    int
		expected    float64
	}{
		{
			description: "Quantity 9 (no discount)",
			price:       10.0,
			quantity:    9,
			expected:    90.0, // 10 * 9
		},
		{
			description: "Quantity 10 (10% discount boundary)",
			price:       10.0,
			quantity:    10,
			expected:    90.0, // 10 * 10 * 0.9
		},
		{
			description: "Quantity 49 (10% discount upper boundary)",
			price:       10.0,
			quantity:    49,
			expected:    441.0, // 10 * 49 * 0.9
		},
		{
			description: "Quantity 50 (25% discount boundary)",
			price:       10.0,
			quantity:    50,
			expected:    375.0, // 10 * 50 * 0.75
		},
	}

	allPassed := true

	for _, tc := range testCases {
		got := calculateDiscount(tc.price, tc.quantity)

		// Floating point comparison with a small epsilon for safety, 
		// though these inputs result in exact decimals.
		epsilon := 0.0001
		if abs(got-tc.expected) > epsilon {
			fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.description, tc.expected, got)
			allPassed = false
		} else {
			fmt.Printf("PASS: %s\n", tc.description)
		}
	}

	if allPassed {
		fmt.Println("\nAll tests passed.")
	} else {
		fmt.Println("\nSome tests failed.")
	}
}

// abs returns the absolute value of a float64
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}