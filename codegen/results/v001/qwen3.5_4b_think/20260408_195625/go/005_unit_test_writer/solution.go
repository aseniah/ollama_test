package main

import "fmt"

// calculateDiscount calculates a bulk discount based on quantity
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
        price      float64
        quantity   int
        expected   float64
        description string
    }{
        {10.00, 9, 90.00, "Single item discount (quantity < 10)"},
        {10.00, 10, 90.00, "No discount applied at boundary (quantity = 10)"},
        {10.00, 49, 441.00, "Discount at 0.9 (quantity = 49)"},
        {10.00, 50, 375.00, "Discount at 0.75 (quantity = 50)"},
    }

    for _, tc := range testCases {
        result := calculateDiscount(tc.price, tc.quantity)
        
        // Allow a small floating point error
        if result == tc.expected {
            fmt.Printf("PASS: %s\n", tc.description)
        } else {
            fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", tc.description, tc.expected, result)
        }
    }
}