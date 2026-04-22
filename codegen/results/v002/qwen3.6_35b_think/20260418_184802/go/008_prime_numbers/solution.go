package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Ensure exactly one argument is provided
	if len(os.Args) != 2 {
		return
	}

	// Parse the command line argument
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 2 {
		// No primes to print if N < 2 or invalid input
		return
	}

	// Sieve of Eratosthenes to find all primes up to N
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	// Print each prime number on its own line
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			fmt.Println(i)
		}
	}
}