package main

import (
	"fmt"
)

// calculateDiscount calculates the discount based on the quantity of items.
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

// check function to test calculateDiscount with expected and actual values.
func check(description string, expected float64, got float64) {
	if expected == got {
		fmt.Printf("PASS: %s\n", description)
	} else {
		fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", description, expected, got)
	}
}

func main() {
	// Test cases
	check("quantity 9", 10*5.0, calculateDiscount(5.0, 9))
	check("quantity 10", 10*4.5, calculateDiscount(5.0, 10))
	check("quantity 49", 49*3.6, calculateDiscount(5.0, 49))
	check("quantity 50", 50*3.75, calculateDiscount(5.0, 50))

	// Additional test cases
	check("quantity 1", 1*5.0, calculateDiscount(5.0, 1))
	check("quantity 51", 51*3.75, calculateDiscount(5.0, 51))
}