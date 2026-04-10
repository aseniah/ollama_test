package main

import (
	"fmt"
	"os"
	"strconv"
)

// isPrime checks if a number is prime
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
	
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Check if we have at least one argument
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: provide a single integer argument N")
		os.Exit(1)
	}
	
	// Parse the command line argument
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: argument must be a valid integer")
		os.Exit(1)
	}
	
	// Print all prime numbers up to and including N
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}