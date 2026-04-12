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

// printPrimes prints all prime numbers up to and including N, one per line.
func printPrimes(n int) {
	for num := 1; num <= n; num++ {
		if isPrime(num) {
			fmt.Println(num)
		}
	}
}

func main() {
	var N int
	// Read the single integer argument from command line.
	args := os.Args
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: program <N>")
		os.Exit(1)
	}
	N, err := parseInt(strings.NewReader(args[1]))
	if err != nil || N <= 0 {
		fmt.Fprintln(os.Stderr, "Error: invalid input")
		os.Exit(1)
	}

	printPrimes(N)
}