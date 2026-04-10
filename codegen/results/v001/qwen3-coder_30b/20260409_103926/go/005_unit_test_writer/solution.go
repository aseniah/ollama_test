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
    // Test case 1: quantity = 9 (boundary condition - less than 10)
    expected1 := 90.0
    got1 := calculateDiscount(10.0, 9)
    if got1 == expected1 {
        fmt.Printf("PASS: quantity 9 (less than 10)\n")
    } else {
        fmt.Printf("FAIL: quantity 9 (less than 10) (expected: %.1f, got: %.1f)\n", expected1, got1)
    }

    // Test case 2: quantity = 10 (boundary condition - first discount tier)
    expected2 := 90.0
    got2 := calculateDiscount(10.0, 10)
    if got2 == expected2 {
        fmt.Printf("PASS: quantity 10 (first discount tier)\n")
    } else {
        fmt.Printf("FAIL: quantity 10 (first discount tier) (expected: %.1f, got: %.1f)\n", expected2, got2)
    }

    // Test case 3: quantity = 49 (boundary condition - less than 50)
    expected3 := 441.0
    got3 := calculateDiscount(10.0, 49)
    if got3 == expected3 {
        fmt.Printf("PASS: quantity 49 (less than 50)\n")
    } else {
        fmt.Printf("FAIL: quantity 49 (less than 50) (expected: %.1f, got: %.1f)\n", expected3, got3)
    }

    // Test case 4: quantity = 50 (boundary condition - second discount tier)
    expected4 := 375.0
    got4 := calculateDiscount(10.0, 50)
    if got4 == expected4 {
        fmt.Printf("PASS: quantity 50 (second discount tier)\n")
    } else {
        fmt.Printf("FAIL: quantity 50 (second discount tier) (expected: %.1f, got: %.1f)\n", expected4, got4)
    }

    // Additional test case: quantity = 0 (edge case)
    expected5 := 0.0
    got5 := calculateDiscount(10.0, 0)
    if got5 == expected5 {
        fmt.Printf("PASS: quantity 0 (edge case)\n")
    } else {
        fmt.Printf("FAIL: quantity 0 (edge case) (expected: %.1f, got: %.1f)\n", expected5, got5)
    }
}