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
	d := a - b
	if d < 0 {
		d = -d
	}
	return d < 1e-9
}

type testCase struct {
	description string
	price       float64
	quantity    int
	expected    float64
}

func main() {
	tests := []testCase{
		{"quantity below 10 (no discount)", 100.0, 9, 900.0},
		{"quantity exactly 10 (10% discount)", 100.0, 10, 900.0},
		{"quantity 49 (10% discount)", 100.0, 49, 4410.0},
		{"quantity exactly 50 (25% discount)", 100.0, 50, 3750.0},
	}

	for _, tc := range tests {
		got := calculateDiscount(tc.price, tc.quantity)
		if almostEqual(got, tc.expected) {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %g, got: %g)\n", tc.description, tc.expected, got)
		}
	}
}