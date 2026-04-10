package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Get the command line argument N
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Error: please provide an integer argument N")
		os.Exit(1)
	}

	// Parse the argument
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: argument must be an integer")
		os.Exit(1)
	}

	// If N < 1, print nothing
	if n < 1 {
		return
	}

	// Generate Fibonacci numbers
	a := 1
	b := 1

	// Print first Fibonacci number if it's within range
	if a <= n {
		fmt.Println(a)
	}

	// Generate and print subsequent Fibonacci numbers
	for b <= n {
		fmt.Println(b)
		// Generate next number in sequence
		next := a + b
		a = b
		b = next
	}
}