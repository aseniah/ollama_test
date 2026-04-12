package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if argument is provided
	if len(os.Args) < 2 {
		return
	}

	// Parse the argument
	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	// If N < 1, print nothing
	if N < 1 {
		return
	}

	// Generate and print Fibonacci numbers
	a, b := 1, 1
	for a <= N {
		fmt.Println(a)
		a, b = b, a+b
	}
}