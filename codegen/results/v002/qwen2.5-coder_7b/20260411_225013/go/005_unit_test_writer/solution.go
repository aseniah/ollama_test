package main

import (
    "fmt"
)

func calculateDiscount(price float64, quantity int) float64 {
    if quantity < 10 {
        return price * float64(quantity)
    } else if quantity < 50 {
        return price * float64(quantity) * 0.9
    }
    return price * float64(quantity) * 0.75
}

func testCalculateDiscount() {
    tests := []struct {
        description string
        price       float64
        quantity    int
        expected    float64
    }{
        {"quantity < 10", 10, 9, 90},
        {"quantity == 10", 10, 10, 100},
        {"quantity < 50", 10, 49, 441},
        {"quantity == 50", 10, 50, 375},
    }

    for _, test := range tests {
        got := calculateDiscount(test.price, test.quantity)
        if got != test.expected {
            fmt.Printf("FAIL: %s (expected: %v, got: %v)\n", test.description, test.expected, got)
        } else {
            fmt.Printf("PASS: %s\n", test.description)
        }
    }
}

func main() {
    testCalculateDiscount()
}
```

This Go program includes a `calculateDiscount` function that calculates a discount based on the price and quantity of items. The program also contains a `testCalculateDiscount` function that tests the `calculateDiscount` function with different scenarios, including boundary conditions. The results of the tests are printed to stdout in either "PASS" or "FAIL" format, indicating whether each test passed or failed.