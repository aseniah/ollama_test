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
    // Test case 1: quantity = 9 (boundary: below 10, no discount)
    test1Price := 20.0
    test1Quantity := 9
    expected1 := test1Price * float64(test1Quantity)
    got1 := calculateDiscount(test1Price, test1Quantity)
    if fmt.Sprintf("%.2f", got1) == fmt.Sprintf("%.2f", expected1) {
        fmt.Printf("PASS: quantity=9 (no discount applied)\n")
    } else {
        fmt.Printf("FAIL: quantity=9 (expected: %.2f, got: %.2f)\n", expected1, got1)
    }

    // Test case 2: quantity = 10 (boundary: enter 10-49 range, 10% discount)
    test2Price := 30.0
    test2Quantity := 10
    expected2 := test2Price * float64(test2Quantity) * 0.9
    got2 := calculateDiscount(test2Price, test2Quantity)
    if fmt.Sprintf("%.2f", got2) == fmt.Sprintf("%.2f", expected2) {
        fmt.Printf("PASS: quantity=10 (10% discount applied)\n")
    } else {
        fmt.Printf("FAIL: quantity=10 (expected: %.2f, got: %.2f)\n", expected2, got2)
    }

    // Test case 3: quantity = 49 (boundary: top of 10-49 range, 10% discount)
    test3Price := 50.0
    test3Quantity := 49
    expected3 := test3Price * float64(test3Quantity) * 0.9
    got3 := calculateDiscount(test3Price, test3Quantity)
    if fmt.Sprintf("%.2f", got3) == fmt.Sprintf("%.2f", expected3) {
        fmt.Printf("PASS: quantity=49 (10% discount applied)\n")
    } else {
        fmt.Printf("FAIL: quantity=49 (expected: %.2f, got: %.2f)\n", expected3, got3)
    }

    // Test case 4: quantity = 50 (boundary: enter 50+ range, 25% discount)
    test4Price := 10.0
    test4Quantity := 50
    expected4 := test4Price * float64(test4Quantity) * 0.75
    got4 := calculateDiscount(test4Price, test4Quantity)
    if fmt.Sprintf("%.2f", got4) == fmt.Sprintf("%.2f", expected4) {
        fmt.Printf("PASS: quantity=50 (25% discount applied)\n")
    } else {
        fmt.Printf("FAIL: quantity=50 (expected: %.2f, got: %.2f)\n", expected4, got4)
    }
}