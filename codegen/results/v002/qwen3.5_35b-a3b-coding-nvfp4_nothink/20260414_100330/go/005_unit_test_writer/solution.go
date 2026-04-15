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
	// Helper to compare floats with a small tolerance
	equal := func(a, b float64) bool {
		diff := math.Abs(a - b)
		return diff < 1e-6
	}

	runTest := func(desc string, price float64, qty int, expected float64) {
		got := calculateDiscount(price, float64(qty))
		if !equal(got, expected) {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", desc, expected, got)
		} else {
			fmt.Printf("PASS: %s\n", desc)
		}
	}

	// Test cases covering boundary conditions:
	// quantity < 10 (boundary at 9)
	// 10 <= quantity < 50 (boundary at 10, 49)
	// quantity >= 50 (boundary at 50)

	// Test 1: quantity = 9 (just below first discount threshold)
	runTest("quantity=9 (no discount)", 10.0, 9, 90.0)

	// Test 2: quantity = 10 (first discount threshold: 10% off)
	runTest("quantity=10 (10% discount)", 10.0, 10, 10.0*10*0.9) // 90.0

	// Test 3: quantity = 49 (still in 10% discount range)
	runTest("quantity=49 (10% discount)", 10.0, 49, 10.0*49*0.9) // 441.0

	// Test 4: quantity = 50 (second discount threshold: 25% off)
	runTest("quantity=50 (25% discount)", 10.0, 50, 10.0*50*0.75) // 375.0
}