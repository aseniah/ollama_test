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

func runTest(description string, price float64, quantity int, expected float64) bool {
	result := calculateDiscount(price, quantity)
	if result == expected {
		fmt.Printf("PASS: %s\n", description)
		return true
	} else {
		fmt.Printf("FAIL: %s (expected: %.1f, got: %.1f)\n", description, expected, result)
		return false
	}
}

func main() {
	passed := 0
	total := 0

	// Test case 1: quantity < 10 (no discount)
	total++
	if runTest("quantity 9, no discount applied", 10.0, 9, 90.0) {
		passed++
	}

	// Test case 2: quantity >= 10 and < 50 (10% discount) - boundary
	total++
	if runTest("quantity 10, 10%% discount applied", 10.0, 10, 90.0) {
		passed++
	}

	// Test case 3: quantity >= 10 and < 50 (10% discount) - upper boundary
	total++
	if runTest("quantity 49, 10%% discount applied", 10.0, 49, 441.0) {
		passed++
	}

	// Test case 4: quantity >= 50 (25% discount) - boundary
	total++
	if runTest("quantity 50, 25%% discount applied", 10.0, 50, 375.0) {
		passed++
	}

	fmt.Printf("\nResults: %d/%d tests passed\n", passed, total)
}