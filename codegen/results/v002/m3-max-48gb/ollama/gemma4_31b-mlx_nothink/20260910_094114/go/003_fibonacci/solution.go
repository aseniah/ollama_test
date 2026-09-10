package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if the command line argument is provided
	if len(os.Args) < 2 {
		return
	}

	// Parse the argument N as an integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	// If N < 1, print nothing as per requirements
	if n < 1 {
		return
	}

	// Initialize the first two Fibonacci numbers
	a, b := 1, 1

	// The sequence starts: 1, 1, 2, 3, 5...
	// We print the first 1 since N >= 1
	fmt.Println(a)

	// Calculate and print subsequent numbers until we exceed N
	for b <= n {
		fmt.Println(b)
		
		// Compute next Fibonacci number
		next := a + b
		a = b
		b = next
	}
}