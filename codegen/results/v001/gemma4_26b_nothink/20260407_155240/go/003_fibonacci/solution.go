package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if an argument was provided
	if len(os.Args) < 2 {
		return
	}

	// Parse the input argument N
	nStr := os.Args[1]
	n, err := strconv.Atoi(nStr)
	if err != nil {
		// If the input is not a valid integer, we exit silently 
		// (or you could handle error specifically)
		return
	}

	// If N < 1, the requirement says print nothing.
	if n < 1 {
		return
	}

	// Initialize the first two numbers of the Fibonacci sequence
	a := 1
	b := 1

	// Print the first number
	fmt.Println(a)

	// If the second number is also <= N, we enter the loop
	// Note: The sequence starts 1, 1, 2...
	// We handle the second '1' separately to manage the loop logic correctly
	if b <= n {
		fmt.Println(b)
	} else {
		// This case handles if N was 1, but the second 1 is already handled.
		// However, with N >= 1, the loop below handles subsequent numbers.
		return
	}

	// Calculate and print subsequent Fibonacci numbers
	for {
		next := a + b
		if next > n {
			break
		}
		fmt.Println(next)
		// Update values for the next iteration
		a = b
		b = next
	}
}