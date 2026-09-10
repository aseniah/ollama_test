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
	// Check odd numbers from 3 up to the square root of n
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Ensure an argument is provided
	if len(os.Args) < 2 {
		return
	}

	// Convert the first command line argument to an integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If the argument is not a valid integer, exit silently or handle error
		return
	}

	// Iterate from 2 up to and including N
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}