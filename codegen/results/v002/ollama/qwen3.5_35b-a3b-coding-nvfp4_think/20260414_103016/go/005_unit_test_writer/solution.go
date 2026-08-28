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

func assertApproxEqual(expected, got, tolerance float64, description string) {
	if abs(expected-got) <= tolerance {
		fmt.Printf("PASS: %s\n", description)
	} else {
		fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", description, expected, got)
	}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	testCases := []struct {
		price       float64
		quantity    int
		expected    float64
		description string
	}{
		{price: 100.00, quantity: 9, expected: 900.00, description: "No discount (quantity < 10)"},
		{price: 100.00, quantity: 10, expected: 900.00, description: "10% discount (quantity = 10)"},
		{price: 100.00, quantity: 49, expected: 4410.00, description: "10% discount (quantity = 49)"},
		{price: 100.00, quantity: 50, expected: 3750.00, description: "25% discount (quantity = 50)"},
	}

	for _, tc := range testCases {
		result := calculateDiscount(tc.price, tc.quantity)
		assertApproxEqual(tc.expected, result, 0.01, tc.description)
	}
}