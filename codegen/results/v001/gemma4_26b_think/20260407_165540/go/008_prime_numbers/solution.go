package main

import (
	"fmt"
	"os"
	"strconv"
)

// main implements the logic to find and print primes up to N using the Sieve of Eratosthenes.
func main() {
	// Check if an argument was provided
	if len(os.Args) != 2 {
		return
	}

	// Parse the input argument to an integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	// If N is less than 2, there are no prime numbers to print
	if n < 2 {
		return
	}

	// Use the Sieve of Eratosthenes algorithm
	// Create a slice to track prime status for numbers up to N
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	// Iterate from 2 up to the square root of N
	for p := 2; p*p <= n; p++ {
		// If isPrime[p] is not changed, then it is a prime
		if isPrime[p] {
			// Update all multiples of p starting from p*p
			for i := p * p; i <= n; i += p {
				isPrime[i] = false
			}
		}
	}

	// Print all indices that are still marked as true
	for p := 2; p <= n; p++ {
		if isPrime[p] {
			fmt.Println(p)
		}
	}
}