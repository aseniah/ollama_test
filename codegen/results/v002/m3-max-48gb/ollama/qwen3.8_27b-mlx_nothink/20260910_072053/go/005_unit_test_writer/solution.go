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
	// Test case 1: quantity = 9 (below first threshold, no discount)
	// Expected: 10.0 * 9 = 90.0
	expected1 := 90.0
	got1 := calculateDiscount(10.0, 9)
	if approxEqual(got1, expected1) {
		fmt.Println("PASS: quantity 9 (no discount)")
	} else {
		fmt.Printf("FAIL: quantity 9 (no discount) (expected: %f, got: %f)\n", expected1, got1)
	}

	// Test case 2: quantity = 10 (at boundary, 10% discount)
	// Expected: 10.0 * 10 * 0.9 = 90.0
	expected2 := 90.0
	got2 := calculateDiscount(10.0, 10)
	if approxEqual(got2, expected2) {
		fmt.Println("PASS: quantity 10 (10%% discount boundary)")
	} else {
		fmt.Printf("FAIL: quantity 10 (10%% discount boundary) (expected: %f, got: %f)\n", expected2, got2)
	}

	// Test case 3: quantity = 49 (below 50 threshold, 10% discount)
	// Expected: 20.0 * 49 * 0.9 = 882.0
	expected3 := 882.0
	got3 := calculateDiscount(20.0, 49)
	if approxEqual(got3, expected3) {
		fmt.Println("PASS: quantity 49 (10%% discount)")
	} else {
		fmt.Printf("FAIL: quantity 49 (10%% discount) (expected: %f, got: %f)\n", expected3, got3)
	}

	// Test case 4: quantity = 50 (at boundary, 25% discount)
	// Expected: 20.0 * 50 * 0.75 = 750.0
	expected4 := 750.0
	got4 := calculateDiscount(20.0, 50)
	if approxEqual(got4, expected4) {
		fmt.Println("PASS: quantity 50 (25%% discount boundary)")
	} else {
		fmt.Printf("FAIL: quantity 50 (25%% discount boundary) (expected: %f, got: %f)\n", expected4, got4)
	}
}