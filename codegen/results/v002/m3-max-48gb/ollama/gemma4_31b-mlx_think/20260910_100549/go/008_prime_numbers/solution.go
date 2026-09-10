package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if the integer argument N is provided
	if len(os.Args) < 2 {
		return
	}

	// Parse the command line argument to an integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If the input is not a valid integer, we exit silently or could print an error.
		// Based on the prompt, we simply ensure it behaves correctly for the task.
		return
	}

	// Primes start at 2. If N < 2, there are no primes to print.
	if n < 2 {
		return
	}

	// Use the Sieve of Eratosthenes to find all primes up to N
	// Create a boolean slice "isPrime[0..n]" and initialize all entries it as true.
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	for p := 2; p*p <= n; p++ {
		// If isPrime[p] is not changed, then it is a prime
		if isPrime[p] == true {
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