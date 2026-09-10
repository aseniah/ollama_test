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
	// Use 6k +/- 1 rule to optimize primality test
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Check if the command line argument N is provided
	if len(os.Args) < 2 {
		return
	}

	// Parse the argument to an integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If input is not a valid integer, we exit silently or could print an error.
		// Based on the prompt requirements, we focus on printing primes.
		return
	}

	// Iterate from 2 up to N and print if the number is prime
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}