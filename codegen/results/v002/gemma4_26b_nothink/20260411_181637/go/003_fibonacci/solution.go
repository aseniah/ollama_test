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

	// Parse the command line argument to an integer
	nStr := os.Args[1]
	n, err := strconv.Atoi(nStr)
	if err != nil {
		// If the argument is not a valid integer, exit silently or handle error
		return
	}

	// If N < 1, print nothing as per instructions
	if n < 1 {
		return
	}

	// Initialize the first two numbers of the Fibonacci sequence
	a := 1
	b := 1

	// Print the first number
	fmt.Println(a)

	// If N is at least 1 and the first number is 1, 
	// we check if we should print the second 1.
	// The sequence starts 1, 1, 2, 3...
	// If n is 1, we should print 1 and 1.
	if n >= 1 {
		// We check the second number in the sequence
		// This loop handles the sequence generation
		for b <= n {
			fmt.Println(b)
			
			// Calculate next Fibonacci number
			next := a + b
			a = b
			b = next
		}
	}
}