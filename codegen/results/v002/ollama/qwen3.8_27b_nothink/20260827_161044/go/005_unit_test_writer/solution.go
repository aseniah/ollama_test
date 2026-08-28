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

func testDiscount(price float64, quantity int, expected float64) {
	got := calculateDiscount(price, quantity)
	// Use a small epsilon for floating point comparison
	const epsilon = 1e-9
	if (got - expected) < epsilon && (expected - got) < epsilon {
		fmt.Printf("PASS: price=%v, quantity=%d\n", price, quantity)
	} else {
		fmt.Printf("FAIL: price=%v, quantity=%d (expected: %v, got: %v)\n", price, quantity, expected, got)
	}
}

func main() {
	// Test case 1: quantity = 9 (below 10, no discount)
	testDiscount(10.0, 9, 10.0*9.0)

	// Test case 2: quantity = 10 (boundary, 10% discount)
	testDiscount(10.0, 10, 10.0*10.0*0.9)

	// Test case 3: quantity = 49 (below 50, 10% discount)
	testDiscount(10.0, 49, 10.0*49.0*0.9)

	// Test case 4: quantity = 50 (boundary, 25% discount)
	testDiscount(10.0, 50, 10.0*50.0*0.75)

	// Additional test cases for thoroughness
	// Test case 5: quantity = 1 (minimum)
	testDiscount(5.0, 1, 5.0*1.0)

	// Test case 6: quantity = 25 (middle of 10-49 range)
	testDiscount(20.0, 25, 20.0*25.0*0.9)

	// Test case 7: quantity = 100 (well above 50)
	testDiscount(15.0, 100, 15.0*100.0*0.75)
}