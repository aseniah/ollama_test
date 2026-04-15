package main

import (
	"fmt"
	"math"
)

// calculateDiscount calculates the bulk discount based on quantity
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

// approximateEqual checks if two float64 values are approximately equal
func approximateEqual(a, b float64, epsilon float64) bool {
	diff := math.Abs(a - b)
	return diff <= epsilon
}

func main() {
	tests := []struct {
		description string
		price       float64
		quantity    int
		expected    float64
	}{
		{
			description: "Quantity 9 (no discount)",
			price:       10.0,
			quantity:    9,
			expected:    10.0 * 9.0, // 90.0
		},
		{
			description: "Quantity 10 (10% discount)",
			price:       10.0,
			quantity:    10,
			expected:    10.0 * 10.0 * 0.9, // 90.0
		},
		{
			description: "Quantity 49 (10% discount)",
			price:       10.0,
			quantity:    49,
			expected:    10.0 * 49.0 * 0.9, // 441.0
		},
		{
			description: "Quantity 50 (25% discount)",
			price:       10.0,
			quantity:    50,
			expected:    10.0 * 50.0 * 0.75, // 375.0
		},
	}

	epsilon := 0.0001 // tolerance for floating point comparison

	for _, t := range tests {
		result := calculateDiscount(t.price, t.quantity)
		if approximateEqual(result, t.expected, epsilon) {
			fmt.Printf("PASS: %s\n", t.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", t.description, t.expected, result)
		}
	}
}