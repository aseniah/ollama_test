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
	// Test Case 1: Boundary condition (quantity = 9, just below 10)
	expected := calculateDiscount(10.0, 9)
	got := calculateDiscount(10.0, 9)
	if expected == got {
		fmt.Println("PASS: quantity=9 (expecting:", expected, ")")
	} else {
		fmt.Printf("FAIL: quantity=9 (expected: %.2f, got: %.2f)\n", expected, got)
	}

	// Test Case 2: Boundary condition (quantity = 10, transition point)
	expected = calculateDiscount(10.0, 10)
	got = calculateDiscount(10.0, 10)
	if expected == got {
		fmt.Println("PASS: quantity=10 (expecting:", expected, ")")
	} else {
		fmt.Printf("FAIL: quantity=10 (expected: %.2f, got: %.2f)\n", expected, got)
	}

	// Test Case 3: Boundary condition (quantity = 49, just before 50)
	expected = calculateDiscount(10.0, 49)
	got = calculateDiscount(10.0, 49)
	if expected == got {
		fmt.Println("PASS: quantity=49 (expecting:", expected, ")")
	} else {
		fmt.Printf("FAIL: quantity=49 (expected: %.2f, got: %.2f)\n", expected, got)
	}

	// Test Case 4: Boundary condition (quantity = 50, transition point)
	expected = calculateDiscount(10.0, 50)
	got = calculateDiscount(10.0, 50)
	if expected == got {
		fmt.Println("PASS: quantity=50 (expecting:", expected, ")")
	} else {
		fmt.Printf("FAIL: quantity=50 (expected: %.2f, got: %.2f)\n", expected, got)
	}

	// Test Case 5: Normal case (well within first tier)
	expected = calculateDiscount(10.0, 5)
	got = calculateDiscount(10.0, 5)
	if expected == got {
		fmt.Println("PASS: quantity=5 (expecting:", expected, ")")
	} else {
		fmt.Printf("FAIL: quantity=5 (expected: %.2f, got: %.2f)\n", expected, got)
	}

	// Test Case 6: Normal case (well within second tier)
	expected = calculateDiscount(10.0, 30)
	got = calculateDiscount(10.0, 30)
	if expected == got {
		fmt.Println("PASS: quantity=30 (expecting:", expected, ")")
	} else {
		fmt.Printf("FAIL: quantity=30 (expected: %.2f, got: %.2f)\n", expected, got)
	}

	// Test Case 7: Normal case (well within third tier)
	expected = calculateDiscount(10.0, 60)
	got = calculateDiscount(10.0, 60)
	if expected == got {
		fmt.Println("PASS: quantity=60 (expecting:", expected, ")")
	} else {
		fmt.Printf("FAIL: quantity=60 (expected: %.2f, got: %.2f)\n", expected, got)
	}

	fmt.Println("\nAll tests completed.")
}