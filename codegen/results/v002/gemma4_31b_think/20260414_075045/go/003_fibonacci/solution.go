package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Ensure a command line argument is provided
	if len(os.Args) < 2 {
		return
	}

	// Parse the argument N as an integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If the argument is not a valid integer, we exit silently or could print an error
		return
	}

	// If N < 1, print nothing as per the requirements
	if n < 1 {
		return
	}

	// Fibonacci sequence starts: 1, 1, 2, 3, 5...
	// We initialize current as the first number and next as the second number.
	current := 1
	next := 1

	for current <= n {
		// Print the current Fibonacci number
		fmt.Println(current)

		// Calculate the next number in the sequence
		// current becomes the old 'next', and next becomes 'current + next'
		temp := current + next
		current = next
		next = temp
	}
}