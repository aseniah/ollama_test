package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: <program> <N>")
		os.Exit(1)
	}

	n, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil || n < 0 {
		fmt.Fprintln(os.Stderr, "Error: invalid integer input")
		os.Exit(1)
	}

	if n < 2 {
		return
	}

	primes := isPrimeUntil(n)
	for _, p := range primes {
		fmt.Println(p)
	}
}

// isPrimeUntil returns a slice of prime numbers up to and including maxVal.
func isPrimeUntil(maxVal int64) []int64 {
	primes := make([]int64, 0)

	// Special handling for 2
	if maxVal >= 2 {
		primes = append(primes, 2)
	}

	for i := 3; i <= maxVal; i += 2 {
		if isPrime(i) {
			primes = append(primes, int64(i))
		}
	}

	return primes
}

// isPrime checks if a number is prime using trial division up to sqrt(n).
func isPrime(n int64) bool {
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	i := int64(5)
	for ; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}