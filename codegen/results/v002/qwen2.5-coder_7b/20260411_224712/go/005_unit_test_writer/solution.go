package main

import (
    "fmt"
)

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
        {"quantity less than 10", 10, 9, 90},
        {"quantity equal to 10", 10, 10, 100},
        {"quantity less than 50", 10, 49, 441},
        {"quantity equal to 50", 10, 50, 375},
    }

    for _, tc := range testCases {
        result := calculateDiscount(tc.price, tc.quantity)
        if result == tc.expected {
            fmt.Printf("PASS: %s (expected: %.2f, got: %.2f)\n", tc.description, tc.expected, result)
        } else {
            fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.description, tc.expected, result)
        }
    }
}