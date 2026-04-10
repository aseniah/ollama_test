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

func approxEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}

func check(description string, expected, got float64) {
	if approxEqual(expected, got) {
		fmt.Printf("PASS: %s\n", description)
	} else {
		fmt.Printf("FAIL: %s (expected: %g, got: %g)\n", description, expected, got)
	}
}

func main() {
	price := 10.0

	// quantity 9: no discount, 9 * 10 = 90
	check("quantity 9 (no discount)", 90.0, calculateDiscount(price, 9))

	// quantity 10: 10% discount, 10 * 10 * 0.9 = 90
	check("quantity 10 (10% discount boundary)", 90.0, calculateDiscount(price, 10))

	// quantity 49: 10% discount, 49 * 10 * 0.9 = 441
	check("quantity 49 (10% discount upper boundary)", 441.0, calculateDiscount(price, 49))

	// quantity 50: 25% discount, 50 * 10 * 0.75 = 375
	check("quantity 50 (25% discount boundary)", 375.0, calculateDiscount(price, 50))

	// quantity 1: no discount, 1 * 10 = 10
	check("quantity 1 (single unit)", 10.0, calculateDiscount(price, 1))

	// quantity 100: 25% discount, 100 * 10 * 0.75 = 750
	check("quantity 100 (bulk discount)", 750.0, calculateDiscount(price, 100))
}
