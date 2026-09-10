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

func main() {
	// Define test cases with boundary conditions
	type testCase struct {
		description string
		price       float64
		quantity    int
		expected    float64
	}

	tests := []testCase{
		{"quantity 9 (no discount boundary)", 10.0, 9, 10.0 * 9},
		{"quantity 10 (10% discount boundary)", 10.0, 10, 10.0 * 10 * 0.9},
		{"quantity 49 (10% discount boundary)", 10.0, 49, 10.0 * 49 * 0.9},
		{"quantity 50 (25% discount boundary)", 10.0, 50, 10.0 * 50 * 0.75},
		{"quantity 1 (no discount, single item)", 5.5, 1, 5.5 * 1},
		{"quantity 100 (25% discount, large qty)", 20.0, 100, 20.0 * 100 * 0.75},
	}

	for _, tc := range tests {
		got := calculateDiscount(tc.price, tc.quantity)
		if got == tc.expected {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %f, got: %f)\n", tc.description, tc.expected, got)
		}
	}
}