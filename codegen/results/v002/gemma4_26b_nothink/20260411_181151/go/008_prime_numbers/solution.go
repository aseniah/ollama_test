package main

import (
	"fmt"
	"os"
	"strconv"
)

// sieveOfEratosthenes returns a slice of all prime numbers up to N.
func sieveOfEratosthenes(n int) []int {
	if n < 2 {
		return []int{}
	}

	// Create a boolean slice "isPrime[0..n]" and initialize
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
			// numbers which are multiples of p and are less than p^2
			// are already been marked.
			for i := p * p; i <= n; i += p {
				isPrime[i] = false
			}
		}
	}

	// Collect all prime numbers into a slice
	primes := []int{}
	for p := 2; p <= n; p++ {
		if isPrime[p] {
			primes = append(primes, p)
		}
	}
	return primes
}

func main() {
	// Check if an argument was provided
	if len(os.Args) < 2 {
		return
	}

	// Parse the argument from string to integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If the argument is not a valid integer, exit silently or handle error
		return
	}

	// Get the list of primes up to N
	primes := sieveOfEratosthenes(n)

	// Print each prime number on a new line
	for _, prime := range primes {
		fmt.Println(prime)
	}
}