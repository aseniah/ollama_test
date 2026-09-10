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
	// Test cases with boundary conditions
	// quantity 9: no discount, price * 9
	// quantity 10: 10% discount, price * 10 * 0.9
	// quantity 49: 10% discount, price * 49 * 0.9
	// quantity 50: 25% discount, price * 50 * 0.75

	price := 10.0

	// Test case 1: quantity 9 (boundary, no discount)
	expected1 := price * 9.0
	got1 := calculateDiscount(price, 9)
	if math.Abs(got1-expected1) < 1e-9 {
		fmt.Printf("PASS: quantity 9, no discount\n")
	} else {
		fmt.Printf("FAIL: quantity 9, no discount (expected: %f, got: %f)\n", expected1, got1)
	}

	// Test case 2: quantity 10 (boundary, 10% discount)
	expected2 := price * 10.0 * 0.9
	got2 := calculateDiscount(price, 10)
	if math.Abs(got2-expected2) < 1e-9 {
		fmt.Printf("PASS: quantity 10, 10%% discount\n")
	} else {
		fmt.Printf("FAIL: quantity 10, 10%% discount (expected: %f, got: %f)\n", expected2, got2)
	}

	// Test case 3: quantity 49 (boundary, 10% discount)
	expected3 := price * 49.0 * 0.9
	got3 := calculateDiscount(price, 49)
	if math.Abs(got3-expected3) < 1e-9 {
		fmt.Printf("PASS: quantity 49, 10%% discount\n")
	} else {
		fmt.Printf("FAIL: quantity 49, 10%% discount (expected: %f, got: %f)\n", expected3, got3)
	}

	// Test case 4: quantity 50 (boundary, 25% discount)
	expected4 := price * 50.0 * 0.75
	got4 := calculateDiscount(price, 50)
	if math.Abs(got4-expected4) < 1e-9 {
		fmt.Printf("PASS: quantity 50, 25%% discount\n")
	} else {
		fmt.Printf("FAIL: quantity 50, 25%% discount (expected: %f, got: %f)\n", expected4, got4)
	}
}