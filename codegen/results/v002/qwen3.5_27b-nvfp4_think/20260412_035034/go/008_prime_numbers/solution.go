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
	// Check if we have at least one command line argument
	if len(os.Args) < 2 {
		return
	}

	// Parse the integer argument
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	// Print all prime numbers up to and including n
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}