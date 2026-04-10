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
    // Test case 1: quantity < 10 (boundary condition - 9)
    expected := 90.0
    got := calculateDiscount(10.0, 9)
    if got == expected {
        fmt.Println("PASS: quantity 9 (no discount)")
    } else {
        fmt.Printf("FAIL: quantity 9 (no discount) (expected: %.1f, got: %.1f)\n", expected, got)
    }

    // Test case 2: quantity = 10 (boundary condition - first discount tier)
    expected = 90.0
    got = calculateDiscount(10.0, 10)
    if got == expected {
        fmt.Println("PASS: quantity 10 (10% discount)")
    } else {
        fmt.Printf("FAIL: quantity 10 (10% discount) (expected: %.1f, got: %.1f)\n", expected, got)
    }

    // Test case 3: quantity = 49 (boundary condition - last discount tier before highest)
    expected = 441.0
    got = calculateDiscount(10.0, 49)
    if got == expected {
        fmt.Println("PASS: quantity 49 (10% discount)")
    } else {
        fmt.Printf("FAIL: quantity 49 (10% discount) (expected: %.1f, got: %.1f)\n", expected, got)
    }

    // Test case 4: quantity = 50 (boundary condition - highest discount tier)
    expected = 375.0
    got = calculateDiscount(10.0, 50)
    if got == expected {
        fmt.Println("PASS: quantity 50 (25% discount)")
    } else {
        fmt.Printf("FAIL: quantity 50 (25% discount) (expected: %.1f, got: %.1f)\n", expected, got)
    }

    // Test case 5: quantity = 0 (edge case)
    expected = 0.0
    got = calculateDiscount(10.0, 0)
    if got == expected {
        fmt.Println("PASS: quantity 0 (no items)")
    } else {
        fmt.Printf("FAIL: quantity 0 (no items) (expected: %.1f, got: %.1f)\n", expected, got)
    }
}