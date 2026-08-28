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

func main() {
	// Test case 1: quantity 9 (boundary, below 10)
	price := 100.0
	quantity := 9
	expected := price * float64(quantity)
	got := calculateDiscount(price, quantity)
	if abs(expected-got) < 1e-9 {
		fmt.Println("PASS: quantity 9 (below 10)")
	} else {
		fmt.Printf("FAIL: quantity 9 (below 10) (expected: %v, got: %v)\n", expected, got)
	}

	// Test case 2: quantity 10 (boundary, starts 10% discount)
	price = 100.0
	quantity = 10
	expected = price * float64(quantity) * 0.9
	got = calculateDiscount(price, quantity)
	if abs(expected-got) < 1e-9 {
		fmt.Println("PASS: quantity 10 (starts 10% discount)")
	} else {
		fmt.Printf("FAIL: quantity 10 (starts 10% discount) (expected: %v, got: %v)\n", expected, got)
	}

	// Test case 3: quantity 49 (boundary, still 10% discount)
	price = 100.0
	quantity = 49
	expected = price * float64(quantity) * 0.9
	got = calculateDiscount(price, quantity)
	if abs(expected-got) < 1e-9 {
		fmt.Println("PASS: quantity 49 (still 10% discount)")
	} else {
		fmt.Printf("FAIL: quantity 49 (still 10% discount) (expected: %v, got: %v)\n", expected, got)
	}

	// Test case 4: quantity 50 (boundary, starts 25% discount)
	price = 100.0
	quantity = 50
	expected = price * float64(quantity) * 0.75
	got = calculateDiscount(price, quantity)
	if abs(expected-got) < 1e-9 {
		fmt.Println("PASS: quantity 50 (starts 25% discount)")
	} else {
		fmt.Printf("FAIL: quantity 50 (starts 25% discount) (expected: %v, got: %v)\n", expected, got)
	}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}