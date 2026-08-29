```go
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

func almostEqual(a, b float64) bool {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff < 1e-6
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
			description: "quantity 9: no discount",
			price:       10,
			quantity:    9,
			expected:    90,
		},
		{
			description: "quantity 10: 10% discount boundary",
			price:       10,
			quantity:    10,
			expected:    90,
		},
		{
			description: "quantity 49: 10% discount upper boundary",
			price:       10,
			quantity:    49,
			expected:    441,
		},
		{
			description: "quantity 50: 25% discount boundary",
			price:       10,
			quantity:    50,
			expected:    375,
		},
	}

	for _, tc := range tests {
		got := calculateDiscount(tc.price, tc.quantity)

		if almostEqual(tc.expected, got) {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("