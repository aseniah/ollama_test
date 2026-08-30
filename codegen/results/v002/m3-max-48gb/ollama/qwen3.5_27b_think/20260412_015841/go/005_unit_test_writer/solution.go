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
    testCases := []struct {
        description string
        price       float64
        quantity    int
        expected    float64
    }{
        {"Quantity 9 (no discount)", 10.0, 9, 10.0 * 9},
        {"Quantity 10 (10% discount)", 10.0, 10, 10.0 * 10 * 0.9},
        {"Quantity 49 (10% discount)", 10.0, 49, 10.0 * 49 * 0.9},
        {"Quantity 50 (25% discount)", 10.0, 50, 10.0 * 50 * 0.75},
    }

    for _, tc := range testCases {
        result := calculateDiscount(tc.price, tc.quantity)
        if result == tc.expected {
            fmt.Printf("PASS: %s\n", tc.description)
        } else {
            fmt.Printf("FAIL: %s (expected: %f, got: %f)\n", tc.description, tc.expected, result)
        }
    }
}