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

func runTest(description string, price float64, quantity int, expected float64) {
	got := calculateDiscount(price, quantity)
	epsilon := 1e-9
	if math.Abs(got-expected) < epsilon {
		fmt.Printf("PASS: %s\n", description)
	} else {
		fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", description, expected, got)
	}
}

func main() {
	runTest("quantity 9 (no discount)", 10.0, 9, 90.0)
	runTest("quantity 10 (10% discount)", 10.0, 10, 90.0)
	runTest("quantity 49 (10% discount)", 10.0, 49, 441.0)
	runTest("quantity 50 (25% discount)", 10.0, 50, 375.0)
}