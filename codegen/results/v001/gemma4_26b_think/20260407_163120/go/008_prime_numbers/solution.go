package main

import (
	"fmt"
	"os"
	"strconv"
)

// main implements the logic to find and print all prime numbers up to N.
func main() {
	// Check if an argument was provided via command line
	if len(os.Args) < 2 {
		return
	}

	// Parse the first argument from string to integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If the argument is not a valid integer, we exit silently
		return
	}

	// If N is less than 2, there are no prime numbers to print
	if n < 2 {
		return
	}

	// Use the Sieve of Eratosthenes algorithm to find primes up to N.
	// isPrime[i] will be true if i is a prime number.
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	// Iterate from the first prime (2) up to the square root of N
	for p := 2; p*p <= n; p++ {
		// If isPrime[p] is not changed, then it is a prime
		if isPrime[p] {
			// Update all multiples of p starting from p*p
			// Multiples smaller than p*p have already been marked by smaller primes
			for i := p * p; i <= n; i += p {
				isPrime[i] = false
			}
		}
	}

	// Iterate through the boolean slice and print the indices that are marked true
	for p := 2; p <= n; p++ {
		if isPrime[p] {
			fmt.Println(p)
		}
	}
}