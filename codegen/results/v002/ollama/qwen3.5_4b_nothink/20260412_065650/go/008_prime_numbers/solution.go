package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if an argument was provided
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run main.go <N>")
		os.Exit(1)
	}

	n, err := os.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: argument must be an integer")
		os.Exit(1)
	}

	// Handle edge cases
	if n < 2 {
		return
	}

	isPrime := func(num int) bool {
		if num < 2 {
			return false
		}
		if num == 2 {
			return true
		}
		if num%2 == 0 {
			return false
		}
		for i := 3; i <= int(float64(num)/float64(2)); i += 2 {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}