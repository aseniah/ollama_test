package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Ensure exactly one argument is provided
	if len(os.Args) < 2 {
		return
	}

	// Parse the command line argument
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	// If N < 1, print nothing as per requirements
	if n < 1 {
		return
	}

	// Generate and print Fibonacci numbers up to N
	a, b := 1, 1
	for a <= n {
		fmt.Println(a)
		a, b = b, a+b
	}
}