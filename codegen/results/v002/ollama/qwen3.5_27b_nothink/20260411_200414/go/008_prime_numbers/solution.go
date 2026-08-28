package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if exactly one argument is provided
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <N>")
		os.Exit(1)
	}

	// Parse the integer argument
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: argument must be an integer")
		os.Exit(1)
	}

	// Handle cases where N is less than 2 (no primes exist)
	if n < 2 {
		return
	}

	// Sieve of Eratosthenes
	isPrime := make([]bool, n+1)
	
	// Initialize all numbers as prime (true) by default
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	// 0 and 1 are not prime
	isPrime[0] = false
	isPrime[1] = false

	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			// Mark multiples of i as non-prime
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	// Print all primes up to N
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			fmt.Println(i)
		}
	}
}