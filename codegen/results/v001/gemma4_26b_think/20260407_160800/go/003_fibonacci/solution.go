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

	// Parse the argument from string to integer
	nStr := os.Args[1]
	n, err := strconv.Atoi(nStr)
	if err != nil {
		// If the argument is not a valid integer, we exit silently 
		// or handle as needed. Per instructions, we just need to process N.
		return
	}

	// If N < 1, print nothing.
	if n < 1 {
		return
	}

	// Initialize the first two numbers of the Fibonacci sequence
	// We use a loop that calculates the next term based on the current and next
	curr := 1
	next := 1

	for curr <= n {
		// Print the current Fibonacci number
		fmt.Println(curr)

		// Calculate the next number in the sequence
		// temp = curr + next, then shift current to the old next, and next to temp
		temp := curr + next
		curr = next
		next = temp
	}
}