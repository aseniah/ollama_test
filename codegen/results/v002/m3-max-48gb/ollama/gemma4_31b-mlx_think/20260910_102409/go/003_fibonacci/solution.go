package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Ensure that an argument was provided
	if len(os.Args) < 2 {
		return
	}

	// Parse the command line argument into an integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If the argument is not a valid integer, we exit silently
		return
	}

	// Requirement: If N < 1, print nothing.
	if n < 1 {
		return
	}

	// Fibonacci sequence starting: 1, 1, 2, 3, 5...
	// We initialize curr to 1 and next to 1 to produce the starting 1, 1.
	curr := 1
	next := 1

	// Continue printing as long as the current Fibonacci number does not exceed N
	for curr <= n {
		fmt.Println(curr)
		
		// Calculate the next number in the sequence
		// Go allows tuple assignment: curr becomes the previous next, 
		// and next becomes the sum of the two.
		curr, next = next, curr+next
	}
}