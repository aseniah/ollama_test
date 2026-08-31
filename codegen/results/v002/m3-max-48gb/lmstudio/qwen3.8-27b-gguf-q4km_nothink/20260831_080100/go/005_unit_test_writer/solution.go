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

func approxEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}

func main() {
	type testCase struct {
		description string
		price       float64
		quantity    int
		expected    float64
	}

	testCases := []testCase{
		{"quantity 9 (below first tier)", 10.0, 9, 90.0},
		{"quantity 10 (first tier boundary)", 10.0, 10, 90.0},
		{"quantity 49 (second tier boundary)", 10.0, 49, 441.0},
		{"quantity 50 (third tier boundary)", 10.0, 50, 375.0},
		{"quantity 1 (single item)", 5.0, 1, 5.0},
		{"quantity 25 (middle of second tier)", 20.0, 25, 450.0},
		{"quantity 100 (third tier)", 10.0, 100, 750.0},
	}

	allPassed := true
	for _, tc := range testCases {
		got := calculateDiscount(tc.price, tc.quantity)
		if approxEqual(got, tc.expected) {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", tc.description, tc.expected, got)
			allPassed = false
		}
	}

	if allPassed {
		fmt.Println("All tests passed!")
	}
}