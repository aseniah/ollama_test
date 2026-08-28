package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if an argument was provided
	if len(os.Args) < 2 {
		// If no argument, print nothing as per requirement (or could error, but spec says handle input)
		return
	}

	// Parse the integer argument
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If the argument is not a valid integer, print nothing and exit
		return
	}

	// If N is less than 2, there are no primes
	if n < 2 {
		return
	}

	// Sieve of Eratosthenes to find primes up to N
	// isPrime[i] will be true if i is prime, false otherwise
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	// 0 and 1 are not prime (already false by default, but explicit logic helps)
	isPrime[0] = false
	isPrime[1] = false

	for p := 2; p*p <= n; p++ {
		// If isPrime[p] is not changed, then it is a prime
		if isPrime[p] {
			// Update all multiples of p
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