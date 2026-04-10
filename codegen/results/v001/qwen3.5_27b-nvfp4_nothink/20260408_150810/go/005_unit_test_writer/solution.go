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

func runTest(name string, price float64, quantity int, expected float64) {
	got := calculateDiscount(price, quantity)

	// Use a small epsilon for floating point comparison to handle precision issues
	epsilon := 1e-9
	if abs(got-expected) < epsilon {
		fmt.Printf("PASS: %s\n", name)
	} else {
		fmt.Printf("FAIL: %s (expected: %.4f, got: %.4f)\n", name, expected, got)
	}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println("Running tests for calculateDiscount...")

	// Test Case 1: Boundary condition (quantity 9, no discount)
	// Price: 10.00, Quantity: 9 -> Expected: 10 * 9 = 90.00
	runTest("Quantity 9 (No Discount)", 10.0, 9, 90.0)

	// Test Case 2: Boundary condition (quantity 10, 10% discount starts)
	// Price: 10.00, Quantity: 10 -> Expected: 10 * 10 * 0.9 = 90.00
	runTest("Quantity 10 (10% Discount)", 10.0, 10, 90.0)

	// Test Case 3: Boundary condition (quantity 49, still 10% discount)
	// Price: 20.00, Quantity: 49 -> Expected: 20 * 49 * 0.9 = 882.00
	runTest("Quantity 49 (10% Discount)", 20.0, 49, 882.0)

	// Test Case 4: Boundary condition (quantity 50, 25% discount starts)
	// Price: 20.00, Quantity: 50 -> Expected: 20 * 50 * 0.75 = 750.00
	runTest("Quantity 50 (25% Discount)", 20.0, 50, 750.0)

	// Additional Test Case 5: High quantity to ensure logic holds
	// Price: 10.00, Quantity: 100 -> Expected: 10 * 100 * 0.75 = 750.00
	runTest("Quantity 100 (25% Discount)", 10.0, 100, 750.0)

	fmt.Println("Tests completed.")
}