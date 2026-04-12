package main

import (
	"fmt"
	"math"
)

// calculateDiscount calculates the total price based on bulk discounts.
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

// testCase holds the parameters and expected result for a single test case.
type testCase struct {
	description string
	price       float64
	quantity    int
	expected    float64
}

// runTest executes a single test case and prints the result.
func runTest(tc testCase) {
	got := calculateDiscount(tc.price, tc.quantity)
	
	// Using a small epsilon for float64 comparison to handle precision issues
	if math.Abs(got-tc.expected) < 0.0001 {
		fmt.Printf("PASS: %s\n", tc.description)
	} else {
		fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.description, tc.expected, got)
	}
}

func main() {
	tests := []testCase{
		{
			description: "Quantity 9 (No discount)",
			price:       10.0,
			quantity:    9,
			expected:    90.0, // 10 * 9 * 1.0
		},
		{
			description: "Quantity 10 (10% discount starts)",
			price:       10.0,
			quantity:    10,
			expected:    90.0, // 10 * 10 * 0.9
		},
		{
			description: "Quantity 49 (End of 10% discount)",
			price:       100.0,
			quantity:    49,
			expected:    4410.0, // 100 * 49 * 0.9
		},
		{
			description: "Quantity 50 (25% discount starts)",
			price:       100.0,
			quantity:    50,
			expected:    3750.0, // 100 * 50 * 0.75
		},
	}

	fmt.Println("Running calculateDiscount tests...")
	for _, tc := range tests {
		runTest(tc)
	}
}