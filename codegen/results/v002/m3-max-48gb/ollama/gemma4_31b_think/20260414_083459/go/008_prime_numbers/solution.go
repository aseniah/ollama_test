package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Ensure a command line argument is provided
	if len(os.Args) < 2 {
		return
	}

	// Convert the argument to an integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If input is not a valid integer, we treat it as no primes found
		return
	}

	// Primes start from 2. If N is less than 2, there are no primes.
	if n < 2 {
		return
	}

	// Use the Sieve of Eratosthenes algorithm to find all primes up to N
	// Create a boolean slice "isPrime[0..n]" and initialize all entries it as true.
	// A value in isPrime[i] will finally be false if i is Not a prime, else true.
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	for p := 2; p*p <= n; p++ {
		// If isPrime[p] is not changed, then it is a prime
		if isPrime[p] {
			// Update all multiples of p starting from p*p
			for i := p * p; i <= n; i += p {
				isPrime[i] = false
			}
		}
	}

	// Print all prime numbers
	for p := 2; p <= n; p++ {
		if isPrime[p] {
			fmt.Println(p)
		}
	}
}