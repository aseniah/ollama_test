package main

import (
	"fmt"
	"os"
	"strconv"
)

// isPrime checks if a number is prime
func isPrime(n int) bool {
	if n < 2 {
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
	// Check if exactly one argument is provided
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: program <N>")
		os.Exit(1)
	}

	// Parse the integer argument
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Invalid integer argument")
		os.Exit(1)
	}

	// If N < 2, there are no primes up to N
	if n < 2 {
		return
	}

	// Print all primes up to and including N
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}