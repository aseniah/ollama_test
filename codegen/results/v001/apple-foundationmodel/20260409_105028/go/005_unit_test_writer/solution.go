package main

import (
	"fmt"
)

func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

func testCalculateDiscount() {
	// Test case 1: Quantity 9
	price := 100.0
	quantity := 9
	expected := price * float64(quantity)
	result := calculateDiscount(price, quantity)
	if result == expected {
		fmt.Printf("PASS: Quantity 9\n")
	} else {
		fmt.Printf("FAIL: Quantity 9\nExpected: %f, got: %f\n", expected, result)
	}

	// Test case 2: Quantity 10
	price := 100.0
	quantity := 10
	expected := price * float64(quantity)
	result := calculateDiscount(price, quantity)
	if result == expected {
		fmt.Printf("PASS: Quantity 10\n")
	} else {
		fmt.Printf("FAIL: Quantity 10\nExpected: %f, got: %f\n", expected, result)
	}

	// Test case 3: Quantity 49
	price := 100.0
	quantity := 49
	expected := price * float64(quantity) * 0.9
	result := calculateDiscount(price, quantity)
	if result == expected {
		fmt.Printf("PASS: Quantity 49\n")
	} else {
		fmt.Printf("FAIL: Quantity 49\nExpected: %f, got: %f\n", expected, result)
	}

	// Test case 4: Quantity 50
	price := 100.0
	quantity := 50
	expected := price * float64(quantity) * 0.9
	result := calculateDiscount(price, quantity)
	if result == expected {
		fmt.Printf("PASS: Quantity 50\n")
	} else {
		fmt.Printf("FAIL: Quantity 50\nExpected: %f, got: %f\n", expected, result)
	}
}

func main() {
	testCalculateDiscount()
}