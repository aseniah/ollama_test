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
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	// Check odd divisors up to the square root of n
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Check if the command line argument is provided
	if len(os.Args) < 2 {
		return
	}

	// Parse the argument to an integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If input is not a valid integer, we simply exit
		return
	}

	// Iterate from 2 up to N and print primes
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}