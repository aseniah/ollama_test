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

	// Parse the argument N from the command line
	nStr := os.Args[1]
	n, err := strconv.Atoi(nStr)
	if err != nil {
		// If input is not a valid integer, exit silently or handle error
		return
	}

	// If N < 1, print nothing as per requirements
	if n < 1 {
		return
	}

	// Initialize the first two numbers of the Fibonacci sequence
	a := 1
	b := 1

	// Print the first number
	fmt.Println(a)

	// If N is 1, we've already printed the first 1. 
	// We need to handle the second 1 in the sequence 1, 1, 2...
	if n >= 1 {
		// The sequence starts 1, 1...
		// We track the current number and the next number.
		
		// Special case: The sequence starts with two 1s.
		// If N is at least 1, we print the second 1.
		fmt.Println(b)
	}

	// Calculate subsequent numbers
	for {
		next := a + b
		// Stop if the next Fibonacci number exceeds N
		if next > n {
			break
		}
		fmt.Println(next)
		
		// Update values for the next iteration
		a = b
		b = next
	}
}