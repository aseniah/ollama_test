package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if the command line argument is provided
	if len(os.Args) < 2 {
		return
	}

	// Parse the integer argument N
	n, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		return
	}

	// If N < 1, print nothing
	if n < 1 {
		return
	}

	// Initialize the first two Fibonacci numbers in the sequence 1, 1, 2, ...
	a := int64(1)
	b := int64(1)

	// Print Fibonacci numbers while they do not exceed N
	for a <= n {
		fmt.Println(a)
		sum := a + b
		a = b
		b = sum
	}
}