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
    description string
    price       float64
    quantity    int
    expected    float64
}

func runTest(tc testCase) {
    got := calculateDiscount(tc.price, tc.quantity)
    
    // Use a small epsilon for floating point comparison
    const epsilon = 0.0001
    if abs(got - tc.expected) < epsilon {
        fmt.Printf("PASS: %s\n", tc.description)
    } else {
        fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.description, tc.expected, got)
    }
}

func abs(x float64) float64 {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
    // Test cases with boundary conditions
    testCases := []testCase{
        {
            description: "Quantity 9 (no discount, boundary before 10)",
            price:       10.0,
            quantity:    9,
            expected:    10.0 * 9.0, // 90.0
        },
        {
            description: "Quantity 10 (10% discount, boundary at 10)",
            price:       10.0,
            quantity:    10,
            expected:    10.0 * 10.0 * 0.9, // 90.0
        },
        {
            description: "Quantity 49 (10% discount, boundary before 50)",
            price:       10.0,
            quantity:    49,
            expected:    10.0 * 49.0 * 0.9, // 441.0
        },
        {
            description: "Quantity 50 (25% discount, boundary at 50)",
            price:       10.0,
            quantity:    50,
            expected:    10.0 * 50.0 * 0.75, // 375.0
        },
        {
            description: "Quantity 100 (25% discount, large quantity)",
            price:       5.0,
            quantity:    100,
            expected:    5.0 * 100.0 * 0.75, // 375.0
        },
        {
            description: "Quantity 1 (no discount, minimum quantity)",
            price:       20.0,
            quantity:    1,
            expected:    20.0 * 1.0, // 20.0
        },
    }

    fmt.Println("Running calculateDiscount function tests...")
    fmt.Println("===========================================")
    
    passCount := 0
    for _, tc := range testCases {
        runTest(tc)
        passCount++
    }
    
    fmt.Println("===========================================")
    fmt.Printf("Total tests: %d\n", len(testCases))
}