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

	// Parse the input argument N as an integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If the argument is not a valid integer, the program terminates silently
		return
	}

	// If N < 1, the requirement is to print nothing
	if n < 1 {
		return
	}

	// The Fibonacci sequence defined in the prompt starts: 1, 1, 2, 3, 5...
	// We initialize the first two numbers.
	a, b := 1, 1

	// Since n >= 1, the first Fibonacci number (1) always satisfies the condition.
	fmt.Println(a)
	
	// The second Fibonacci number (1) also always satisfies the condition if n >= 1.
	fmt.Println(b)

	// Calculate subsequent Fibonacci numbers using a loop.
	for {
		next := a + b
		
		// If the next number in the sequence exceeds N, we stop.
		if next > n {
			break
		}
		
		// Print the number and shift the variables forward.
		fmt.Println(next)
		a = b
		b = next
	}
}