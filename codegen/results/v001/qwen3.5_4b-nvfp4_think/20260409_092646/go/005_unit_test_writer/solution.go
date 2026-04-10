package main

import (
	"fmt"
	"math"
)

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
		name       string
		price      float64
		quantity   int
		expected   float64
		assertions int
	}{
		{
			name:       "Price $10, Quantity 9",
			price:      10.0,
			quantity:   9,
			expected:   90.0,
			assertions: 0,
		},
		{
			name:       "Price $20, Quantity 10",
			price:      20.0,
			quantity:   10,
			expected:   180.0, // 20 * 10 * 0.9
			assertions: 1,
		},
		{
			name:       "Price $50, Quantity 49",
			price:      50.0,
			quantity:   49,
			expected:   2205.0, // 50 * 49 * 0.9
			assertions: 2,
		},
		{
			name:       "Price $30, Quantity 50",
			price:      30.0,
			quantity:   50,
			expected:   1125.0, // 30 * 50 * 0.75
			assertions: 3,
		},
	}

	passCount := 0
	failCount := 0

	for i := range tests {
		test := &tests[i]

		got := calculateDiscount(test.price, test.quantity)
		expected := math.Float64fromInt(int64(test.expected))

		if math.Abs(got-expected) < 0.001 {
			fmt.Printf("PASS: %s\n", test.name)
			passCount++
		} else {
			fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", test.name, expected, got)
			failCount++
		}
	}

	fmt.Println("")
	if passCount > 0 && failCount == 0 {
		fmt.Printf("All tests passed! (+%d tests)\n", passCount)
	} else if failCount > 0 {
		fmt.Printf("Total: +%d passed, +%d failed\n", passCount, failCount)
	} else {
		fmt.Println("No tests were run.")
	}
}