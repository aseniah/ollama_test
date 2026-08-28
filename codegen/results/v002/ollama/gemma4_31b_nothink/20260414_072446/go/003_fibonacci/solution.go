package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if a command line argument is provided
	if len(os.Args) < 2 {
		return
	}

	// Parse the input argument N as an integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If input is not a valid integer, we stop execution
		return
	}

	// If N < 1, print nothing as per requirements
	if n < 1 {
		return
	}

	// Initialize the first two numbers of the sequence
	a := 1
	b := 1

	// Print the first number
	fmt.Println(a)

	// If N is 1, we only print the first '1' once if we follow 1, 1, 2... 
	// but the prompt specifies 1, 1, 2, 3, 5... 
	// So we handle the sequence generation loop.
	for b <= n {
		fmt.Println(b)
		
		// Calculate next Fibonacci number
		next := a + b
		a = b
		b = next
	}
}