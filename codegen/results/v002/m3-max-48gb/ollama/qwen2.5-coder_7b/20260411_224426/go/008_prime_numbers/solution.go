package main

import (
	"fmt"
	"os"
	"strconv"
)

// Helper function to check if a number is prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Check if the command line argument is provided and convert it to an integer
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run program.go <integer>")
		os.Exit(1)
	}

	nStr := os.Args[1]
	n, err := strconv.Atoi(nStr)
	if err != nil {
		fmt.Println("Please provide a valid integer.")
		os.Exit(1)
	}

	// Find and print all prime numbers up to and including N
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}