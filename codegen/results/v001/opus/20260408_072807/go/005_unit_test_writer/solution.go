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

func check(description string, expected, got float64) {
	if math.Abs(expected-got) < 1e-9 {
		fmt.Printf("PASS: %s\n", description)
	} else {
		fmt.Printf("FAIL: %s (expected: %g, got: %g)\n", description, expected, got)
	}
}

func main() {
	price := 10.0

	// quantity 9: no discount
	check("quantity 9 (no discount)", price*9, calculateDiscount(price, 9))

	// quantity 10: 10% discount boundary
	check("quantity 10 (10% discount boundary)", price*10*0.9, calculateDiscount(price, 10))

	// quantity 49: still 10% discount
	check("quantity 49 (10% discount)", price*49*0.9, calculateDiscount(price, 49))

	// quantity 50: 25% discount boundary
	check("quantity 50 (25% discount boundary)", price*50*0.75, calculateDiscount(price, 50))

	// quantity 1: minimum purchase
	check("quantity 1 (minimum)", price*1, calculateDiscount(price, 1))

	// quantity 100: large order
	check("quantity 100 (large order)", price*100*0.75, calculateDiscount(price, 100))
}
