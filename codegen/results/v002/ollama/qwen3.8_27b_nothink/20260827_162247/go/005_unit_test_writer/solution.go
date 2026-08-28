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

func main() {
	// Test case 1: quantity = 9 (boundary: just below 10, no discount)
	desc1 := "quantity 9 (no discount)"
	price1 := 100.0
	expected1 := price1 * 9 // 900
	got1 := calculateDiscount(price1, 9)
	if approxEqual(got1, expected1) {
		fmt.Printf("PASS: %s\n", desc1)
	} else {
		fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", desc1, expected1, got1)
	}

	// Test case 2: quantity = 10 (boundary: exactly 10, 10% discount)
	desc2 := "quantity 10 (10% discount)"
	price2 := 100.0
	expected2 := price2 * 10 * 0.9 // 900
	got2 := calculateDiscount(price2, 10)
	if approxEqual(got2, expected2) {
		fmt.Printf("PASS: %s\n", desc2)
	} else {
		fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", desc2, expected2, got2)
	}

	// Test case 3: quantity = 49 (boundary: just below 50, 10% discount)
	desc3 := "quantity 49 (10% discount)"
	price3 := 100.0
	expected3 := price3 * 49 * 0.9 // 4410
	got3 := calculateDiscount(price3, 49)
	if approxEqual(got3, expected3) {
		fmt.Printf("PASS: %s\n", desc3)
	} else {
		fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", desc3, expected3, got3)
	}

	// Test case 4: quantity = 50 (boundary: exactly 50, 25% discount)
	desc4 := "quantity 50 (25% discount)"
	price4 := 100.0
	expected4 := price4 * 50 * 0.75 // 3750
	got4 := calculateDiscount(price4, 50)
	if approxEqual(got4, expected4) {
		fmt.Printf("PASS: %s\n", desc4)
	} else {
		fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", desc4, expected4, got4)
	}

	// Test case 5: quantity = 1 (minimum, no discount)
	desc5 := "quantity 1 (no discount)"
	price5 := 50.0
	expected5 := price5 * 1 // 50
	got5 := calculateDiscount(price5, 1)
	if approxEqual(got5, expected5) {
		fmt.Printf("PASS: %s\n", desc5)
	} else {
		fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", desc5, expected5, got5)
	}

	// Test case 6: quantity = 100 (well above 50, 25% discount)
	desc6 := "quantity 100 (25% discount)"
	price6 := 10.0
	expected6 := price6 * 100 * 0.75 // 750
	got6 := calculateDiscount(price6, 100)
	if approxEqual(got6, expected6) {
		fmt.Printf("PASS: %s\n", desc6)
	} else {
		fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", desc6, expected6, got6)
	}
}