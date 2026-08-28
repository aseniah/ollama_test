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
		name     string
		price    float64
		quantity int
		expected float64
	}{
		{"quantity 9 (below discount threshold)", 10.0, 9, 90.0},
		{"quantity 10 (10% discount boundary)", 10.0, 10, 90.0},
		{"quantity 49 (10% discount upper boundary)", 10.0, 49, 441.0},
		{"quantity 50 (25% discount boundary)", 10.0, 50, 375.0},
	}

	const epsilon = 1e-9

	for _, t := range tests {
		result := calculateDiscount(t.price, t.quantity)
		if math.Abs(result-t.expected) < epsilon {
			fmt.Printf("PASS: %s\n", t.name)
		} else {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", t.name, t.expected, result)
		}
	}
}