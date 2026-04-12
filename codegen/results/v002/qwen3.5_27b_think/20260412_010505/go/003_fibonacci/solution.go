package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if we have at least one argument
	if len(os.Args) < 2 {
		return
	}

	// Parse the integer argument
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	// If N < 1, print nothing
	if n < 1 {
		return
	}

	// Start the Fibonacci sequence with 1, 1
	a, b := 1, 1

	// Print the first Fibonacci number if it doesn't exceed N
	if a <= n {
		fmt.Println(a)
	}

	// Print the second Fibonacci number if it doesn't exceed N
	if b <= n {
		fmt.Println(b)
	}

	// Generate and print subsequent Fibonacci numbers
	for {
		next := a + b
		if next > n {
			break
		}
		fmt.Println(next)
		a, b = b, next
	}
}