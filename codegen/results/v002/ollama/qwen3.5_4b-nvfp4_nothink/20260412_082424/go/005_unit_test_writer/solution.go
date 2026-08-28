package main

import (
	"fmt"
)

// The function to test, copied as requested.
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

func main() {
	tests := []struct {
		price   float64
		quantity int
		expected float64
		desc    string
	}{
		{10.0, 5, 50.0, "quantity < 10: no discount"},
		{10.0, 9, 90.0, "quantity = 9 (boundary, no discount)"},
		{10.0, 10, 90.0, "quantity = 10 (boundary, 10% off)"},
		{10.0, 49, 450.0, "quantity = 49 (boundary, 10% off)"},
		{10.0, 50, 375.0, "quantity = 50 (boundary, 25% off)"},
		{10.0, 100, 750.0, "large quantity: 25% off"},
	}

	passCount := 0
	failCount := 0

	for _, tc := range tests {
		result := calculateDiscount(tc.price, tc.quantity)
		
		// Use a small tolerance for floating point comparison (optional but safe)
		isPass := fmt.Sprintf("%.2f", result) == fmt.Sprintf("%.2f", tc.expected) || 
		          (result == tc.expected)

		if isPass {
			fmt.Printf("PASS: %s\n", tc.desc)
			passCount++
		} else {
			expectedStr := fmt.Sprintf("%.2f", tc.expected)
			actualStr := fmt.Sprintf("%.2f", result)
			fmt.Printf("FAIL: %s (expected: %s, got: %s)\n", tc.desc, expectedStr, actualStr)
			failCount++
		}
	}

	fmt.Println()
	fmt.Printf("Test Summary: %d passed, %d failed\n", passCount, failCount)
	
	if passCount == len(tests) {
		fmt.Println("All tests passed!")
	} else {
		fmt.Println("Some tests failed.")
	}
}