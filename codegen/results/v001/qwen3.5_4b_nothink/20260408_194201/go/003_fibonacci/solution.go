package main

import (
	"fmt"
	"os"
)

func main() {
	// Get N from command line arguments
	n, err := parseInt(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	// If N < 1, print nothing
	if n < 1 {
		return
	}

	// Generate Fibonacci numbers up to N
	a, b := 1, 1
	count := 0
	for {
		if a > n {
			break
		}
		fmt.Println(a)
		count++
		if b > n {
			break
		}
		fmt.Println(b)
		count++
		// Check if b exceeds n to prevent overflow in next iteration if N is large
		if a+n > n {
			break
		}
		a, b = b, a+b
	}
}

func parseInt(s string) (int, error) {
	var n int
	_, err := fmt.Sscanf(s, "%d", &n)
	if err != nil {
		return 0, err
	}
	return n, nil
}