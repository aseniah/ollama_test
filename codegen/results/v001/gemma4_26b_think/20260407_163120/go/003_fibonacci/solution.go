package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if an argument was provided via the command line
	if len(os.Args) < 2 {
		return
	}

	// Convert the first command-line argument to an integer N
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If the argument is not a valid integer, exit quietly
		return
	}

	// If N < 1, per requirements, print nothing
	if n < 1 {
		return
	}

	// The sequence starts with 1, 1, ...
	// Since we have already verified N >= 1, the first two 1s are always <= N.
	fmt.Println(1)
	fmt.Println(1)

	// Initialize the first two terms of the Fibonacci sequence
	a := 1
	b := 1

	// Iterate to calculate subsequent Fibonacci numbers
	for {
		next := a + b

		// If the next Fibonacci number exceeds N, we stop the sequence
		if next > n {
			break
		}

		// Print the current Fibonacci number
		fmt.Println(next)

		// Update variables for the next iteration of the sequence
		a = b
		b = next
	}
}