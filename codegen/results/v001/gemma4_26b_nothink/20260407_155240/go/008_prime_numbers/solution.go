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
	// We only need to check up to the square root of n.
	// We use i*i <= n to avoid importing math for Sqrt.
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Check if an argument was provided.
	if len(os.Args) < 2 {
		return
	}

	// Parse the argument from string to integer.
	input := os.Args[1]
	n, err := strconv.Atoi(input)
	if err != nil {
		// If the input is not a valid integer, we exit silently or could print an error.
		return
	}

	// Iterate through all numbers from 2 up to and including N.
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}