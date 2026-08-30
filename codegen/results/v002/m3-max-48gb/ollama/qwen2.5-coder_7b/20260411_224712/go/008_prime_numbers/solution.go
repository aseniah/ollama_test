package main

import (
	"fmt"
	"os"
	"strconv"
)

// isPrime checks if a number n is prime.
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Check if the correct number of arguments is provided
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run program.go <integer>")
		os.Exit(1)
	}

	// Parse the command line argument to an integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 0 {
		fmt.Println("Please provide a non-negative integer as an argument.")
		os.Exit(1)
	}

	// Print all prime numbers up to and including N
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}