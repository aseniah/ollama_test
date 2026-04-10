package main

import (
	"fmt"
	"os"
	"strconv"
)

// getPrimes implements the Sieve of Eratosthenes algorithm.
// It returns a slice of all prime numbers up to and including n.
func getPrimes(n int) []int {
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
			// numbers which are multiples of p and less than p^2 are already marked
			for i := p * p; i <= n; i += p {
				isPrime[i] = false
			}
		}
	}

	// Collect all prime numbers into a slice
	var primes []int
	for p := 2; p <= n; p++ {
		if isPrime[p] {
			primes = append(primes, p)
		}
	}
	return primes
}

func main() {
	// Check if the command line argument is provided
	if len(os.Args) < 2 {
		return
	}

	// Parse the argument from string to integer
	arg := os.Args[1]
	n, err := strconv.Atoi(arg)
	if err != nil {
		// If the argument is not a valid integer, exit silently or handle error
		return
	}

	// Get the list of primes
	primes := getPrimes(n)

	// Print each prime on a new line
	for _, prime := range primes {
		fmt.Println(prime)
	}
}