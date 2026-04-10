package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if command line argument was provided
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: program N")
		os.Exit(1)
	}

	// Parse the integer argument
	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: N must be an integer")
		os.Exit(1)
	}

	if N < 2 {
		return
	}

	// Sieve of Eratosthenes to find primes up to N
	isPrime := make([]bool, N+1)
	for i := 2; i <= N; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= N; i++ {
		if isPrime[i] {
			for j := i * i; j <= N; j += i {
				isPrime[j] = false
			}
		}
	}

	// Print all primes
	for i := 2; i <= N; i++ {
		if isPrime[i] {
			fmt.Println(i)
		}
	}
}