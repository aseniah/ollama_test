package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if the command line argument is provided
	if len(os.Args) < 2 {
		return
	}

	// Parse the input argument N to an integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	// If N < 1, print nothing as per requirements
	if n < 1 {
		return
	}

	// Starting Fibonacci numbers: 1, 1, 2, 3, 5...
	a := 1
	b := 1

	// Print the first Fibonacci number
	fmt.Println(a)

	// If N is 1, we only print the first '1' (or both if we follow the 1, 1 sequence)
	// The prompt asks for the sequence starting 1, 1, 2, 3, 5...
	// We continue while the current Fibonacci number does not exceed N.
	for b <= n {
		fmt.Println(b)
		
		// Calculate next Fibonacci number
		next := a + b
		a = b
		b = next
	}
}