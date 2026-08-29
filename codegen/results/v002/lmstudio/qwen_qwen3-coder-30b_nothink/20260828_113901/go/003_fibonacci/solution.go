package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if command line argument is provided
	if len(os.Args) < 2 {
		return
	}

	// Parse the command line argument
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	// If N < 1, print nothing
	if n < 1 {
		return
	}

	// Generate and print Fibonacci numbers
	a, b := 1, 1
	
	// Print first Fibonacci number
	if a <= n {
		fmt.Println(a)
	}
	
	// Print second Fibonacci number
	if b <= n {
		fmt.Println(b)
	}
	
	// Generate and print remaining Fibonacci numbers
	for {
		next := a + b
		if next > n {
			break
		}
		fmt.Println(next)
		a, b = b, next
	}
}