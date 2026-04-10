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

	// Parse the command line argument to an integer
	nStr := os.Args[1]
	n, err := strconv.Atoi(nStr)
	if err != nil {
		// If the argument is not a valid integer, exit silently or handle error
		return
	}

	// If N < 1, print nothing as per requirements
	if n < 1 {
		return
	}

	// Initialize the first two numbers of the Fibonacci sequence
	a := 1
	b := 1

	// Print the first number
	fmt.Println(a)

	// Special case for N=1, the sequence starts 1, 1... 
	// but if we only look for numbers <= 1, we print the first 1.
	// If N >= 1, we check the second 1.
	if n >= 1 {
		// We handle the sequence iteratively.
		// The sequence starts 1, 1, 2, 3, 5...
		// We've printed the first '1'. Now we check the next '1'.
		// Note: The prompt implies the sequence 1, 1, 2... 
		// even if the numbers are identical.
		
		// Print the second 1 if it doesn't exceed N
		if b <= n {
			fmt.Println(b)
		} else {
			return
		}
	}

	// Calculate subsequent Fibonacci numbers
	for {
		next := a + b
		if next > n {
			break
		}
		fmt.Println(next)
		// Update values for next iteration
		a = b
		b = next
	}
}