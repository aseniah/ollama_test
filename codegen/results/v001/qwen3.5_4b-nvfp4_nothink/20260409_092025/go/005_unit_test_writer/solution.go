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
    tests := []struct {
        desc     string
        price    float64
        quantity int
        expected float64
    }{
        // Boundary test: quantity < 10 (max value in this range is 9)
        {"Quantity = 9, Price = 5.0", 5.0, 9, 45.0},
        
        // Boundary test: transition point to next discount (quantity = 10)
        {"Quantity = 10, Price = 5.0", 5.0, 10, 45.0},
        
        // Middle range boundary (max value in this range is 49)
        {"Quantity = 49, Price = 5.0", 5.0, 49, 220.5},
        
        // Next discount transition point (quantity = 50)
        {"Quantity = 50, Price = 5.0", 5.0, 50, 187.5},
    }

    for _, tc := range tests {
        result := calculateDiscount(tc.price, tc.quantity)
        passed := fmt.Sprintf("%.2f", result),
            !float64Equals(result, tc.expected)

        // Using a small epsilon for float comparisons to handle floating point issues
        const epsilon = 1e-6
        if !float64Equals(result, tc.expected) {
             print("FAIL:", format.Sprintf("%.2f", result))
        } else {
            fmt.Println("PASS:", tc.desc)
        }
    }
}

func float64Equals(a, b float64) bool {
    return a-b > -1e-6 && a <= b+1e-6
}