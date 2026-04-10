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

func testCalculateDiscount() {
    tests := []struct {
        description string
        price       float64
        quantity    int
        expected    float64
    }{
        {"quantity less than 10", 10.0, 9, 90.0},
        {"quantity equal to 10", 10.0, 10, 100.0},
        {"quantity between 10 and 50", 10.0, 49, 441.0},
        {"quantity greater than or equal to 50", 10.0, 50, 375.0},
    }

    for _, test := range tests {
        result := calculateDiscount(test.price, test.quantity)
        if result != test.expected {
            fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", test.description, test.expected, result)
        } else {
            fmt.Printf("PASS: %s\n", test.description)
        }
    }
}

func main() {
    testCalculateDiscount()
}