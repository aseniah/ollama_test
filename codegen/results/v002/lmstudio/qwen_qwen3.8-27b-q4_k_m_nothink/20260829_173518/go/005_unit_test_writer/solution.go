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
    // Test case 1: quantity 9 (below 10, no discount)
    // price = 10.0, quantity = 9
    // expected = 10.0 * 9 = 90.0
    price1 := 10.0
    quantity1 := 9
    expected1 := price1 * float64(quantity1)
    got1 := calculateDiscount(price1, quantity1)
    if fmt.Sprintf("%.2f", expected1) == fmt.Sprintf("%.2f", got1) {
        fmt.Println("PASS: quantity 9 (no discount)")
    } else {
        fmt.Printf("FAIL: quantity 9 (no discount) (expected: %.2f, got: %.2f)\n", expected1, got1)
    }

    // Test case 2: quantity 10 (boundary, 10% discount)
    // price = 10.0, quantity = 10
    // expected = 10.0 * 10 * 0.9 = 90.0
    price2 := 10.0
    quantity2 := 10
    expected2 := price2 * float64(quantity2) * 0.9
    got2 := calculateDiscount(price2, quantity2)
    if fmt.Sprintf("%.2f", expected2) == fmt.Sprintf("%.2f", got2) {
        fmt.Println("PASS: quantity 10 (10%% discount)")
    } else {
        fmt.Printf("FAIL: quantity 10 (10%% discount) (expected: %.2f, got: %.2f)\n", expected2, got2)
    }

    // Test case 3: quantity 49 (just below 50, 10% discount)
    // price = 10.0, quantity = 49
    // expected = 10.0 * 49 * 0.9 = 441.0
    price3 := 10.0
    quantity3 := 49
    expected3 := price3 * float64(quantity3) * 0.9
    got3 := calculateDiscount(price3, quantity3)
    if fmt.Sprintf("%.2f", expected3) == fmt.Sprintf("%.2f", got3) {
        fmt.Println("PASS: quantity 49 (10%% discount)")
    } else {
        fmt.Printf("FAIL: quantity 49 (10%% discount) (expected: %.2f, got: %.2f)\n", expected3, got3)
    }

    // Test case 4: quantity 50 (boundary, 25% discount)
    // price = 10.0, quantity = 50
    // expected = 10.0 * 50 * 0.75 = 375.0
    price4 := 10.0
    quantity4 := 50
    expected4 := price4 * float64(quantity4) * 0.75
    got4 := calculateDiscount(price4, quantity4)
    if fmt.Sprintf("%.2f", expected4) == fmt.Sprintf("%.2f", got4) {
        fmt.Println("PASS: quantity 50 (25%% discount)")
    } else {
        fmt.Printf("FAIL: quantity 50 (25%% discount) (expected: %.2f, got: %.2f)\n", expected4, got4)
    }
}