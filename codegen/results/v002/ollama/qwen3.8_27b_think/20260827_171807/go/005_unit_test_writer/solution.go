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

func runTest(name string, price float64, quantity int, expected float64) {
	got := calculateDiscount(price, quantity)
	if math.Abs(got-expected) < 1e-9 {
		fmt.Printf("PASS: %s\n", name)
	} else {
		fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", name, expected, got)
	}
}

func main() {
	// Boundary: just below 10 (no discount)
	runTest("quantity 9 - no discount (price=10.0)", 10.0, 9, 90.0)

	// Boundary: exactly 10 (10% discount kicks in)
	runTest("quantity 10 - 10%% discount (price=10.0)", 10.0, 10, 90.0)

	// Boundary: just below 50 (10% discount)
	runTest("quantity 49 - 10%% discount (price=10.0)", 10.0, 49, 441.0)

	// Boundary: exactly 50 (25% discount kicks in)
	runTest("quantity 50 - 25%% discount (price=10.0)", 10.0, 50, 375.0)

	// Extra: well above 50 (25% discount)
	runTest("quantity 100 - 25%% discount (price=5.0)", 5.0, 100, 375.0)
}