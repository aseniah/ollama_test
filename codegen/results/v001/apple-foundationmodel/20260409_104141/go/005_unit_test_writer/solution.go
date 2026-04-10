package main

import (
	"fmt"
)

// Function to calculate discount
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

// Test case structure
type TestCase struct {
	description string
	price       float64
	quantity    int
	expected    float64
}

// Function to run tests
func runTests(cases []TestCase) {
	for _, test := range cases {
		result := calculateDiscount(test.price, test.quantity)
		if result == test.expected {
			fmt.Printf("PASS: %s\n", test.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", test.description, test.expected, result)
		}
	}
}

func main() {
	testCases := []TestCase{
		{description: "Quantity 9", price: 10.0, quantity: 9, expected: 90.0},
		{description: "Quantity 10", price: 20.0, quantity: 10, expected: 200.0},
		{description: "Quantity 49", price: 15.0, quantity: 49, expected: 110.25},
		{description: "Quantity 50", price: 25.0, quantity: 50, expected: 1250.0},
	}

	runTests(testCases)
}