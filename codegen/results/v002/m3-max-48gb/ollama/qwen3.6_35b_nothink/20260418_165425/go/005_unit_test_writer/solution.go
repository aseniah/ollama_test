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

func main() {
	testCases := []struct {
		description string
		price       float64
		quantity    int
		expected    float64
	}{
		{"Bulk discount for quantity 9 (less than 10)", 100.0, 9, 900.0},
		{"Bulk discount for quantity 10 (first threshold)", 100.0, 10, 900.0},
		{"Bulk discount for quantity 49 (just before second threshold)", 100.0, 49, 4410.0},
		{"Bulk discount for quantity 50 (second threshold)", 100.0, 50, 3750.0},
	}

	allPassed := true
	for _, tc := range testCases {
		result := calculateDiscount(tc.price, tc.quantity)
		if math.Abs(result-tc.expected) < 1e-9 {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %.4f, got: %.4f)\n", tc.description, tc.expected, result)
			allPassed = false
		}
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	} else {
		fmt.Println("\nSome tests failed.")
	}
}