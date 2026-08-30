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

	// Primes start at 2, so if N < 2 there are no primes to print.
	if n < 2 {
		return
	}

	// Sieve of Eratosthenes
	sieve := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		sieve[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if sieve[p] {
			for i := p * p; i <= n; i += p {
				sieve[i] = false
			}
		}
	}

	// Print all numbers marked as prime
	for p := 2; p <= n; p++ {
		if sieve[p] {
			fmt.Println(p)
		}
	}
}