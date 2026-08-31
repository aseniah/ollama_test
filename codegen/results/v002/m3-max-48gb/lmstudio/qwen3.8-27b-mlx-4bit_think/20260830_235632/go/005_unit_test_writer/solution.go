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
		description string
		price       float64
		quantity    int
		expected    float64
	}

	tests := []testCase{
		{
			description: "quantity 9 (below 10, no discount)",
			price:       10,
			quantity:    9,
			expected:    90,
		},
		{
			description: "quantity 10 (boundary, 10% discount)",
			price:       10,
			quantity:    10,
			expected:    90,
		},
		{
			description: "quantity 49 (below 50, 10% discount)",
			price:       10,
			quantity:    49,
			expected:    441,
		},
		{
			description: "quantity 50 (boundary, 25% discount)",
			price:       10,
			quantity:    50,
			expected:    375,
		},
	}

	for _, tc := range tests {
		got := calculateDiscount(tc.price, tc.quantity)
		if got == tc.expected {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.description, tc.expected, got)
		}
	}
}