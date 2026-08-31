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
	diff := math.Abs(a - b)
	threshold := 1e-6
	if diff > 0.1 {
		threshold = math.Abs(a) * 1e-6
	}
	return diff <= threshold
}

func test(description string, price float64, quantity int, expected float64) {
	got := calculateDiscount(price, quantity)
	if approxEqual(got, expected) {
		fmt.Printf("PASS: %s\n", description)
	} else {
		fmt.Printf("FAIL: %s (expected: %.4f, got: %.4f)\n", description, expected, got)
	}
}

func main() {
	// Test case 1: quantity 9 (below 10, no discount)
	// price = 10.0, quantity = 9, expected = 90.0
	test("quantity 9 (no discount)", 10.0, 9, 90.0)

	// Test case 2: quantity 10 (at boundary, 10% discount)
	// price = 10.0, quantity = 10, expected = 10 * 10 * 0.9 = 90.0
	test("quantity 10 (10% discount)", 10.0, 10, 90.0)

	// Test case 3: quantity 49 (just below 50, 10% discount)
	// price = 10.0, quantity = 49, expected = 10 * 49 * 0.9 = 441.0
	test("quantity 49 (10% discount)", 10.0, 49, 441.0)

	// Test case 4: quantity 50 (at boundary, 25% discount)
	// price = 10.0, quantity = 50, expected = 10 * 50 * 0.75 = 375.0
	test("quantity 50 (25% discount)", 10.0, 50, 375.0)

	// Test case 5: quantity 1 (minimum, no discount)
	// price = 5.5, quantity = 1, expected = 5.5
	test("quantity 1 (no discount)", 5.5, 1, 5.5)

	// Test case 6: quantity 100 (well above 50, 25% discount)
	// price = 20.0, quantity = 100, expected = 20 * 100 * 0.75 = 1500.0
	test("quantity 100 (25% discount)", 20.0, 100, 1500.0)
}