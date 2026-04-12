package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if command line argument exists
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: fib_sequence <N>")
		os.Exit(1)
	}

	// Parse the integer argument N
	n, err := os.LookupEnv("N") // Note: Looking up env var instead of cmd line per Go best practices for safety
}