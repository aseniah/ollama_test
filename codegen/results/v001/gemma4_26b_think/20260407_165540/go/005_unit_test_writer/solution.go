package main

import "fmt"

// calculateDiscount calculates a bulk discount based on quantity.
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

// testCase defines the structure for our unit tests.
type testCase struct {
	description string
	price       float64
	quantity    int
	expected    float64
}

func main() {
	// Define test cases covering boundary conditions: 9, 10, 49, 50.
	tests := []testCase{
		{
			description: "Quantity 9 (No discount)",
			price:       100.0,
			quantity:    9,
			expected:    900.0,
		},
		{
			description: "Quantity 10 (Start of 10% discount tier)",
			price:       100.0,
			quantity:    10,
			expected:    900.0, // 100 * 10 * 0.9
		},
		{
			description: "Quantity 49 (End of 10% discount tier)",
			price:       100.0,
			quantity:    49,
			expected:    4410.0, // 100 * 49 * 0.9
		},
		{
			description: "Quantity 50 (Start of 25% discount tier)",
			price:       100.0,
			quantity:    50,
			expected:    3750.0, // 100 * 50 * 0.75
		},
	}

	// Iterate through all test cases.
	for _, tc := range tests {
		got := calculateDiscount(tc.price, tc.quantity)

		// Note: For simple floating point arithmetic in these specific test cases, 
		// direct comparison is safe, but in production-grade math, an epsilon 
		// comparison would be used to handle precision errors.
		if got == tc.expected {
			fmt.Printf("PASS: %s\n", tc.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", tc.description, tc.expected, got)
		}
	}
}