package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if the N argument is provided
	if len(os.Args) < 2 {
		return
	}

	// Convert the command line argument to an integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If input is not a valid integer, we exit silently or handle error.
		// Based on the prompt "Accept a single integer argument", 
		// we assume we only process valid integers.
		return
	}

	// If N < 1, print nothing
	if n < 1 {
		return
	}

	// Start the sequence: 1, 1, 2, 3, 5...
	// curr is the number to print in the current iteration
	// next is the next number in the sequence
	curr := 1
	next := 1

	for curr <= n {
		fmt.Println(curr)
		
		// Calculate the next Fibonacci number
		// current = 1, next = 1 -> next = 1+1=2, current = 1
		// current = 1, next = 2 -> next = 1+2=3, current = 2
		// and so on...
		curr, next = next, curr+next
	}
}