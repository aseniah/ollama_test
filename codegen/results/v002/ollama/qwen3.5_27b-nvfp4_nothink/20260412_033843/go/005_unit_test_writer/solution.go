package main

import "fmt"

func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

func runTest(description string, price float64, quantity int, expected float64) {
	result := calculateDiscount(price, quantity)
	if result == expected {
		fmt.Printf("PASS: %s\n", description)
	} else {
		fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", description, expected, result)
	}
}

func main() {
	// Test Case 1: Boundary - quantity < 10 (no discount)
	runTest("Quantity 9 (no discount)", 10.0, 9, 90.0)

	// Test Case 2: Boundary - quantity >= 10 and < 50 (10% discount)
	runTest("Quantity 10 (10%% discount)", 10.0, 10, 90.0)

	// Test Case 3: Boundary - quantity just below 50 (10% discount)
	runTest("Quantity 49 (10%% discount)", 10.0, 49, 441.0) // 10 * 49 * 0.9 = 441

	// Test Case 4: Boundary - quantity >= 50 (25% discount)
	runTest("Quantity 50 (25%% discount)", 10.0, 50, 375.0) // 10 * 50 * 0.75 = 375

	// Additional Test Case: Non-boundary value in the middle range
	runTest("Quantity 25 (10%% discount)", 20.0, 25, 450.0) // 20 * 25 * 0.9 = 450
}