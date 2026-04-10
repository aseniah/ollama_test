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

func testCase(description string, price float64, quantity int, expected float64) {
	result := calculateDiscount(price, quantity)
	// Use a small epsilon for floating point comparison
	if math.Abs(result-expected) < 0.0001 {
		fmt.Printf("PASS: %s\n", description)
	} else {
		fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", description, expected, result)
	}
}

func main() {
	// Test case 1: quantity < 10 (boundary at 9)
	testCase("quantity 9 with price 10 - no discount", 10.0, 9, 90.0)

	// Test case 2: quantity >= 10 and < 50 (boundary at 10)
	testCase("quantity 10 with price 10 - 10% discount", 10.0, 10, 90.0)

	// Test case 3: quantity >= 50 (boundary at 49)
	testCase("quantity 49 with price 10 - 10% discount", 10.0, 49, 441.0)

	// Test case 4: quantity >= 50 (boundary at 50)
	testCase("quantity 50 with price 10 - 25% discount", 10.0, 50, 375.0)

	// Additional test cases for more coverage
	testCase("quantity 1 with price 5 - no discount", 5.0, 1, 5.0)

	testCase("quantity 25 with price 20 - 10% discount", 20.0, 25, 450.0)

	testCase("quantity 100 with price 15 - 25% discount", 15.0, 100, 1125.0)
}
