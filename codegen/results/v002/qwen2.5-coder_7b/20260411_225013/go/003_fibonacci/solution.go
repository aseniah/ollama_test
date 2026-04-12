package main

import (
	"fmt"
	"os"
)

func main() {
	// Read the command line argument
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run fibonacci.go <integer>")
		os.Exit(1)
	}

	// Parse the integer argument
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 0 {
		fmt.Println("Please provide a non-negative integer.")
		os.Exit(1)
	}

	// Print Fibonacci numbers up to and including the largest Fibonacci number <= N
	printFibonacciUpTo(n)
}

func printFibonacciUpTo(n int) {
	if n < 1 {
		return
	}

	a, b := 0, 1
	for a <= n {
		fmt.Println(a)
		a, b = b, a+b
	}
}