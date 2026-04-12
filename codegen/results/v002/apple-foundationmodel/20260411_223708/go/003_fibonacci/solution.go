package main

import (
	"fmt"
	"os"
)

func main() {
	// Read the command line argument
	n, _ := os.ParseInt(os.Args[1], 10, 32)

	// Check if the input is valid
	if n < 1 {
		fmt.Println("N must be at least 1.")
		return
	}

	// Initialize the first two Fibonacci numbers
	fib1, fib2 := 1, 1

	// Print Fibonacci numbers until the next number would exceed N
	for fib1 <= n {
		fmt.Println(fib1)
		fibNext := fib1 + fib2
		fib1, fib2 = fib2, fibNext
	}
}