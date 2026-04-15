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

	// Parse the input argument N as an integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If input is not a valid integer, we simply exit
		return
	}

	// Requirement: If N < 1, print nothing.
	if n < 1 {
		return
	}

	// Start the sequence: 1, 1, 2, 3, 5, ...
	// We use a and b to keep track of the current and next numbers in the sequence.
	a := 1
	b := 1

	for a <= n {
		// Print the current Fibonacci number
		fmt.Println(a)

		// Calculate the next Fibonacci number
		// Sequence: 1, 1, 2, 3, 5...
		// Step 1: a=1, b=1 -> Print 1, next a=1, b=2
		// Step 2: a=1, b=2 -> Print 1, next a=2, b=3
		// Step 3: a=2, b=3 -> Print 2, next a=3, b=5
		a, b = b, a+b
	}
}