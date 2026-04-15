package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if an argument was provided
	if len(os.Args) < 2 {
		return
	}

	// Parse the command line argument N
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// If the input is not a valid integer, we stop execution
		return
	}

	// If N < 1, print nothing as per requirements
	if n < 1 {
		return
	}

	// Initialize the first two Fibonacci numbers
	a := 1
	b := 1

	// Print the first number
	fmt.Println(a)

	// We handle the sequence iteratively.
	// Since the sequence starts 1, 1, 2... and we already printed the first 1,
	// we check if the second 1 is within range before entering the loop.
	if b <= n {
		fmt.Println(b)
	} else {
		// If n was 1, we already printed the first 1. 
		// However, based on the start 1, 1, 2... if N=1, both 1s should be printed.
		// But if N < 1, nothing is printed.
		// Since N >= 1 here, and the second number is 1, it will always print if N >= 1.
	}

	// Generate subsequent Fibonacci numbers
	for {
		next := a + b
		if next > n {
			break
		}
		fmt.Println(next)
		a = b
		b = next
	}
}