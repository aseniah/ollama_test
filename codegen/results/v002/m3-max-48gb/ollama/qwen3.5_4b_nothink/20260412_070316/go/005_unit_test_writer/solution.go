package main

import (
	"fmt"
)

// calculateDiscount calculates a bulk discount
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

func main() {
	// Helper function to run a test case
	runTest := func(price, quantity int, expected float64, description string) {
		result := calculateDiscount(float64(price), quantity)
		if result == expected {
			fmt.Printf("PASS: %s\n", description)
		} else {
			fmt.Printf("FAIL: %s (expected: %.6f, got: %.6f)\n", description, expected, result)
		}
	}

	// Test case 1: quantity < 10
	runTest(10.0, 5, 50.0, "quantity 5 with no discount")

	// Test case 2: boundary case quantity = 9
	runTest(10.0, 9, 90.0, "quantity 9 (boundary, no discount)")

	// Test case 3: boundary case quantity = 10
	runTest(10.0, 10, 90.0, "quantity 10 (boundary, 10% discount)")

	// Test case 4: quantity just below 50
	runTest(10.0, 49, 441.0, "quantity 49 (boundary, 10% discount)")

	// Test case 5: boundary case quantity = 50
	runTest(10.0, 50, 375.0, "quantity 50 (boundary, 25% discount)")
}