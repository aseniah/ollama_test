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
    // Test case 1: quantity = 9 (boundary condition - less than 10)
    expected1 := 90.0 // price = 10, quantity = 9, no discount
    actual1 := calculateDiscount(10.0, 9)
    if actual1 == expected1 {
        fmt.Println("PASS: quantity 9 with price 10")
    } else {
        fmt.Printf("FAIL: quantity 9 with price 10 (expected: %.1f, got: %.1f)\n", expected1, actual1)
    }

    // Test case 2: quantity = 10 (boundary condition - first discount tier)
    expected2 := 90.0 // price = 10, quantity = 10, 10% discount
    actual2 := calculateDiscount(10.0, 10)
    if actual2 == expected2 {
        fmt.Println("PASS: quantity 10 with price 10")
    } else {
        fmt.Printf("FAIL: quantity 10 with price 10 (expected: %.1f, got: %.1f)\n", expected2, actual2)
    }

    // Test case 3: quantity = 49 (boundary condition - second discount tier)
    expected3 := 441.0 // price = 10, quantity = 49, 10% discount
    actual3 := calculateDiscount(10.0, 49)
    if actual3 == expected3 {
        fmt.Println("PASS: quantity 49 with price 10")
    } else {
        fmt.Printf("FAIL: quantity 49 with price 10 (expected: %.1f, got: %.1f)\n", expected3, actual3)
    }

    // Test case 4: quantity = 50 (boundary condition - second discount tier)
    expected4 := 375.0 // price = 10, quantity = 50, 25% discount
    actual4 := calculateDiscount(10.0, 50)
    if actual4 == expected4 {
        fmt.Println("PASS: quantity 50 with price 10")
    } else {
        fmt.Printf("FAIL: quantity 50 with price 10 (expected: %.1f, got: %.1f)\n", expected4, actual4)
    }
}