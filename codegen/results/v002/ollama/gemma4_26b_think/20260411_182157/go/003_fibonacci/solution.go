package main

import (
	"fmt"
	"os"
	"strconv"
)

/**
 * The program accepts a single integer argument N from the command line.
 * It prints the Fibonacci sequence (1, 1, 2, 3, 5, ...) up to the largest
 * Fibonacci number that does not exceed N.
 * If N < 1, the program prints nothing.
 */

func main() {
	// Check if an argument was provided via the command line.
	// os.Args[0] is the program name, so we look for os.Args[1].
	if len(os.Args) < 2 {
		return
	}

	// Parse the first command-line argument into an integer.
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If the argument is not a valid integer, terminate silently.
		return
	}

	// As per requirements: If N < 1, print nothing.
	if n < 1 {
		return
	}

	// The Fibonacci sequence begins with 1, 1, 2, 3, 5...
	// Since we have already verified N >= 1, we can safely print the first '1'.
	fmt.Println(1)

	// Since N is at least 1, the second '1' in the sequence (1, 1, 2...) 
	// is also guaranteed to be less than or equal to N.
	fmt.Println(1)

	// Initialize the Fibonacci calculation variables.
	// 'prev' and 'curr' represent the two most recent numbers in the sequence.
	prev := 1
	curr := 1
	// 'next' represents the next number to be calculated and checked against N.
	next := prev + curr

	// Loop as long as the next number in the sequence does not exceed N.
	for next <= n {
		// Print the number that satisfies the condition.
		fmt.Println(next)

		// Advance the sequence:
		// The current number becomes the 'previous' number.
		// The 'next' number becomes the 'current' number.
		prev = curr
		curr = next

		// Calculate the new next number in the sequence.
		next = prev + curr
	}
}