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
		{
			description: "Quantity 9 (below first threshold)",
			price:       10.0,
			quantity:    9,
			expected:    90.0, // 10 * 9
		},
		{
			description: "Quantity 10 (at first threshold)",
			price:       10.0,
			quantity:    10,
			expected:    90.0, // 10 * 10 * 0.9
		},
		{
			description: "Quantity 49 (at second threshold)",
			price:       20.0,
			quantity:    49,
			expected:    882.0, // 20 * 49 * 0.9
		},
		{
			description: "Quantity 50 (at third threshold)",
			price:       20.0,
			quantity:    50,
			expected:    750.0, // 20 * 50 * 0.75
		},
	}

	for _, t := range tests {
		got := calculateDiscount(t.price, t.quantity)
		// Use a small epsilon for float comparison
		diff := got - t.expected
		if diff < 0 {
			diff = -diff
		}
		const epsilon = 1e-9
		if diff <= epsilon {
			fmt.Printf("PASS: %s\n", t.description)
		} else {
			fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", t.description, t.expected, got)
		}
	}
}