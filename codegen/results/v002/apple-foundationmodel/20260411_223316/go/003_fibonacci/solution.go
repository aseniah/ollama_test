package main

import (
    "fmt"
    "os"
)

func main() {
    // Check if a valid integer argument N is provided
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run fibonacci.go <N>")
        return
    }

    // Parse the command-line argument N
    var n int
    n, _ = os.Atoi(os.Args[1])

    // Ensure N is a valid non-negative integer
    if n < 1 {
        fmt.Println("N must be a non-negative integer")
        return
    }

    // Initialize the first two Fibonacci numbers
    a, b := 1, 1

    // Print Fibonacci numbers up to and including the largest number not exceeding N
    for i := 0; i <= n; i++ {
        fmt.Print(a, " ")
        // Calculate the next Fibonacci number
        a, b = b, a + b
    }

    fmt.Println() // Newline after printing all Fibonacci numbers
}