package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if a single integer argument is provided
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <integer N>\n", os.Args[0])
		os.Exit(1)
	}

	// Parse the integer argument N
	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid integer argument: %s\n", os.Args[1])
		os.Exit(1)
	}

	// If N < 1, print nothing
	if N < 1 {
		return
	}

	// Generate Fibonacci numbers up to N
	a, b := 1, 1
	for {
		if a > N {
			break
		}
		fmt.Println(a)
		next := a + b
		a = b
		b = next
	}
}