package main

import "fmt"

func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

func main() {
	type testCase struct {
		price       float64
		quantity    int
		expected    float64
		description string
	}

	testCases := []testCase{
		{100.0, 9, 900.0, "No discount for quantity 9"},
		{100.0, 10, 900.0, "10% discount for quantity 10"},
		{100.0, 49, 4410.0, "10% discount for quantity 49"},
		{100.0, 50, 3750.0, "25% discount for quantity 50"},
	}

	for _, tc := range testCases {
		got := calculateDiscount(tc.price, tc.quantity)
		if got == tc.expected {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", tc.description, tc.expected, got)
		}
	}
}