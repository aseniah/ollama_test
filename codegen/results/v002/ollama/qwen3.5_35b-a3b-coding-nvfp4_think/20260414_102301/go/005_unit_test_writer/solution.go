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
        {"No discount (quantity=9)", 100.0, 9, 900.0},
        {"10% discount (quantity=10)", 100.0, 10, 900.0},
        {"10% discount (quantity=49)", 100.0, 49, 4410.0},
        {"25% discount (quantity=50)", 100.0, 50, 3750.0},
    }

    for _, tc := range testCases {
        result := calculateDiscount(tc.price, tc.quantity)
        if approxEqual(result, tc.expected) {
            fmt.Printf("PASS: %s\n", tc.description)
        } else {
            fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.description, tc.expected, result)
        }
    }
}

func approxEqual(a, b float64) bool {
    return abs(a-b) < 0.001
}

func abs(x float64) float64 {
    if x < 0 {
        return -x
    }
    return x
}