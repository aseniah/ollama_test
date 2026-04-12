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
    // Test case 1: Quantity 9
    price := 10.0
    quantity := 9
    expected := price * float64(quantity)
    result := calculateDiscount(price, quantity)
    if result == expected {
        fmt.Printf("PASS: Test case 1 - Quantity 9\n")
    } else {
        fmt.Printf("FAIL: Test case 1 - Quantity 9 expected: %f, got: %f\n", expected, result)
    }

    // Test case 2: Quantity 10
    price := 10.0
    quantity := 10
    expected := price * float64(quantity)
    result := calculateDiscount(price, quantity)
    if result == expected {
        fmt.Printf("PASS: Test case 2 - Quantity 10\n")
    } else {
        fmt.Printf("FAIL: Test case 2 - Quantity 10 expected: %f, got: %f\n", expected, result)
    }

    // Test case 3: Quantity 49
    price := 10.0
    quantity := 49
    expected := price * float64(quantity) * 0.9
    result := calculateDiscount(price, quantity)
    if result == expected {
        fmt.Printf("PASS: Test case 3 - Quantity 49\n")
    } else {
        fmt.Printf("FAIL: Test case 3 - Quantity 49 expected: %f, got: %f\n", expected, result)
    }

    // Test case 4: Quantity 50
    price := 10.0
    quantity := 50
    expected := price * float64(quantity) * 0.75
    result := calculateDiscount(price, quantity)
    if result == expected {
        fmt.Printf("PASS: Test case 4 - Quantity 50\n")
    } else {
        fmt.Printf("FAIL: Test case 4 - Quantity 50 expected: %f, got: %f\n", expected, result)
    }
}