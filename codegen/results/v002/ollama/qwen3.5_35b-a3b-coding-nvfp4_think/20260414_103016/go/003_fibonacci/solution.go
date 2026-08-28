package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if we have the command line argument
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

	// Generate Fibonacci sequence starting with 1, 1, 2, 3, 5, ...
	a, b := 1, 1
	for a <= n {
		fmt.Println(a)
		next := a + b
		a = b
		b = next
	}
}