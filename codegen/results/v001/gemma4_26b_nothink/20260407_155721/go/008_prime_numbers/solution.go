package main

import (
	"fmt"
	"os"
	"strconv"
)

// sieveOfEratosthenes returns a slice of booleans where true indicates
// the index is a prime number.
func sieveOfEratosthenes(n int) []bool {
	if n < 2 {
		return []bool{}
	}

	// Create a boolean array "isPrime[0..n]" and initialize
	// all entries it as true. A value in isPrime[i] will
	// finally be false if i is Not a prime, else true.
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	for p := 2; p*p <= n; p++ {
		// If isPrime[p] is not changed, then it is a prime
		if isPrime[p] {
			// Update all multiples of p greater than or equal to the square of it
			// numbers which are multiple of p and are smaller than p^2
			// are already been marked.
			for i := p * p; i <= n; i += p {
				isPrime[i] = false
			}
		}
	}

	return isPrime
}

func main() {
	// Check if an argument is provided
	if len(os.Args) < 2 {
		return
	}

	// Parse the argument from string to integer
	input := os.Args[1]
	n, err := strconv.Atoi(input)
	if err != nil {
		// If the input is not a valid integer, we exit silently or could print error
		return
	}

	// If N is less than 2, there are no primes to print
	if n < 2 {
		return
	}

	// Generate primes using the Sieve of Eratosthenes algorithm
	primesMask := sieveOfEratosthenes(n)

	// Iterate through the mask and print indices that are true
	for i := 2; i <= n; i++ {
		if primesMask[i] {
			fmt.Println(i)
		}
	}
}