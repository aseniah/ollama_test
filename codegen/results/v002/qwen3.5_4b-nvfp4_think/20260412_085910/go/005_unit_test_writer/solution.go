package main

import (
	"fmt"
)

func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

func testDiscount() bool {
	passed := true
	tests := []struct {
		name        string
		price       float64
		quantity    int
		expected    float64
		description string
	}{
		{
			name:        "1 - Basic calculation (no discount)",
			price:       10.0,
			quantity:    5,
			expected:    50.0,
			description: "quantity < 10",
		},
		{
			name:        "2 - Boundary condition (quantity = 9)",
			price:       10.0,
			quantity:    9,
			expected:    90.0,
			description: "quantity < 10 boundary",
		},
		{
			name:        "3 - Boundary condition (quantity = 10)",
			price:       10.0,
			quantity:    10,
			expected:    90.0,
			description: "quantity >= 10, < 50",
		},
		{
			name:        "4 - Boundary condition (quantity = 49)",
			price:       10.0,
			quantity:    49,
			expected:    782.50,
			description: "quantity < 50 boundary",
		},
		{
			name:        "5 - Boundary condition (quantity = 50)",
			price:       10.0,
			quantity:    50,
			expected:    625.00,
			description: "quantity >= 50",
		},
	}

	totalTests := len(tests)
	passCount := 0

	for i, tc := range tests {
		result := calculateDiscount(tc.price, tc.quantity)
		expected := tc.expected
		
		if fmt.Sprintf("%.2f", result) != fmt.Sprintf("%.2f", expected) {
			if fmt.Sprintf("%.2f", result) == fmt.Sprintf("%.2f", expected) {
				fmt.Printf("PASS: %s\n", tc.description)
				passCount++
			} else {
				expectedStr := fmt.Sprintf("%.2f", expected)
				gotStr := fmt.Sprintf("%.2f", result)
				fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.description, expected, result)
				passed = false
			}
		} else {
			fmt.Printf("PASS: %s\n", tc.description)
			passCount++
		}
	}

	if passed {
		fmt.Println("\nAll tests passed successfully!")
	} else {
		fmt.Printf("\nResults - Passed: %d/%d tests failed", passCount, totalTests)
	}

	return passed
}

func main() {
	testDiscount()
}