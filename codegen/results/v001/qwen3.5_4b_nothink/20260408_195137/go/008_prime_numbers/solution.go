package main

import (
	"fmt"
	"os"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Check if N is provided
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s N\n", os.Args[0])
		os.Exit(1)
	}

	N, err := fmt.Sscanf(os.Args[1], "%d", nil)
	if err != nil || N != 1 {
		fmt.Fprintf(os.Stderr, "Error parsing argument: %s\n", os.Args[1])
		os.Exit(1)
	}

	// Validate input range
	if N < 2 {
		// No primes below 2
		return
	}

	// Collect primes
	var primes []int
	for i := 2; i <= N; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	// Output primes
	for _, p := range primes {
		fmt.Println(p)
	}
}