package main

import (
	"fmt"
	"os"
)

func main() {
	// Read the integer argument from command line
	if len(os.Args) != 2 {
		fmt.Println("Usage: go fibonacci.go <N>")
		return
	}

	n := os.Args[1]

	// Check if N is a valid integer
	if, ok := n.(int); !ok {
		fmt.Println("Invalid input. Please provide a valid integer.")
		return
	}

	if n < 1 {
		fmt.Println("N must be at least 1.")
		return
	}

	// Initialize the first two Fibonacci numbers
	a, b := 1, 1
	for {
		// Check if the next Fibonacci number exceeds N
		c := a + b
		if c > n {
			break
		}
		// Print the Fibonacci number
		fmt.Println(c)
		// Move to the next pair
		a, b = b, c
	}
}