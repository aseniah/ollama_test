package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		// No argument provided, print nothing
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// Invalid integer, print nothing
		return
	}

	if n < 1 {
		// N < 1, print nothing
		return
	}

	// Generate Fibonacci numbers up to N
	// Sequence: 1, 1, 2, 3, 5, ...
	a := 1
	b := 1

	for a <= n {
		fmt.Println(a)
		a, b = b, a+b
	}
}