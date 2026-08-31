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
    // Test case 1: quantity 9 (boundary condition - less than 10)
    expected1 := 90.0 // 10 * 9
    got1 := calculateDiscount(10.0, 9)
    if math.Abs(expected1-got1) < 0.0001 {
        fmt.Printf("PASS: quantity 9 (less than 10)\n")
    } else {
        fmt.Printf("FAIL: quantity 9 (less than 10) (expected: %.1f, got: %.1f)\n", expected1, got1)
    }

    // Test case 2: quantity 10 (boundary condition - first discount tier)
    expected2 := 90.0 // 10 * 10 * 0.9
    got2 := calculateDiscount(10.0, 10)
    if math.Abs(expected2-got2) < 0.0001 {
        fmt.Printf("PASS: quantity 10 (first discount tier)\n")
    } else {
        fmt.Printf("FAIL: quantity 10 (first discount tier) (expected: %.1f, got: %.1f)\n", expected2, got2)
    }

    // Test case 3: quantity 49 (boundary condition - less than 50)
    expected3 := 441.0 // 10 * 49 * 0.9
    got3 := calculateDiscount(10.0, 49)
    if math.Abs(expected3-got3) < 0.0001 {
        fmt.Printf("PASS: quantity 49 (less than 50)\n")
    } else {
        fmt.Printf("FAIL: quantity 49 (less than 50) (expected: %.1f, got: %.1f)\n", expected3, got3)
    }

    // Test case 4: quantity 50 (boundary condition - second discount tier)
    expected4 := 375.0 // 10 * 50 * 0.75
    got4 := calculateDiscount(10.0, 50)
    if math.Abs(expected4-got4) < 0.0001 {
        fmt.Printf("PASS: quantity 50 (second discount tier)\n")
    } else {
        fmt.Printf("FAIL: quantity 50 (second discount tier) (expected: %.1f, got: %.1f)\n", expected4, got4)
    }

    // Test case 5: quantity 100 (additional test case)
    expected5 := 750.0 // 10 * 100 * 0.75
    got5 := calculateDiscount(10.0, 100)
    if math.Abs(expected5-got5) < 0.0001 {
        fmt.Printf("PASS: quantity 100 (second discount tier)\n")
    } else {
        fmt.Printf("FAIL: quantity 100 (second discount tier) (expected: %.1f, got: %.1f)\n", expected5, got5)
    }
}