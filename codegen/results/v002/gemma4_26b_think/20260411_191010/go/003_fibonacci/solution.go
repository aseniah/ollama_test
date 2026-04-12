package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if an argument was provided
	if len(os.Args) < 2 {
		return
	}

	// Parse the command line argument N to an integer
	nStr := os.Args[1]
	n, err := strconv.Atoi(nStr)
	if err != nil {
		// If the argument is not a valid integer, we stop execution.
		return
	}

	// If N < 1, the requirement specifies to print nothing.
	if n < 1 {
		return
	}

	// We start the sequence with 1, 1, 2, 3, 5...
	// We maintain two variables: 'a' is the current number to print,
	// and 'b' is the next number in the sequence.
	a, b := 1, 1

	for a <= n {
		// Print the current Fibonacci number
		fmt.Println(a)

		// Calculate the next Fibonacci number in the sequence
		// The new 'a' becomes the old 'b'
		// The new 'b' becomes the sum of the old 'a' and 'b'
		next := a + b
		a = b
		b = next
	}
}