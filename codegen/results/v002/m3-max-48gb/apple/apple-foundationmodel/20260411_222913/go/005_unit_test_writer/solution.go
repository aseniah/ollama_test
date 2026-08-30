package main

import "fmt"

// Function to calculate bulk discount
func calculateDiscount(price float64, quantity int) float64 {
    if quantity < 10 {
        return price * float64(quantity)
    } else if quantity < 50 {
        return price * float64(quantity) * 0.9
    }
    return price * float64(quantity) * 0.75
}

func main() {
    // Test cases
    testCases := []struct {
        description string
        quantity    int
        expected    float64
    }{
        {"Test case 1: Quantity 9", 9, 9.0, 9.0},
        {"Test case 2: Quantity 10", 10, 9.0, 9.0},
        {"Test case 3: Quantity 49", 49, 34.5, 34.5},
        {"Test case 4: Quantity 50", 50, 37.5, 37.5},
    }

    // Run tests
    for _, test := range testCases {
        result := calculateDiscount(10.0, test.quantity)
        fmt.Printf("PASS: %s\n", test.description)
        fmt.Printf("  Expected: %f, Got: %f\n", test.expected, result)
    }
}