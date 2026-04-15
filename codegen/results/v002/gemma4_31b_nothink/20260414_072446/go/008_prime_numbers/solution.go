package main

import (
	"fmt"
	"os"
	"strconv"
)

// isPrime checks if a number is prime.
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
	// Optimization: check divisors up to the square root of n
	// All primes greater than 3 are of the form 6k ± 1
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Ensure a command line argument was provided
	if len(os.Args) < 2 {
		return
	}

	// Parse the command line argument N to an integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If input is not a valid integer, we exit silently or can print error
		// Based on the prompt, we assume valid integer input or print nothing
		return
	}

	// Iterate through all integers from 2 up to and including N
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}