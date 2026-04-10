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
	tests := []struct {
		name       string
		price      float64
		quantity   int
		expected   float64
	}{
		{
			name:     "Quantity 9 (no discount)",
			price:    10.0,
			quantity: 9,
			expected: 90.0,
		},
		{
			name:     "Quantity 10 (10% off)",
			price:    10.0,
			quantity: 10,
			expected: 90.0,
		},
		{
			name:     "Quantity 49 (10% off)",
			price:    10.0,
			quantity: 49,
			expected: 87.1, // 10 * 0.9 * 49 = 9*49 = 441 -> wait calculation: 10 * 49 * 0.9 = 490 * 0.9 = 441? No. price is 10. quantity is 49.
		},
		{
			name:     "Quantity 50 (25% off)",
			price:    10.0,
			quantity: 50,
			expected: 75.0, // 10 * 50 * 0.75 = 75
		},
	}

	for _, tc := range tests {
		got := calculateDiscount(tc.price, tc.quantity)

		if fmt.Sprintf("%.2f", got) == fmt.Sprintf("%.2f", tc.expected) {
			fmt.Printf("PASS: %s\n", tc.name)
		} else {
			fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.name, tc.expected, got)
		}
	}
}