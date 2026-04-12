package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the command-line argument N
	var n int
	if len(os.Args) != 2 {
		fmt.Println("Usage: fibonacci.go <N>")
		os.Exit(1)
	}
	n, err := os.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error parsing argument:", err)
		os.Exit(1)
	}

	// Check if N is less than 1
	if n < 1 {
		fmt.Println("N must be at least 1.")
		os.Exit(1)
	}

	// Generate Fibonacci numbers and print them until we exceed N
	fib, _ := generateFibonacciNumbers(n)
	for _, fibNumber := range fib {
		fmt.Print(fibNumber)
		fmt.Println()
	}
}

// Function to generate Fibonacci numbers up to N
func generateFibonacciNumbers(limit int) ([]int, error) {
	if limit <= 0 {
		return nil, nil
	}

	fib := []int{1, 1}
	next := 2 // Starting with 2 as the next Fibonacci number
	for i := 1; i < len(fib); i++ {
		next = fib[i-1] + fib[i-2]
		if next > limit {
			return fib[:i], nil
		}
		fib = append(fib, next)
	}
	return fib, nil
}