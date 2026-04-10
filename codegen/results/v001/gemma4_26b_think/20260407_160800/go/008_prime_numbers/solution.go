package main

import (
	"fmt"
	"os"
	"strconv"
)

// main is the entry point of the program.
func main() {
	// Check if an argument was provided.
	if len(os.Args) < 2 {
		return
	}

	// Parse the command line argument to an integer.
	nStr := os.Args[1]
	n, err := strconv.Atoi(nStr)
	if err != nil {
		// If the argument is not a valid integer, we exit.
		return
	}

	// If N is less than 2, there are no prime numbers.
	if n < 2 {
		return
	}

	// Use the Sieve of Eratosthenes algorithm to find primes up to N.
	// Create a boolean slice "isPrime[0..n]" and initialize
	// all entries it as true. A value in isPrime[i] will
	// finally be false if i is Not a prime, else true.
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	for p := 2; p*p <= n; p++ {
		// If isPrime[p] is not changed, then it is a prime.
		if isPrime[p] {
			// Update all multiples of p starting from p*p.
			// Numbers smaller than p*p have already been marked.
			for i := p * p; i <= n; i += p {
				isPrime[i] = false
			}
		}
	}

	// Print all prime numbers found.
	for p := 2; p <= n; p++ {
		if isPrime[p] {
			fmt.Println(p)
		}
	}
}