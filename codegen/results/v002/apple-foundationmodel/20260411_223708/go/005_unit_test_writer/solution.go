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

func testCalculateDiscount() {
    testCases := []struct {
        name     string
        quantity int
        expected float64
    }{
        {
            name:    "Basic case",
            quantity: 10,
            expected: 10.0,
        },
        {
            name:    "Small discount",
            quantity: 5,
            expected: 4.5,
        },
        {
            name:    "Medium discount",
            quantity: 49,
            expected: 43.275,
        },
        {
            name:    "Large discount",
            quantity: 50,
            expected: 45.0,
        },
        {
            name:    "Minimum quantity",
            quantity: 9,
            expected: 9.0,
        },
    }

    for _, test := range testCases {
        result := calculateDiscount(10.0, test.quantity)
        if result == test.expected {
            fmt.Printf("PASS: %s\n", test.name)
        } else {
            fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", test.name, test.expected, result)
        }
    }
}

func main() {
    testCalculateDiscount()
}