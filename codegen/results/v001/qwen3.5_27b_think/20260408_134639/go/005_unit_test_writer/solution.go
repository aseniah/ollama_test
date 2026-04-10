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
    // Test case 1: Quantity 9 (no discount)
    price1 := 10.0
    quantity1 := 9
    expected1 := price1 * float64(quantity1)
    got1 := calculateDiscount(price1, quantity1)
    if got1 == expected1 {
        fmt.Println("PASS: Quantity 9 - No discount applied")
    } else {
        fmt.Printf("FAIL: Quantity 9 - No discount applied (expected: %.2f, got: %.2f)\n", expected1, got1)
    }

    // Test case 2: Quantity 10 (10% discount - boundary)
    price2 := 10.0
    quantity2 := 10
    expected2 := price2 * float64(quantity2) * 0.9
    got2 := calculateDiscount(price2, quantity2)
    if got2 == expected2 {
        fmt.Println("PASS: Quantity 10 - 10%% discount applied (boundary)")
    } else {
        fmt.Printf("FAIL: Quantity 10 - 10%% discount applied (boundary) (expected: %.2f, got: %.2f)\n", expected2, got2)
    }

    // Test case 3: Quantity 49 (10% discount - boundary)
    price3 := 10.0
    quantity3 := 49
    expected3 := price3 * float64(quantity3) * 0.9
    got3 := calculateDiscount(price3, quantity3)
    if got3 == expected3 {
        fmt.Println("PASS: Quantity 49 - 10%% discount applied (boundary)")
    } else {
        fmt.Printf("FAIL: Quantity 49 - 10%% discount applied (boundary) (expected: %.2f, got: %.2f)\n", expected3, got3)
    }

    // Test case 4: Quantity 50 (25% discount - boundary)
    price4 := 10.0
    quantity4 := 50
    expected4 := price4 * float64(quantity4) * 0.75
    got4 := calculateDiscount(price4, quantity4)
    if got4 == expected4 {
        fmt.Println("PASS: Quantity 50 - 25%% discount applied (boundary)")
    } else {
        fmt.Printf("FAIL: Quantity 50 - 25%% discount applied (boundary) (expected: %.2f, got: %.2f)\n", expected4, got4)
    }
}