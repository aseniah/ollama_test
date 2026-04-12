package main

import (
	"fmt"
	"math"
)

func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

type testCase struct {
	description string
	price       float64
	quantity    int
	expected    float64
}

func runTest(tc testCase) {
	result := calculateDiscount(tc.price, tc.quantity)
	// Allow for small floating point differences
	if math.Abs(result-tc.expected) < 0.001 {
		fmt.Printf("PASS: %s\n", tc.description)
	} else {
		fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.description, tc.expected, result)
	}
}

func main() {
	tests := []testCase{
		{
			description: "quantity 9 - no discount (below 10)",
			price:       100.0,
			quantity:    9,
			expected:    900.0, // 100 * 9 = 900
		},
		{
			description: "quantity 10 - 10%% discount (boundary)",
			price:       100.0,
			quantity:    10,
			expected:    900.0, // 100 * 10 * 0.9 = 900
		},
		{
			description: "quantity 49 - 10%% discount (upper boundary)",
			price:       100.0,
			quantity:    49,
			expected:    4410.0, // 100 * 49 * 0.9 = 4410
		},
		{
			description: "quantity 50 - 25%% discount (boundary)",
			price:       100.0,
			quantity:    50,
			expected:    3750.0, // 100 * 50 * 0.75 = 3750
		},
		{
			description: "quantity 100 - 25%% discount (above 50)",
			price:       50.0,
			quantity:    100,
			expected:    3750.0, // 50 * 100 * 0.75 = 3750
		},
	}

	fmt.Println("Running calculateDiscount tests...")
	fmt.Println("-----------------------------------")

	for _, tc := range tests {
		runTest(tc)
	}

	fmt.Println("-----------------------------------")
	fmt.Println("Tests completed.")
}