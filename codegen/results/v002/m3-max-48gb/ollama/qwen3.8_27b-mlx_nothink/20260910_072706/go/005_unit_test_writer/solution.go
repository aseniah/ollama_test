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
	// Test case 1: quantity 9 (below 10, no discount)
	{
		description := "quantity 9, no discount"
		expected := 10.0 * 9
		got := calculateDiscount(10.0, 9)
		if math.Abs(got-expected) < 1e-9 {
			fmt.Printf("PASS: %s\n", description)
		} else {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", description, expected, got)
		}
	}

	// Test case 2: quantity 10 (boundary, 10% discount)
	{
		description := "quantity 10, 10%% discount"
		expected := 10.0 * 10 * 0.9
		got := calculateDiscount(10.0, 10)
		if math.Abs(got-expected) < 1e-9 {
			fmt.Printf("PASS: %s\n", description)
		} else {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", description, expected, got)
		}
	}

	// Test case 3: quantity 49 (below 50, 10% discount)
	{
		description := "quantity 49, 10%% discount"
		expected := 10.0 * 49 * 0.9
		got := calculateDiscount(10.0, 49)
		if math.Abs(got-expected) < 1e-9 {
			fmt.Printf("PASS: %s\n", description)
		} else {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", description, expected, got)
		}
	}

	// Test case 4: quantity 50 (boundary, 25% discount)
	{
		description := "quantity 50, 25%% discount"
		expected := 10.0 * 50 * 0.75
		got := calculateDiscount(10.0, 50)
		if math.Abs(got-expected) < 1e-9 {
			fmt.Printf("PASS: %s\n", description)
		} else {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", description, expected, got)
		}
	}

	// Test case 5: quantity 0 (edge case, no discount)
	{
		description := "quantity 0, no discount"
		expected := 10.0 * 0
		got := calculateDiscount(10.0, 0)
		if math.Abs(got-expected) < 1e-9 {
			fmt.Printf("PASS: %s\n", description)
		} else {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", description, expected, got)
		}
	}
}