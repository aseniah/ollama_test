package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run . <number>")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 0 {
		fmt.Fprintln(os.Stderr, "Error: Please provide a non-negative integer")
		os.Exit(1)
	}

	// Handle numbers below 2 where no primes exist
	if n < 2 {
		return
	}

	// Use Sieve of Eratosthenes to find primes up to N
	initialized := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		initialized[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if initialized[i] {
			for j := i * i; j <= n; j += i {
				initialized[j] = false
			}
		}
	}

	for i := 2; i <= n; i++ {
		if initialized[i] {
			fmt.Println(i)
		}
	}
}