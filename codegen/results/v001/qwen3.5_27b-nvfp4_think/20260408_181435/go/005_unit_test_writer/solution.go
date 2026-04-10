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

func abs(x float64) float64 {
    if x < 0 {
        return -x
    }
    return x
}

func runTest(description string, price float64, quantity int, expected float64) {
    got := calculateDiscount(price, quantity)

    const epsilon = 0.0001

    if abs(got-expected) < epsilon {
        fmt.Printf("PASS: %s\n", description)
    } else {
        fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", description, expected, got)
    }
}

func main() {
    // Test case 1: quantity 9 - no discount (< 10)
    runTest("Quantity 9 (no discount)", 10.0, 9, 90.0)

    // Test case 2: quantity 10 - 10% discount (>= 10 and < 50)
    runTest("Quantity 10 (10%% discount)", 10.0, 10, 90.0)

    // Test case 3: quantity 49 - 10% discount (< 50)
    runTest("Quantity 49 (10%% discount)", 10.0, 49, 441.0)

    // Test case 4: quantity 50 - 25% discount (>= 50)
    runTest("Quantity 50 (25%% discount)", 10.0, 50, 375.0)
}