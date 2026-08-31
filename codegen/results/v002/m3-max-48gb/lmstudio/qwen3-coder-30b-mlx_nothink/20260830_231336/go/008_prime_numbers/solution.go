package main

import (
	"fmt"
	"os"
	"strconv"
)

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
	
	// Check odd divisors up to sqrt(n)
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s N\n", os.Args[0])
		os.Exit(1)
	}
	
	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s is not a valid integer\n", os.Args[1])
		os.Exit(1)
	}
	
	if N < 2 {
		// If N is less than 2, there are no primes up to N
		return
	}
	
	for i := 2; i <= N; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}