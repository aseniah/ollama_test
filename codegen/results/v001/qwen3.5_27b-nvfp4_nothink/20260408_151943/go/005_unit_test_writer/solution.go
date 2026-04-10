package main

import "fmt"

// calculateDiscount calculates a bulk discount based on price and quantity.
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

// testResult holds the result of a single test case.
type testResult struct {
	desc     string
	price    float64
	qty      int
	expected float64
	got      float64
	passed   bool
}

func main() {
	// Define test cases covering boundaries and standard scenarios
	testCases := []struct {
		desc     string
		price    float64
		qty      int
		expected float64
	}{
		{
			desc:     "Boundary: quantity 9 (no discount)",
			price:    100.0,
			qty:      9,
			expected: 900.0, // 100 * 9
		},
		{
			desc:     "Boundary: quantity 10 (10% discount)",
			price:    100.0,
			qty:      10,
			expected: 900.0, // 100 * 10 * 0.9
		},
		{
			desc:     "Boundary: quantity 49 (10% discount)",
			price:    100.0,
			qty:      49,
			expected: 4410.0, // 100 * 49 * 0.9
		},
		{
			desc:     "Boundary: quantity 50 (25% discount)",
			price:    100.0,
			qty:      50,
			expected: 3750.0, // 100 * 50 * 0.75
		},
	}

	passCount := 0

	for _, tc := range testCases {
		got := calculateDiscount(tc.price, tc.qty)
		passed := got == tc.expected
		
		if passed {
			fmt.Printf("PASS: %s\n", tc.desc)
			passCount++
		} else {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", tc.desc, tc.expected, got)
		}
	}

	fmt.Printf("\nTest Summary: %d/%d tests passed.\n", passCount, len(testCases))
}