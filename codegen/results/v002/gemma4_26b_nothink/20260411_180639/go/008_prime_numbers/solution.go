package main

import (
	"fmt"
	"os"
	"strconv"
)

// isPrime checks if a number is prime using the trial division method.
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	// Optimization: Check divisors up to the square root of n.
	// All primes greater than 3 are of the form 6k ± 1.
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Ensure there is exactly one command line argument.
	if len(os.Args) != 2 {
		// If the user provides wrong arguments, we exit silently or could print usage.
		// The prompt implies a specific task, so we handle the input expectation.
		return
	}

	// Parse the argument as an integer.
	input := os.Args[1]
	n, err := strconv.Atoi(input)
	if err != nil {
		// If the input is not a valid integer, we exit.
		return
	}

	// Iterate from 2 up to and including N.
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}