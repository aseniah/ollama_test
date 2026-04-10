package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if command line argument exists
	if len(os.Args) < 2 {
		fmt.Println("Usage: fibonacci <N>")
		os.Exit(1)
	}

	// Parse the integer argument
	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: argument must be an integer")
		os.Exit(1)
	}

	// If N < 1, print nothing
	if N < 1 {
		return
	}

	// Generate Fibonacci numbers up to N
	// Sequence starts with 1, 1, 2, 3, 5, ...
	if N >= 1 {
		fmt.Println(1)
	}

	a := 1
	b := 1

	// Generate and print subsequent Fibonacci numbers
	for b <= N {
		fmt.Println(b)
		newNum := a + b
		a = b
		b = newNum
	}
}