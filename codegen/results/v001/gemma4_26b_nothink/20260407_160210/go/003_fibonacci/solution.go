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

	// Parse the command line argument N
	nStr := os.Args[1]
	n, err := strconv.Atoi(nStr)
	if err != nil {
		// If N is not a valid integer, we exit silently or handle error
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

	// If N is at least 1, we check if we can print the second 1
	// The sequence starts 1, 1, 2, 3...
	if n >= 1 {
		// We handle the second '1' specifically to maintain the sequence structure
		// but we must ensure we don't print a second 1 if the logic requires 
		// only numbers <= N. Since N >= 1, the second 1 is always valid.
		// However, to avoid infinite loops or logic errors, we check b <= n.
		if b <= n {
			// We only print it if it's not the same instance as the first 'a'
			// But for the standard sequence 1, 1, 2..., we just print it.
			// To avoid duplicate logic, we can use a loop.
		}
	}

	// Alternative approach: use a loop that tracks the current and next value
	curr := 1
	next := 1

	// We already handled the first '1' via logic above, 
	// let's rewrite the loop for clarity and correctness.
	
	// Resetting logic for a clean loop:
	first := 1
	second := 1

	if n >= 1 {
		fmt.Println(first)
	}
	if n >= 1 {
		// The requirement says 1, 1, 2, 3... 
		// If N=1, we should print 1, 1.
		// Let's use a loop that maintains the sequence.
	}
}