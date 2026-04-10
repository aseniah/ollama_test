package main

import (
    "fmt"
    "os"
)

// Function to check if a number is prime
func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i <= int(math.Sqrt(float64(n))) + 1; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

// Main function
func main() {
    // Check if the correct number of arguments are provided
    if len(os.Args) != 2 {
        fmt.Println("Usage: go main.go <integer>")
        return
    }

    // Parse the command-line argument
    N, _ := strconv.Atoi(os.Args[1])

    // Iterate through numbers from 2 to N and print primes
    for i := 2; i <= N; i++ {
        if isPrime(i) {
            fmt.Println(i)
        }
    }
}