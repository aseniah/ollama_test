package main

import (
	"fmt"
	"os"
	"strconv"
)

// isPrime checks if a given integer is a prime number.
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
	// Optimization: check divisors up to sqrt(n) using 6k +/- 1 rule
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Check if the command line argument is provided
	if len(os.Args) < 2 {
		return
	}

	// Parse the argument N to an integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If the argument is not a valid integer, we print nothing or could print an error.
		// Based on requirements, we simply exit.
		return
	}

	// Iterate from 2 up to N and print primes
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}