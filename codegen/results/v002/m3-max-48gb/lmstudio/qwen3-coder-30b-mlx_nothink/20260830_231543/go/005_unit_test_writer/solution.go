package main

import (
    "fmt"
    "math"
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
    // Test case 1: quantity = 9 (below threshold)
    expected1 := 90.0 // price = 10, quantity = 9, no discount
    actual1 := calculateDiscount(10.0, 9)
    if math.Abs(expected1-actual1) < 0.0001 {
        fmt.Printf("PASS: quantity 9 with price 10.0\n")
    } else {
        fmt.Printf("FAIL: quantity 9 with price 10.0 (expected: %.1f, got: %.1f)\n", expected1, actual1)
    }

    // Test case 2: quantity = 10 (at boundary, 10-49 gets 10% discount)
    expected2 := 90.0 // price = 10, quantity = 10, 10% discount
    actual2 := calculateDiscount(10.0, 10)
    if math.Abs(expected2-actual2) < 0.0001 {
        fmt.Printf("PASS: quantity 10 with price 10.0\n")
    } else {
        fmt.Printf("FAIL: quantity 10 with price 10.0 (expected: %.1f, got: %.1f)\n", expected2, actual2)
    }

    // Test case 3: quantity = 49 (at boundary, 10-49 gets 10% discount)
    expected3 := 441.0 // price = 10, quantity = 49, 10% discount
    actual3 := calculateDiscount(10.0, 49)
    if math.Abs(expected3-actual3) < 0.0001 {
        fmt.Printf("PASS: quantity 49 with price 10.0\n")
    } else {
        fmt.Printf("FAIL: quantity 49 with price 10.0 (expected: %.1f, got: %.1f)\n", expected3, actual3)
    }

    // Test case 4: quantity = 50 (at boundary, 50+ gets 25% discount)
    expected4 := 375.0 // price = 10, quantity = 50, 25% discount
    actual4 := calculateDiscount(10.0, 50)
    if math.Abs(expected4-actual4) < 0.0001 {
        fmt.Printf("PASS: quantity 50 with price 10.0\n")
    } else {
        fmt.Printf("FAIL: quantity 50 with price 10.0 (expected: %.1f, got: %.1f)\n", expected4, actual4)
    }

    // Test case 5: quantity = 5 (below threshold)
    expected5 := 50.0 // price = 10, quantity = 5, no discount
    actual5 := calculateDiscount(10.0, 5)
    if math.Abs(expected5-actual5) < 0.0001 {
        fmt.Printf("PASS: quantity 5 with price 10.0\n")
    } else {
        fmt.Printf("FAIL: quantity 5 with price 10.0 (expected: %.1f, got: %.1f)\n", expected5, actual5)
    }
}