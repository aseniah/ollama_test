package main

func calculateDiscount(price float64, quantity int) float64 {
    if quantity < 10 {
        return price * float64(quantity)
    } else if quantity < 50 {
        return price * float64(quantity) * 0.9
    }
    return price * float64(quantity) * 0.75
}

func main() {
    price := 100.0

    // Test case 1: quantity 9 (no discount)
    result := calculateDiscount(price, 9)
    expected := price * 9
    if fmt.Sprintf("%.2f", result) == fmt.Sprintf("%.2f", expected) {
        fmt.Println("PASS: 9 items no discount")
    } else {
        fmt.Printf("FAIL: 9 items (expected: %.2f, got: %.2f)\n", expected, result)
    }

    // Test case 2: quantity 10 (10% discount)
    result = calculateDiscount(price, 10)
    expected = price * 10 * 0.9
    if fmt.Sprintf("%.2f", result) == fmt.Sprintf("%.2f", expected) {
        fmt.Println("PASS: 10 items 10% discount")
    } else {
        fmt.Printf("FAIL: 10 items (expected: %.2f, got: %.2f)\n", expected, result)
    }

    // Test case 3: quantity 49 (10% discount)
    result = calculateDiscount(price, 49)
    expected = price * 49 * 0.9
    if fmt.Sprintf("%.2f", result) == fmt.Sprintf("%.2f", expected) {
        fmt.Println("PASS: 49 items 10% discount")
    } else {
        fmt.Printf("FAIL: 49 items (expected: %.2f, got: %.2f)\n", expected, result)
    }

    // Test case 4: quantity 50 (25% discount)
    result = calculateDiscount(price, 50)
    expected = price * 50 * 0.75
    if fmt.Sprintf("%.2f", result) == fmt.Sprintf("%.2f", expected) {
        fmt.Println("PASS: 50 items 25% discount")
    } else {
        fmt.Printf("FAIL: 50 items (expected: %.2f, got: %.2f)\n", expected, result)
    }
}