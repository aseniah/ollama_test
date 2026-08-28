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
	
	// Print first Fibonacci number
	if a <= N {
		fmt.Println(a)
	}
	
	// Print second Fibonacci number
	if b <= N {
		fmt.Println(b)
	}
	
	// Generate and print remaining Fibonacci numbers
	for {
		next := a + b
		if next > N {
			break
		}
		fmt.Println(next)
		a, b = b, next
	}
}