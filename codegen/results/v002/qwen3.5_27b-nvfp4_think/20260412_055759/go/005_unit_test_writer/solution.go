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

type testCase struct {
	price      float64
	quantity   int
	expected   float64
	description string
}

func runTest(tc testCase) {
	got := calculateDiscount(tc.price, tc.quantity)

	if got == tc.expected {
		fmt.Printf("PASS: %s\n", tc.description)
	} else {
		fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.description, tc.expected, got)
	}
}

func main() {
	testCases := []testCase{
		{price: 100.0, quantity: 9, expected: 900.0, description: "quantity 9 (no discount)"},
		{price: 100.0, quantity: 10, expected: 900.0, description: "quantity 10 (boundary - 10%% discount)"},
		{price: 100.0, quantity: 49, expected: 4410.0, description: "quantity 49 (10%% discount)"},
		{price: 100.0, quantity: 50, expected: 3750.0, description: "quantity 50 (boundary - 25%% discount)"},
	}

	fmt.Println("Running calculateDiscount tests...")
	fmt.Println()

	for _, tc := range testCases {
		runTest(tc)
	}
}