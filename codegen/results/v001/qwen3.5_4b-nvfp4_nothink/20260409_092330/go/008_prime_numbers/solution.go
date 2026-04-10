package main

import (
	"fmt"
	"os"
)

// isPrime checks if a number n is prime.
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Get the single integer argument N from command line
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid argument: %v\n", err)
		os.Exit(1)
	}

	// Iterate and print primes up to N
	var primes []int
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	// If there are primes, print each one on a new line
	for _, p := range primes {
		fmt.Println(p)
	}
}