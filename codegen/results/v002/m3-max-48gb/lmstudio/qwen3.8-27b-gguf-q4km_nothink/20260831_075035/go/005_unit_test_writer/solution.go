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

func floatEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}

func main() {
	// Test case 1: quantity 9 (boundary below 10)
	expected1 := 10.0 * 9 // 90.0
	got1 := calculateDiscount(10.0, 9)
	if floatEqual(got1, expected1) {
		fmt.Println("PASS: quantity 9 (no discount)")
	} else {
		fmt.Printf("FAIL: quantity 9 (no discount) (expected: %v, got: %v)\n", expected1, got1)
	}

	// Test case 2: quantity 10 (boundary at 10, 10% discount)
	expected2 := 10.0 * 10 * 0.9 // 90.0
	got2 := calculateDiscount(10.0, 10)
	if floatEqual(got2, expected2) {
		fmt.Println("PASS: quantity 10 (10% discount)")
	} else {
		fmt.Printf("FAIL: quantity 10 (10% discount) (expected: %v, got: %v)\n", expected2, got2)
	}

	// Test case 3: quantity 49 (boundary below 50, 10% discount)
	expected3 := 10.0 * 49 * 0.9 // 441.0
	got3 := calculateDiscount(10.0, 49)
	if floatEqual(got3, expected3) {
		fmt.Println("PASS: quantity 49 (10% discount)")
	} else {
		fmt.Printf("FAIL: quantity 49 (10% discount) (expected: %v, got: %v)\n", expected3, got3)
	}

	// Test case 4: quantity 50 (boundary at 50, 25% discount)
	expected4 := 10.0 * 50 * 0.75 // 375.0
	got4 := calculateDiscount(10.0, 50)
	if floatEqual(got4, expected4) {
		fmt.Println("PASS: quantity 50 (25% discount)")
	} else {
		fmt.Printf("FAIL: quantity 50 (25% discount) (expected: %v, got: %v)\n", expected4, got4)
	}

	// Additional test case: quantity 100 (well above 50, 25% discount)
	expected5 := 10.0 * 100 * 0.75 // 750.0
	got5 := calculateDiscount(10.0, 100)
	if floatEqual(got5, expected5) {
		fmt.Println("PASS: quantity 100 (25% discount)")
	} else {
		fmt.Printf("FAIL: quantity 100 (25% discount) (expected: %v, got: %v)\n", expected5, got5)
	}

	// Additional test case: quantity 1 (single item, no discount)
	expected6 := 10.0 * 1 // 10.0
	got6 := calculateDiscount(10.0, 1)
	if floatEqual(got6, expected6) {
		fmt.Println("PASS: quantity 1 (no discount)")
	} else {
		fmt.Printf("FAIL: quantity 1 (no discount) (expected: %v, got: %v)\n", expected6, got6)
	}
}