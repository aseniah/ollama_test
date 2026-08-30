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
	testCases := []struct {
		desc     string
		price    float64
		quantity int
		expected float64
	}{
		{
			desc:     "Boundary: quantity 9 (no discount)",
			price:    100.0,
			quantity: 9,
			expected: 900.0, // 100 * 9 = 900
		},
		{
			desc:     "Boundary: quantity 10 (10% discount)",
			price:    100.0,
			quantity: 10,
			expected: 900.0, // 100 * 10 * 0.9 = 900
		},
		{
			desc:     "Boundary: quantity 49 (10% discount)",
			price:    100.0,
			quantity: 49,
			expected: 4410.0, // 100 * 49 * 0.9 = 4410
		},
		{
			desc:     "Boundary: quantity 50 (25% discount)",
			price:    100.0,
			quantity: 50,
			expected: 3750.0, // 100 * 50 * 0.75 = 3750
		},
		{
			desc:     "Standard case: quantity 60 (25% discount)",
			price:    50.0,
			quantity: 60,
			expected: 2250.0, // 50 * 60 * 0.75 = 2250
		},
	}

	for _, tc := range testCases {
		got := calculateDiscount(tc.price, tc.quantity)
		
		// Use a small epsilon for float comparison to avoid precision issues
		epsilon := 0.01
		if got < tc.expected-epsilon || got > tc.expected+epsilon {
			fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.desc, tc.expected, got)
		} else {
			fmt.Printf("PASS: %s\n", tc.desc)
		}
	}
}