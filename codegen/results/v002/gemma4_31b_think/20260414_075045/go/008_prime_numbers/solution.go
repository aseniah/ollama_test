package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if an argument was provided
	if len(os.Args) < 2 {
		return
	}

	// Parse the argument N as an integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If the input is not a valid integer, we stop execution
		return
	}

	// Prime numbers are defined as natural numbers greater than 1
	if n < 2 {
		return
	}

	// Use the Sieve of Eratosthenes to find all primes up to N
	// Create a slice to track primality, initialized to true
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	// Iterate through the slice and mark multiples of each prime as non-prime
	for p := 2; p*p <= n; p++ {
		if isPrime[p] {
			// Start marking from p*p, as smaller multiples have already been marked
			for i := p * p; i <= n; i += p {
				isPrime[i] = false
			}
		}
	}

	// Print all indices that remained true
	for p := 2; p <= n; p++ {
		if isPrime[p] {
			fmt.Println(p)
		}
	}
}