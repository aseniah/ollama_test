package main

import (
	"fmt"
	"os"
	"strconv"
)

// main implements the logic to find and print all prime numbers up to N.
func main() {
	// Check if an argument was provided via command line.
	if len(os.Args) < 2 {
		return
	}

	// Parse the first argument as an integer N.
	nStr := os.Args[1]
	n, err := strconv.Atoi(nStr)
	if err != nil {
		// If the argument is not a valid integer, we exit silently.
		return
	}

	// If N is less than 2, there are no prime numbers to print.
	if n < 2 {
		return
	}

	// We use the Sieve of Eratosthenes algorithm to find all primes up to N.
	// isPrime[i] will be true if i is a prime number.
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	// Start sieving from the first prime number, 2.
	for p := 2; p*p <= n; p++ {
		// If isPrime[p] is not changed, then it is a prime.
		if isPrime[p] {
			// Update all multiples of p starting from p*p.
			// Multiples of p less than p*p have already been marked.
			for i := p * p; i <= n; i += p {
				isPrime[i] = false
			}
		}
	}

	// Iterate through the boolean slice and print all indices marked as true.
	for p := 2; p <= n; p++ {
		if isPrime[p] {
			fmt.Println(p)
		}
	}
}