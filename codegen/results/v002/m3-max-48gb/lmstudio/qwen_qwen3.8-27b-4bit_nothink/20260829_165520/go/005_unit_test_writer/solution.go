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

func runTest(description string, price float64, quantity int, expected float64) bool {
	got := calculateDiscount(price, quantity)
	// Use a small epsilon for floating point comparison
	const epsilon = 1e-9
	if math.Abs(got-expected) < epsilon {
		fmt.Printf("PASS: %s\n", description)
		return true
	}
	fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", description, expected, got)
	return false
}

func main() {
	allPass := true

	// Test 1: quantity = 9 (just below 10, no discount)
	price := 10.0
	qty := 9
	expected := price * float64(qty) // 90.0
	allPass = runTest("quantity 9 (no discount)", price, qty, expected) && allPass

	// Test 2: quantity = 10 (boundary, 10% discount)
	price = 10.0
	qty = 10
	expected = price * float64(qty) * 0.9 // 90.0
	allPass = runTest("quantity 10 (10% discount)", price, qty, expected) && allPass

	// Test 3: quantity = 49 (just below 50, 10% discount)
	price = 10.0
	qty = 49
	expected = price * float64(qty) * 0.9 // 441.0
	allPass = runTest("quantity 49 (10% discount)", price, qty, expected) && allPass

	// Test 4: quantity = 50 (boundary, 25% discount)
	price = 10.0
	qty = 50
	expected = price * float64(qty) * 0.75 // 375.0
	allPass = runTest("quantity 50 (25% discount)", price, qty, expected) && allPass

	// Test 5: quantity = 0 (no discount, edge case)
	price = 5.0
	qty = 0
	expected = price * float64(qty) // 0.0
	allPass = runTest("quantity 0 (no discount)", price, qty, expected) && allPass

	// Test 6: different price to verify calculation
	price = 25.5
	qty = 100
	expected = price * float64(qty) * 0.75 // 1912.5
	allPass = runTest("quantity 100 with price 25.5 (25% discount)", price, qty, expected) && allPass

	if allPass {
		fmt.Println("All tests passed!")
	} else {
		fmt.Println("Some tests failed.")
	}
}