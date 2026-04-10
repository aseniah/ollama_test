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
    testCases := []struct {
        description  string
        price        float64
        quantity    int
        expected     float64
    }{
        {
            "Basic case with quantity 10",
            10.0,
            10,
            10.0,
        },
        {
            "Basic case with quantity 9",
            10.0,
            9,
            9.0,
        },
        {
            "Basic case with quantity 50",
            10.0,
            50,
            50.0,
        },
        {
            "Basic case with quantity 49",
            10.0,
            49,
            44.5,
        },
    }

    for _, test := range testCases {
        result := calculateDiscount(test.price, test.quantity)
        if result == test.expected {
            fmt.Printf("PASS: %s\n", test.description)
        } else {
            fmt.Printf("FAIL: %s (expected: %.2f, got: %.2f)\n", test.description, test.expected, result)
        }
    }
}