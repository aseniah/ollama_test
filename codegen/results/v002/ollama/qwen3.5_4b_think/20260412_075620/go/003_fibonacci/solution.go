package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Get N from command line argument
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Error: need exactly one integer argument")
		os.Exit(1)
	}

	n, err := strconv.ParseInt(os.Args[1], 10, 0)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: invalid integer argument")
		os.Exit(1)
	}

	if n < 1 {
		return
	}

	// Generate Fibonacci numbers starting with 1, 1, 2, 3, 5...
	a, b := 1, 1
	for a <= n {
		fmt.Println(a)
		next := a + b
		a = b
		b = next
	}
}