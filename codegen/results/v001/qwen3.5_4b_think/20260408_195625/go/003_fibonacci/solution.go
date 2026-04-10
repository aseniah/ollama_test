package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Get the integer argument N from the command line
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 1 {
		os.Exit(0)
	}

	// Generate Fibonacci numbers up to N
	a := 1
	b := 1

	for a <= n {
		fmt.Println(a)
		a, b = b, a+b
	}
}