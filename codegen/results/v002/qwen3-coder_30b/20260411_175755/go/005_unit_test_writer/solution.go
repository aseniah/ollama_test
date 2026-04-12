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
    expected := 90.0
    got := calculateDiscount(10.0, 9)
    if got == expected {
        fmt.Printf("PASS: quantity 9 (less than 10)\n")
    } else {
        fmt.Printf("FAIL: quantity 9 (less than 10) (expected: %.1f, got: %.1f)\n", expected, got)
    }

    // Test case 2: quantity = 10 (boundary condition - first discount tier)
    expected = 90.0
    got = calculateDiscount(10.0, 10)
    if got == expected {
        fmt.Printf("PASS: quantity 10 (first discount tier)\n")
    } else {
        fmt.Printf("FAIL: quantity 10 (first discount tier) (expected: %.1f, got: %.1f)\n", expected, got)
    }

    // Test case 3: quantity = 49 (boundary condition - less than 50)
    expected = 441.0
    got = calculateDiscount(10.0, 49)
    if got == expected {
        fmt.Printf("PASS: quantity 49 (less than 50)\n")
    } else {
        fmt.Printf("FAIL: quantity 49 (less than 50) (expected: %.1f, got: %.1f)\n", expected, got)
    }

    // Test case 4: quantity = 50 (boundary condition - second discount tier)
    expected = 375.0
    got = calculateDiscount(10.0, 50)
    if got == expected {
        fmt.Printf("PASS: quantity 50 (second discount tier)\n")
    } else {
        fmt.Printf("FAIL: quantity 50 (second discount tier) (expected: %.1f, got: %.1f)\n", expected, got)
    }

    // Additional test case: quantity = 0 (edge case)
    expected = 0.0
    got = calculateDiscount(10.0, 0)
    if got == expected {
        fmt.Printf("PASS: quantity 0 (edge case)\n")
    } else {
        fmt.Printf("FAIL: quantity 0 (edge case) (expected: %.1f, got: %.1f)\n", expected, got)
    }
}