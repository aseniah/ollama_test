package main

import (
	"fmt"
)

// calculateDiscount calculates a bulk discount based on quantity.
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

func testDiscount(description string, expected, price, quantity int) {
	result := calculateDiscount(float64(price), quantity)
	
	if result == float64(expected) {
		fmt.Printf("PASS: %s\n", description)
	} else {
		fmt.Printf("FAIL: %s (expected: %d, got: %.2f)\n", description, expected, result)
	}
}

func main() {
	// Test case 1: quantity = 9 (less than 10, no discount)
	testDiscount("Quantity 9 (no discount)", 90, 10, 9)

	// Test case 2: quantity = 10 (10 to 49, 10% discount)
	testDiscount("Quantity 10 (10% discount)", 81, 10, 10)

	// Test case 3: quantity = 49 (10 to 49, 10% discount)
	testDiscount("Quantity 49 (10% discount)", 393.30, 10, 49)

	// Test case 4: quantity = 50 (50+, 25% discount)
	testDiscount("Quantity 50 (25% discount)", 275.00, 10, 50)
}