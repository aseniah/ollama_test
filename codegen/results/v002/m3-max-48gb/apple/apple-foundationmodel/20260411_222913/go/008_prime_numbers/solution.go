package main

import (
    "fmt"
    "os"
)

// isPrime checks if a number is prime.
func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    if n <= 3 {
        return true
    }
    if n%2 == 0 || n%3 == 0 {
        return false
    }
    i := 5
    for i*i <= n {
        if n%i == 0 || n%(i+2) == 0 {
            return false
        }
        i += 6
    }
    return true
}

func main() {
    // Read the command line argument N
    if len(os.Args) != 2 {
        fmt.Println("Usage: go main.go <N>")
        return
    }

    // Convert the command line argument to an integer
    N, _ := os.Atoi(os.Args[1])

    // Check if the argument is a valid integer
    if N < 0 {
        fmt.Println("Please provide a non-negative integer.")
        return
    }

    // Print all prime numbers up to and including N
    for i := 2; i <= N; i++ {
        if isPrime(i) {
            fmt.Println(i)
        }
    }
}