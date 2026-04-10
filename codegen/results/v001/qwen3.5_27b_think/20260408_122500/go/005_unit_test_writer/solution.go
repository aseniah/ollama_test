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
	// Test configuration
	const price = 10.0
	const epsilon = 1e-9 // Tolerance for float comparison

	type TestCase struct {
		quantity int
		expected float64
		desc     string
	}

	tests := []TestCase{
		{9, 90.0, "Quantity 9: No discount applied"},
		{10, 90.0, "Quantity 10: 10% discount threshold"},
		{49, 441.0, "Quantity 49: 10% discount upper bound"},
		{50, 375.0, "Quantity 50: 25% discount threshold"},
	}

	for _, t := range tests {
		got := calculateDiscount(price, t.quantity)
		
		// Use epsilon for float equality check
		if math.Abs(got-t.expected) < epsilon {
			fmt.Printf("PASS: %s\n", t.desc)
		} else {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", t.desc, t.expected, got)
		}
	}
}