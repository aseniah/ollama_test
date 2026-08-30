package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if an argument was provided
	if len(os.Args) < 2 {
		return // Or handle error as needed, but spec implies accepting a single arg
	}

	// Parse the integer N from the command line
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid integer argument\n")
		os.Exit(1)
	}

	// If N is less than 2, there are no primes to print
	if n < 2 {
		return
	}

	// Use Sieve of Eratosthenes to find primes up to N
	primes := sieve(n)

	// Print each prime on a new line
	for _, p := range primes {
		fmt.Println(p)
	}
}

// sieve generates a slice of all prime numbers up to and including limit
func sieve(limit int) []int {
	// isPrime[i] will be true if i is prime, false otherwise
	// Initialize all entries as true
	isPrime := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}

	// Iterate from 2 up to sqrt(limit)
	for p := 2; p*p <= limit; p++ {
		if isPrime[p] {
			// Mark all multiples of p starting from p*p as not prime
			for multiple := p * p; multiple <= limit; multiple += p {
				isPrime[multiple] = false
			}
		}
	}

	// Collect the primes into a slice
	var result []int
	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			result = append(result, i)
		}
	}

	return result
}