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

	// Parse the argument from string to integer.
	arg := os.Args[1]
	n, err := strconv.Atoi(arg)
	if err != nil {
		// If the argument is not a valid integer, we exit.
		return
	}

	// If N is less than 2, there are no prime numbers to print.
	if n < 2 {
		return
	}

	// We use the Sieve of Eratosthenes algorithm to find all primes up to N.
	// Create a boolean slice "isPrime" and initialize all entries to true.
	// A value in isPrime[i] will finally be false if i is Not a prime, else true.
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	// The sieve logic:
	// Start from the first prime number, 2.
	for p := 2; p*p <= n; p++ {
		// If isPrime[p] is not changed, then it is a prime.
		if isPrime[p] {
			// Update all multiples of p starting from p*p.
			// We start at p*p because smaller multiples would have already been marked.
			for i := p * p; i <= n; i += p {
				isPrime[i] = false
			}
		}
	}

	// Iterate through the slice and print all indices marked as true.
	for p := 2; p <= n; p++ {
		if isPrime[p] {
			fmt.Println(p)
		}
	}
}