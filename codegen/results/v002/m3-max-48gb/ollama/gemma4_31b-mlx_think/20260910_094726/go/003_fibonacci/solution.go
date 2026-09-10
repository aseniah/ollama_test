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

	// Parse the input N from the first command line argument
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If the input is not a valid integer, we simply exit as per common CLI patterns
		return
	}

	// If N < 1, print nothing as per requirements
	if n < 1 {
		return
	}

	// Fibonacci sequence starting with 1, 1, 2, 3, 5...
	// a is the current number to print, b is the next number in the sequence
	a, b := 1, 1

	for a <= n {
		fmt.Println(a)
		
		// Calculate next Fibonacci number
		// a becomes the previous b, b becomes the sum of a and b
		a, b = b, a+b
	}
}