package main

// calculateDiscount calculates the total price with bulk discounts applied.
func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

// testCase defines a single test case for the calculateDiscount function.
type testCase struct {
	description string
	price       float64
	quantity    int
	expected    float64
}

func main() {
	testCases := []testCase{
		{
			description: "Quantity 9 (No Discount)",
			price:       10.0,
			quantity:    9,
			expected:    10.0 * 9, // 90.0
		},
		{
			description: "Quantity 10 (First Boundary - 10% Discount)",
			price:       10.0,
			quantity:    10,
			expected:    10.0 * 10 * 0.9, // 90.0
		},
		{
			description: "Quantity 49 (Upper Bound of Second Tier - 10% Discount)",
			price:       20.0,
			quantity:    49,
			expected:    20.0 * 49 * 0.9, // 882.0
		},
		{
			description: "Quantity 50 (Second Boundary - 25% Discount)",
			price:       20.0,
			quantity:    50,
			expected:    20.0 * 50 * 0.75, // 750.0
		},
	}

	allPassed := true

	for _, tc := range testCases {
		result := calculateDiscount(tc.price, tc.quantity)

		// Use a small epsilon for float comparison to handle potential precision issues
		const epsilon = 0.0001
		if abs(result-tc.expected) < epsilon {
			println("PASS:", tc.description)
		} else {
			allPassed = false
			println("FAIL:", tc.description, "(expected:", tc.expected, ", got:", result, ")")
		}
	}

	if allPassed {
		println("\nAll tests passed!")
	} else {
		println("\nSome tests failed.")
	}
}

// abs returns the absolute value of a float64.
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}