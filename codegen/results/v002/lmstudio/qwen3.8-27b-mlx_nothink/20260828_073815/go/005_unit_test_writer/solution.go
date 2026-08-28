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

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	type testCase struct {
		description string
		price       float64
		quantity    int
		expected    float64
	}

	cases := []testCase{
		{"quantity 9", 10.0, 9, 90.0},
		{"quantity 10", 10.0, 10, 90.0},
		{"quantity 49", 10.0, 49, 441.0},
		{"quantity 50", 10.0, 50, 375.0},
	}

	const epsilon = 1e-6

	for _, tc := range cases {
		got := calculateDiscount(tc.price, tc.quantity)

		if abs(got-tc.expected) <= epsilon {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", tc.description, tc.expected, got)
		}
	}
}