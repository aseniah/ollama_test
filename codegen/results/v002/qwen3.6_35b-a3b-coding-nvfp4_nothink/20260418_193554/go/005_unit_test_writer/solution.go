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
	// Test case 1: quantity = 9 (should use no discount, price * quantity)
	expected1 := 100.0 * 9.0
	got1 := calculateDiscount(100.0, 9)
	if expected1 == got1 {
		fmt.Println("PASS: quantity 9 gives no discount")
	} else {
		fmt.Printf("FAIL: quantity 9 gives no discount (expected: %.2f, got: %.2f)\n", expected1, got1)
	}

	// Test case 2: quantity = 10 (should use 10% discount, price * quantity * 0.9)
	expected2 := 100.0 * 10.0 * 0.9
	got2 := calculateDiscount(100.0, 10)
	if expected2 == got2 {
		fmt.Println("PASS: quantity 10 gives 10% discount")
	} else {
		fmt.Printf("FAIL: quantity 10 gives 10% discount (expected: %.2f, got: %.2f)\n", expected2, got2)
	}

	// Test case 3: quantity = 49 (should use 10% discount, price * quantity * 0.9)
	expected3 := 100.0 * 49.0 * 0.9
	got3 := calculateDiscount(100.0, 49)
	if expected3 == got3 {
		fmt.Println("PASS: quantity 49 gives 10% discount")
	} else {
		fmt.Printf("FAIL: quantity 49 gives 10% discount (expected: %.2f, got: %.2f)\n", expected3, got3)
	}

	// Test case 4: quantity = 50 (should use 25% discount, price * quantity * 0.75)
	expected4 := 100.0 * 50.0 * 0.75
	got4 := calculateDiscount(100.0, 50)
	if expected4 == got4 {
		fmt.Println("PASS: quantity 50 gives 25% discount")
	} else {
		fmt.Printf("FAIL: quantity 50 gives 25% discount (expected: %.2f, got: %.2f)\n", expected4, got4)
	}
}