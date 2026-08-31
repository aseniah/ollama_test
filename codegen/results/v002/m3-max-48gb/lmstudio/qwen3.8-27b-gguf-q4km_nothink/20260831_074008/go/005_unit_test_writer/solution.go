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
		{"quantity 9, price 100", 100, 9, 100 * 9},
		{"quantity 10, price 100", 100, 10, 100 * 10 * 0.9},
		{"quantity 49, price 100", 100, 49, 100 * 49 * 0.9},
		{"quantity 50, price 100", 100, 50, 100 * 50 * 0.75},
	}

	for _, t := range tests {
		got := calculateDiscount(t.price, t.quantity)
		if got == t.expected {
			fmt.Printf("PASS: %s\n", t.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", t.description, t.expected, got)
		}
	}
}