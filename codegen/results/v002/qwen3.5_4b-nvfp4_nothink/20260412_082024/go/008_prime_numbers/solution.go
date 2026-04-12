package main

import (
	"fmt"
	"os"
)

// isPrime returns true if n is a prime number, false otherwise.
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	// Check divisors of form 6k ± 1 up to sqrt(n)
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// main executes the program logic.
func main() {
	// Validate command-line argument
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run . <N>")
		os.Exit(1)
	}

	N, err := parseIntInput(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: invalid input. Please provide a positive integer.")
		os.Exit(1)
	}

	// Print all primes up to N (inclusive)
	for i := 2; i <= N; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

// parseIntInput safely converts the string argument to an int.
func parseIntInput(s string) (int, error) {
	var n int
	scan, err := fmt.Sscanf(s, "%d", &n)
	if err != nil || scan != 1 {
		return 0, fmt.Errorf("invalid integer format: %s", s)
	}
	return n, nil
}