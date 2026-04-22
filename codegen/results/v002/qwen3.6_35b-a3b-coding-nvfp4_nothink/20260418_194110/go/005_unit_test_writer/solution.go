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
		{"quantity 9 (just below 10)", 10.0, 9, 90.0},
		{"quantity 10 (at 10, 10% discount)", 10.0, 10, 90.0},
		{"quantity 49 (just below 50, 10% discount)", 10.0, 49, 441.0},
		{"quantity 50 (at 50, 25% discount)", 10.0, 50, 375.0},
	}

	for _, tt := range tests {
		result := calculateDiscount(tt.price, tt.quantity)
		// Use a small epsilon for float comparison
		diff := result - tt.expected
		if diff < 0 {
			diff = -diff
		}
		if diff < 1e-9 {
			fmt.Printf("PASS: %s\n", tt.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", tt.description, tt.expected, result)
		}
	}
}