package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if the required argument N is provided
	if len(os.Args) < 2 {
		return
	}

	// Parse the command line argument to an integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If the argument is not a valid integer, we exit
		return
	}

	// If N < 1, the requirement states we should print nothing
	if n < 1 {
		return
	}

	// We define the start of the Fibonacci sequence.
	// To handle the duplicate '1' at the start (1, 1, 2, 3...),
	// we initialize curr as the first 1 and next as the second 1.
	curr := 1
	next := 1

	// We iterate through the sequence as long as the current 
	// Fibonacci number does not exceed N.
	for curr <= n {
		// Print the current Fibonacci number on its own line
		fmt.Println(curr)

		// Calculate the next number in the sequence
		temp := curr + next
		
		// Move to the next step in the sequence.
		// To ensure we print the second '1' when N >= 1,
		// we shift the current value to the 'next' position
		// and update 'next' to the new sum.
		// Trace N=1: 
		// Loop 1: Prints 1. temp = 1+1=2. curr = 1. next = 2.
		// Loop 2: Prints 1. temp = 1+2=3. curr = 2. next = 3.
		// Loop 3: 2 <= 1 is false. End.
		curr = next
		next = temp
	}
}