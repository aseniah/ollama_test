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
    fmt.Println("=== Testing calculateDiscount Function ===")
    fmt.Println()

    // Test Case 1: quantity = 9 (boundary < 10)
    test1 := TestCase(calculateDiscount, 100.0, 9, 900.0, "quantity = 9 (below threshold)")
    fmt.Println(test1)
    fmt.Println()

    // Test Case 2: quantity = 10 (boundary >= 10, < 50)
    test2 := TestCase(calculateDiscount, 100.0, 10, 900.0, "quantity = 10 (at threshold)")
    fmt.Println(test2)
    fmt.Println()

    // Test Case 3: quantity = 49 (boundary < 50)
    test3 := TestCase(calculateDiscount, 100.0, 49, 441.0, "quantity = 49 (at threshold)")
    fmt.Println(test3)
    fmt.Println()

    // Test Case 4: quantity = 50 (boundary >= 50)
    test4 := TestCase(calculateDiscount, 100.0, 50, 750.0, "quantity = 50 (at threshold)")
    fmt.Println(test4)

    // Additional edge case: quantity = 1 (minimum valid quantity)
    test5 := TestCase(calculateDiscount, 100.0, 1, 100.0, "quantity = 1 (minimum quantity)")
    fmt.Println(test5)
}

// TestCase helper function to run a calculation and compare results
func TestCase(c func(float64, int) float64, price, qty float64, expected float64, description string) string {
    result := c(price, int(qty))
    var passed bool
    if fmt.Sprint(result) == fmt.Sprint(expected) {
        passed = true
    } else {
        // For non-exact float comparisons due to floating point precision
        diff := abs(result - expected)
        passed = diff < 0.01
    }

    if passed {
        return fmt.Sprintf("PASS: %s", description)
    } else {
        return fmt.Sprintf("FAIL: %s (expected: %.2f, got: %.2f)", description, expected, result)
    }
}

// abs returns the absolute value of a float64
func abs(x float64) float64 {
    if x < 0 {
        return -x
    }
    return x
}