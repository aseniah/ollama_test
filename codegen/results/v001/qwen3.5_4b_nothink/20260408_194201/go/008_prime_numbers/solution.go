package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <integer>\n", os.Args[0])
		os.Exit(1)
	}

	nStr := os.Args[1]
	n, err := strconv.Atoi(nStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: expected a valid integer, got '%s'\n", nStr)
		os.Exit(1)
	}

	if n < 2 {
		// No primes less than 2
		return
	}

	// Sieve of Eratosthenes approach to find primes up to n
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	for i := 2; i <= n; i++ {
		if isPrime[i] {
			fmt.Println(i)
		}
	}
}