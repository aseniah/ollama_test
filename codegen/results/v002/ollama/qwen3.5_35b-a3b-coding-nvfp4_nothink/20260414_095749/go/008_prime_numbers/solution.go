package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	if n < 2 {
		// No primes up to N
		return
	}

	// Sieve of Eratosthenes to find primes up to n
	// isPrime[i] will be true if i is prime, false otherwise
	limit := n + 1
	isPrime := make([]bool, limit)
	for i := 2; i < limit; i++ {
		isPrime[i] = true
	}

	for p := 2; p*p < limit; p++ {
		if isPrime[p] {
			// Mark multiples of p as non-prime
			for multiple := p * p; multiple < limit; multiple += p {
				isPrime[multiple] = false
			}
		}
	}

	// Print all primes up to n
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			fmt.Println(i)
		}
	}
}