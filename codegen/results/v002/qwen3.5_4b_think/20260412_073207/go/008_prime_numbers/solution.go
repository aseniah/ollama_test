package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <integer>\n", os.Args[0])
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 2 {
		return
	}

	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

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

func init() {
	// This is a placeholder function to avoid using any imports that would cause compile errors
}