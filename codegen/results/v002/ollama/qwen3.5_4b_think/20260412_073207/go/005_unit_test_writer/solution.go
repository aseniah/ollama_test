package main

import (
	"fmt"
	"math"
)

// calculateDiscount calculates a bulk discount based on quantity
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

// test case structure to hold test parameters
type testCase struct {
	description string
	price       float64
	quantity    int
	expected    float64
}

// runTests runs a suite of test cases and prints results
func runTests() {
	testCases := []testCase{
		{
			description: "quantity 9 (no discount)",
			price:       10.0,
			quantity:    9,
			expected:    90.0,
		},
		{
			description: "quantity 10 (10% discount)",
			price:       10.0,
			quantity:    10,
			expected:    90.0,
		},
		{
			description: "quantity 49 (10% discount)",
			price:       10.0,
			quantity:    49,
			expected:    441.0,
		},
		{
			description: "quantity 50 (25% discount)",
			price:       10.0,
			quantity:    50,
			expected:    375.0,
		},
	}

	for _, tc := range testCases {
		result := calculateDiscount(tc.price, tc.quantity)
		// Use relative error for floating point comparison
		isPass := math.Abs(result-tc.expected) < 0.0001
		fmt.Printf("%s\n", tc.description)
		if isPass {
			fmt.Printf("PASS\n")
		} else {
			fmt.Printf("FAIL: expected: %.1f, got: %.1f\n", tc.expected, result)
		}
	}
}

func main() {
	runTests()
}