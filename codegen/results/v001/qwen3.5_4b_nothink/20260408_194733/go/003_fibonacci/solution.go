package main

import (
	"fmt"
	"os"
)

func main() {
	// Accept a single integer argument N from the command line
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run main.go <N>")
		os.Exit(1)
	}

	var n int
	err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing integer argument")
		os.Exit(1)
	}

	// If N < 1, print nothing
	if n < 1 {
		return
	}

	// Generate and print Fibonacci numbers up to and including the largest <= n
	a, b := 1, 1
	for a <= n {
		fmt.Println(a)
		// Calculate next Fibonacci number
		next := a + b
		a = b
		b = next
	}
}