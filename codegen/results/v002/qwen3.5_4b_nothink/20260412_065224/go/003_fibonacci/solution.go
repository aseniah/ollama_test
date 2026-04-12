package main

import (
	"fmt"
	"os"
)

func main() {
	// Get N from command line arguments
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	n, err := ParseInt(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	// If N < 1, print nothing
	if n < 1 {
		return
	}

	// Generate Fibonacci numbers starting 1, 1, 2, 3, 5, ...
	// a and b are consecutive Fibonacci numbers
	a := 1
	b := 1

	for a <= n {
		fmt.Println(a)
		next := a + b
		a = b
		b = next
	}
}

// ParseInt parses an integer from a string, handling optional sign and surrounding whitespace
func ParseInt(s string) (int, error) {
	// Trim whitespace
	s = trimSpace(s)
	// Remove optional leading '+'
	if len(s) > 0 && s[0] == '+' {
		s = s[1:]
	}
	if len(s) == 0 {
		return 0, fmt.Errorf("empty string")
	}
	// Check for leading '-' if present, but handle as part of string
	if len(s) > 0 && s[0] == '-' {
		s = s[1:]
	}
	if len(s) == 0 {
		return 0, fmt.Errorf("empty string after sign")
	}

	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(val), nil
}

// trimSpace returns a copy of s with leading and trailing spaces removed
func trimSpace(s string) string {
	start := 0
	end := len(s)

	// skip leading whitespace
	for start < end && (s[start] <= ' ') {
		start++
	}

	// skip trailing whitespace
	for end > start && (s[end-1] <= ' ') {
		end--
	}

	return s[start:end]
}

// Note: The above code does not import strconv which is needed for ParseInt
// Let's correct the imports and remove the custom parse logic for simplicity and correctness