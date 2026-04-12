package main

import (
	"fmt"
	"os"
)

// isPrime checks if a number is prime using the Sieve of Eratosthenes approach internally.
// Since we need to print all primes up to N, we'll implement a sieve to find them.
func getPrimesUpTo(n int) []int {
	if n < 2 {
		return []int{}
	}

	// Create a boolean slice "isPrime[0..n]" and initialize all entries as true.
	// A value in isPrime[i] will finally be false if i is not a prime, else true.
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	// Start with the first prime
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			// Mark all multiples of i as non-prime
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	// Collect the primes into a slice
	var primes []int
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

func main() {
	// Get N from command line arguments
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <N>\n", os.Args[0])
		os.Exit(1)
	}

	n, err := parseInt(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid integer argument: %s\n", os.Args[1])
		os.Exit(1)
	}

	// Ensure N is non-negative
	if n < 0 {
		fmt.Fprintf(os.Stderr, "Error: N must be non-negative\n")
		os.Exit(1)
	}

	// Get all primes up to N
	primes := getPrimesUpTo(n)

	// Print each prime on a new line
	for _, p := range primes {
		fmt.Printf("%d\n", p)
	}
}

// parseInt parses a string to an integer
func parseInt(s string) (int, error) {
	result, err := fmt.Sscanf(s, "%d", &result)
	if err != nil || result != 1 {
		return 0, fmt.Errorf("cannot parse '%s' as integer", s)
	}
	return result, nil
}