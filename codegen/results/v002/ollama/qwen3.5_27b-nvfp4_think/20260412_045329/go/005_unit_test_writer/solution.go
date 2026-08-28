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
		description string
		price       float64
		quantity    int
		expected    float64
	}{
		{"Quantity 9 (no discount)", 100, 9, 900.0},
		{"Quantity 10 (10% discount boundary)", 100, 10, 900.0},
		{"Quantity 49 (10% discount boundary)", 100, 49, 4410.0},
		{"Quantity 50 (25% discount boundary)", 100, 50, 3750.0},
	}

	for _, test := range tests {
		result := calculateDiscount(test.price, test.quantity)
		if result == test.expected {
			fmt.Printf("PASS: %s\n", test.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", test.description, test.expected, result)
		}
	}
}