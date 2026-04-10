package main

import (
	"fmt"
	"os"
)

// isPrime checks if a number is prime.
func isPrime(n int) bool {
	if n <= 1 {
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

// printPrimes prints all prime numbers up to and including N.
func printPrimes(N int) {
	for i := 2; i <= N; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: go run primes.go <integer>")
		os.Exit(1)
	}

	N, err := strconv.Atoi(args[0])
	if err != nil || N <= 0 {
		fmt.Println("Please provide a positive integer as an argument.")
		os.Exit(1)
	}

	printPrimes(N)
}