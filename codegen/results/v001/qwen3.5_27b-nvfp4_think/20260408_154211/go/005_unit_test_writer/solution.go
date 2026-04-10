package main

import "fmt"

// calculateDiscount calculates the bulk discount based on quantity
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

// testCase holds a test case for calculateDiscount
type testCase struct {
	description string
	price       float64
	quantity    int
	expected    float64
}

// runTest runs a single test and prints the result
func runTest(tc testCase) bool {
	got := calculateDiscount(tc.price, tc.quantity)

	// Use epsilon for float comparison to handle floating point precision
	const epsilon = 0.01
	diff := got - tc.expected
	if diff < 0 {
		diff = -diff
	}

	if diff < epsilon {
		fmt.Printf("PASS: %s\n", tc.description)
		return true
	} else {
		fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", 
			tc.description, tc.expected, got)
		return false
	}
}

func main() {
	testCases := []testCase{
		{
			description: "Quantity 9 - no discount (< 10)",
			price:       100.00,
			quantity:    9,
			expected:    900.00, // 100 * 9 = 900 (no discount)
		},
		{
			description: "Quantity 10 - 10% discount (>= 10)",
			price:       100.00,
			quantity:    10,
			expected:    900.00, // 100 * 10 * 0.9 = 900
		},
		{
			description: "Quantity 49 - 10% discount (< 50)",
			price:       100.00,
			quantity:    49,
			expected:    4410.00, // 100 * 49 * 0.9 = 4410
		},
		{
			description: "Quantity 50 - 25% discount (>= 50)",
			price:       100.00,
			quantity:    50,
			expected:    3750.00, // 100 * 50 * 0.75 = 3750
		},
		{
			description: "Quantity 0 - no discount",
			price:       50.00,
			quantity:    0,
			expected:    0.00, // 50 * 0 = 0 (no discount)
		},
	}

	fmt.Println("Running calculateDiscount tests...")
	fmt.Println()

	passCount := 0
	failCount := 0

	for _, tc := range testCases {
		if runTest(tc) {
			passCount++
		} else {
			failCount++
		}
	}

	fmt.Println()
	fmt.Printf("Results: %d passed, %d failed\n", passCount, failCount)
}