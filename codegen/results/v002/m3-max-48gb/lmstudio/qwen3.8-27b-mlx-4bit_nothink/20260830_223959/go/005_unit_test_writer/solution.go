package main

import (
	"fmt"
	"math"
)

// calculateDiscount is a bulk discount calculator.
// For quantity < 10: full price
// For 10 <= quantity < 50: 10% discount
// For quantity >= 50: 25% discount
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

// testCase defines a single test case
type testCase struct {
	description   string
	price         float64
	quantity      int
	expectedPrice float64
}

// runTest executes a single test case and prints the result
func runTest(tc testCase) {
	got := calculateDiscount(tc.price, tc.quantity)
	// Use a small epsilon for floating point comparison
	if math.Abs(got-tc.expectedPrice) < 1e-9 {
		fmt.Printf("PASS: %s\n", tc.description)
	} else {
		fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", tc.description, tc.expectedPrice, got)
	}
}

func main() {
	testCases := []testCase{
		{
			description:   "quantity 9 should have no discount (boundary below 10)",
			price:         10.0,
			quantity:      9,
			expectedPrice: 10.0 * 9, // 90.0
		},
		{
			description:   "quantity 10 should have 10% discount (boundary at 10)",
			price:         10.0,
			quantity:      10,
			expectedPrice: 10.0 * 10 * 0.9, // 90.0
		},
		{
			description:   "quantity 49 should have 10% discount (boundary below 50)",
			price:         10.0,
			quantity:      49,
			expectedPrice: 10.0 * 49 * 0.9, // 441.0
		},
		{
			description:   "quantity 50 should have 25% discount (boundary at 50)",
			price:         10.0,
			quantity:      50,
			expectedPrice: 10.0 * 50 * 0.75, // 375.0
		},
		{
			description:   "quantity 1 should have no discount",
			price:         5.0,
			quantity:      1,
			expectedPrice: 5.0 * 1, // 5.0
		},
	}

	for _, tc := range testCases {
		runTest(tc)
	}
}