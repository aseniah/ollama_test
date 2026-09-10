package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if the user provided at least one command line argument
	if len(os.Args) < 2 {
		return
	}

	// Parse the first argument as an integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If the argument is not a valid integer, we exit silently as per usual CLI behavior 
		// or could print an error. Based on requirements, we just need the primes for a valid N.
		return
	}

	// Primes start from 2. If N < 2, there are no primes to print.
	if n < 2 {
		return
	}

	// Sieve of Eratosthenes algorithm
	// Create a boolean slice to keep track of prime numbers. 
	// Initially, assume all numbers from 0 to n are prime.
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	// Iterate from 2 up to the square root of N
	for p := 2; p*p <= n; p++ {
		// If isPrime[p] is still true, then it is a prime
		if isPrime[p] {
			// Update all multiples of p to be false (not prime)
			// Start from p*p because smaller multiples would have already been marked
			for i := p * p; i <= n; i += p {
				isPrime[i] = false
			}
		}
	}

	// Print all numbers that are still marked as prime
	for p := 2; p <= n; p++ {
		if isPrime[p] {
			fmt.Println(p)
		}
	}
}