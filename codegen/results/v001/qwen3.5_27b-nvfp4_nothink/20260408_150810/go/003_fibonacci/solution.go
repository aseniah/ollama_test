package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: provide a single integer argument")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 1 {
		// If N < 1 or invalid input, print nothing.
		return
	}

	// Start the sequence: 1, 1, 2, 3, 5, ...
	a, b := 1, 1

	for a <= n {
		fmt.Println(a)
		a, b = b, a+b
	}
}