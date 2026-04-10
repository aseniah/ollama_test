package main

import (
	"fmt"
	"math"
)

func main() {
	price := 100.0

	// Test case 1: quantity = 9
	desc := "quantity = 9"
	expected := price * 9.0
	got := calculateDiscount(price, 9)
	test(desc, expected, got)

	// Test case 2: quantity = 10
	desc = "quantity = 10"
	expected = price * 10.0 * 0.9
	got = calculateDiscount(price, 10)
	test(desc, expected, got)

	// Test case 3: quantity = 49
	desc = "quantity = 49"
	expected = price * 49.0 * 0.9
	got = calculateDiscount(price, 49)
	test(desc, expected, got)

	// Test case 4: quantity = 50
	desc = "quantity = 50"
	expected = price * 50.0 * 0.75
	got = calculateDiscount(price, 50)
	test(desc, expected, got)
}

func test(desc string, expected float64, got float64) {
	// Use a small epsilon for floating point comparison
	if math.Abs(got - expected) < 0.00001 {
		fmt.Println("PASS:", desc)
	} else {
		fmt.Println("FAIL:", desc, "(expected:", expected, ", got:", got, ")")
	}
}

func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}