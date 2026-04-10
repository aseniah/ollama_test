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

func testCase(description string, price float64, quantity int, expected float64) {
    got := calculateDiscount(price, quantity)
    if got == expected {
        fmt.Printf("PASS: %s\n", description)
    } else {
        fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", description, expected, got)
    }
}

func main() {
    // Test case 1: quantity 9 (below discount threshold, no discount)
    testCase("quantity 9, no discount", 10.0, 9, 90.0)

    // Test case 2: quantity 10 (10% discount threshold)
    testCase("quantity 10, 10%% discount", 10.0, 10, 90.0)

    // Test case 3: quantity 49 (upper limit for 10% discount)
    testCase("quantity 49, 10%% discount", 10.0, 49, 441.0)

    // Test case 4: quantity 50 (25% discount threshold)
    testCase("quantity 50, 25%% discount", 10.0, 50, 375.0)

    // Test case 5: quantity 100 (well into 25% discount range)
    testCase("quantity 100, 25%% discount", 5.0, 100, 375.0)
}