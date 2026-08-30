package main

import (
	"fmt"
	"os"
	"strconv"
)

// isPrime checks if a given integer is a prime number.
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	// We check divisors up to the square root of n.
	// Optimization: check divisors in the form of 6k +/- 1.
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Ensure an argument was provided.
	if len(os.Args) < 2 {
		return
	}

	// Parse the command line argument to an integer.
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If the input is not a valid integer, we simply exit.
		return
	}

	// Iterate from 2 up to N and print primes.
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}