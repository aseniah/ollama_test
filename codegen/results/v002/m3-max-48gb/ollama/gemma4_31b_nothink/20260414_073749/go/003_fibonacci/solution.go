package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if a command line argument was provided
	if len(os.Args) < 2 {
		return
	}

	// Parse the argument N as an integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If input is not a valid integer, we exit silently or could print an error
		return
	}

	// If N < 1, print nothing as per requirements
	if n < 1 {
		return
	}

	// Initial Fibonacci sequence values
	a, b := 1, 1

	// Print the first Fibonacci number (1)
	fmt.Println(a)

	// Handle the case where N is 1 specifically to avoid printing 
	// the second '1' if it exceeds N (though 1 does not exceed 1)
	if n < 2 && n >= 1 {
		// According to the rule: print the largest Fib that does not exceed N.
		// If N=1, the sequence 1, 1 are both <= 1.
		fmt.Println(b)
		return
	}

	// Print the second 1
	fmt.Println(b)

	// Generate subsequent Fibonacci numbers
	for {
		next := a + b
		if next > n {
			break
		}
		fmt.Println(next)
		a = b
		b = next
	}
}