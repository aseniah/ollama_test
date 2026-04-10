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

func main() {
	// Read the command line argument N
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <N>")
		os.Exit(1)
	}

	N, _ := os.Atoi(os.Args[1])

	// Print all prime numbers up to and including N
	for i := 2; i <= N; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}