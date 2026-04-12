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
	fmt.Println("Testing calculateDiscount function:")
	fmt.Println("")

	testCases := []struct {
		description string
		price        float64
		quantity     int
		expected     float64
	}{
		{
			description: "quantity=9 (boundary < 10)",
			price:       25.0,
			quantity:    9,
			expected:    225.0,
		},
		{
			description: "quantity=10 (boundary >= 10, < 50)",
			price:       25.0,
			quantity:    10,
			expected:    225.0,
		},
		{
			description: "quantity=49 (boundary < 50)",
			price:       25.0,
			quantity:    49,
			expected:    1102.5,
		},
		{
			description: "quantity=50 (boundary >= 50)",
			price:       25.0,
			quantity:    50,
			expected:    937.5,
		},
	}

	passCount := 0
	failCount := 0

	for _, tc := range testCases {
		result := calculateDiscount(tc.price, tc.quantity)
		if result == tc.expected {
			fmt.Printf("PASS: %s\n", tc.description)
			passCount++
		} else {
			expectedMsg := fmt.Sprintf("%.2f", tc.expected)
			gotMsg := fmt.Sprintf("%.2f", result)
			fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.description, tc.expected, gotMsg)
			failCount++
		}
	}

	fmt.Println("")
	fmt.Printf("Test Summary: %d passed, %d failed\n", passCount, failCount)
	if failCount == 0 {
		fmt.Println("All tests passed successfully!")
	} else {
		fmt.Println("Some tests failed.")
	}
}