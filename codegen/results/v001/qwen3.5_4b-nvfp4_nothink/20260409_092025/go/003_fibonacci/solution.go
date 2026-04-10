package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if an argument was provided
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <integer>\n", os.Args[0])
		os.Exit(1)
	}

	var n int
	_, err := fmt.Sscanf(os.Args[1], "%d", &n)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing argument: %v\n", err)
		os.Exit(1)
	}

	// If N < 1, print nothing
	if n < 1 {
		return
	}

	// Start the Fibonacci sequence: F0 = 1, F1 = 1, F2 = 2, ...
	var a int = 1 // First term (1)
	var b int = 1 // Second term (1)
	fmt.Println(a) // Print first term

	if n >= b {
		fmt.Println(b) // If current second term <= n, print it
		a = b // Move to next iteration
		b += a // Calculate next term: F2 = F1 + F0 = 1 + 1 = 2
	}

	for b <= n {
		fmt.Println(b)
		// Prepare next terms: a = current 'b', new b = a + old b
		a = b
		b += a
	}
}