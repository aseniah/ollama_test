package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: <program> <integer>")
		os.Exit(1)
	}

	n, err := ParseInt(os.Args[1], 0, 2e9) // Handle up to large integers safely
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	if n < 1 {
		return
	}

	// Start with the first two Fibonacci numbers: F0=1, F1=1
	a := int64(1) // Previous Fib number
	b := int64(1) // Current Fib number (starts at 1)

	// Handle N >= 1 separately to ensure we print at least one '1' if needed
	if n >= 1 {
		fmt.Println(a)
		a, b = b, a+b
	}

	// Loop through the rest of the Fibonacci sequence
	for b <= int64(n) && b > 0 {
		fmt.Println(b)
		a, b = b, a+b
	}
}

// ParseInt parses the input string to an int64 and validates bounds.
func ParseInt(s string, minVal, maxVal int64) (int64, error) {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil || val < minVal || val > maxVal {
		return 0, fmt.Errorf("invalid integer: %s", s)
	}
	return val, nil
}

import "strconv"