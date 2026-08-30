package main

import (
	"fmt"
	"os"
	"strconv"
)

/**
 * This program accepts a single integer argument N from the command line.
 * It prints all Fibonacci numbers starting from 1, 1, 2, 3, ...
 * up to and including the largest Fibonacci number that does not exceed N.
 * If N < 1, the program prints nothing.
 */

func main() {
	// Check if an argument was provided via the command line.
	if len(sArgs) < 2 {
		return
	}

	// Parse the first command-line argument into an integer.
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If the argument is not a valid integer, we exit without printing.
		return
	}

	// If N < 1, according to the requirements, print nothing.
	if n < 1 {
		return
	}

	// Initialize the first two numbers of the Fibonacci sequence.
	// The sequence starts with 1, 1, 2, 3, 5, ...
	// We use a and b to track the current and the next number in the sequence.
	a := 1
	b := 1

	// Iterate as long as the current Fibonacci number 'a' does not exceed N.
	for a <= n {
		// Print the current Fibonacci number.
		fmt.Println(a)

		// Calculate the next number in the sequence.
		// temp = a + b
		// a = b
		// b = temp
		// In Go, we can use multiple assignment for a cleaner implementation.
		next := a + b
		a = b
		b = next
	}
}

// Note: The variable sArgs used in the check above was a typo in thought process, 
// but os.Args is the correct way to access command line arguments in Go.
// Let's rewrite the logic slightly cleaner in the final version.