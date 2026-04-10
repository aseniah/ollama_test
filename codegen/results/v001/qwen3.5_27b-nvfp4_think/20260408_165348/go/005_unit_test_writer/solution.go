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

type testCase struct {
	name       string
	price      float64
	quantity   int
	expected   float64
	description string
}

func runTest(tc testCase) {
	result := calculateDiscount(tc.price, tc.quantity)

	// Use a small epsilon for floating point comparison
	epsilon := 0.01
	if result >= tc.expected-epsilon && result <= tc.expected+epsilon {
		fmt.Printf("PASS: %s\n", tc.description)
	} else {
		fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.description, tc.expected, result)
	}
}

func main() {
	testCases := []testCase{
		{
			name:        "quantity_9",
			price:       100.0,
			quantity:    9,
			expected:    900.0,
			description: "Boundary condition: quantity 9 (no discount)",
		},
		{
			name:        "quantity_10",
			price:       100.0,
			quantity:    10,
			expected:    900.0,
			description: "Boundary condition: quantity 10 (10% discount starts)",
		},
		{
			name:        "quantity_49",
			price:       100.0,
			quantity:    49,
			expected:    4410.0,
			description: "Boundary condition: quantity 49 (still 10% discount)",
		},
		{
			name:        "quantity_50",
			price:       100.0,
			quantity:    50,
			expected:    3750.0,
			description: "Boundary condition: quantity 50 (25% discount starts)",
		},
	}

	fmt.Println("Running calculateDiscount tests...")
	fmt.Println("-----------------------------------")

	for _, tc := range testCases {
		runTest(tc)
	}

	fmt.Println("-----------------------------------")
	fmt.Println("All tests completed.")
}