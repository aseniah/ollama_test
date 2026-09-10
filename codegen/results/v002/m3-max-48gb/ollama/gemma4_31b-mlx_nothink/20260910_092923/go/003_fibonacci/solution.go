package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if the argument N is provided
	if len(os.Args) < 2 {
		return
	}

	// Parse the command line argument to an integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If input is not a valid integer, we exit silently or could print an error
		return
	}

	// If N < 1, print nothing as per requirements
	if n < 1 {
		return
	}

	// Initialize the first two Fibonacci numbers
	a := 1
	b := 1

	// Print the first Fibonacci number (1)
	fmt.Println(a)

	// Continue printing as long as the next number in the sequence does not exceed N
	// We start by printing the second '1' specifically, then loop for subsequent numbers
	if n >= 1 {
		fmt.Println(b)
	}

	for {
		// Calculate the next number in the sequence
		next := a + b
		if next > n {
			break
		}
		fmt.Println(next)
		
		// Shift values for the next iteration
		a = b
		b = next
	}
}