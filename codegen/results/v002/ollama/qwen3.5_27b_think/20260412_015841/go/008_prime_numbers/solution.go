package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Ensure there is at least one command line argument
	if len(os.Args) < 2 {
		return
	}

	// Parse the integer N from the first argument
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If the argument is not a valid integer, do nothing
		return
	}

	// There are no prime numbers less than 2
	if n < 2 {
		return
	}

	// Sieve of Eratosthenes implementation
	// isComposite[i] will be true if i is composite, false if prime
	// We need size n+1 to access index n
	isComposite := make([]bool, n+1)

	// Iterate from 2 up to sqrt(n)
	// Use int64 for the multiplication to prevent overflow on 32-bit systems near MaxInt
	for i := 2; int64(i)*int64(i) <= int64(n); i++ {
		// If i is not marked as composite, it is a prime
		if !isComposite[i] {
			// Mark all multiples of i starting from i*i as composite
			for j := i * i; j <= n; j += i {
				isComposite[j] = true
			}
		}
	}

	// Iterate through the slice and print all numbers that are not composite
	for i := 2; i <= n; i++ {
		if !isComposite[i] {
			fmt.Println(i)
		}
	}
}