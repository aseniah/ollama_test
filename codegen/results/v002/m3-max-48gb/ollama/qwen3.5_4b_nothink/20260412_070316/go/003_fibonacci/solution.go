package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <integer>\n", os.Args[0])
		os.Exit(1)
	}

	n, err := parseInt(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing argument: %v\n", err)
		os.Exit(1)
	}

	if n < 1 {
		return
	}

	// Start Fibonacci sequence: 1, 1, 2, 3, 5, ...
	var a, b int = 1, 1

	for {
		if a > n {
			break
		}
		fmt.Println(a)
		a, b = b, a+b
	}
}

func parseInt(s string) (int, error) {
	n, err := fmt.Sscanf(s, "%d", &n)
	if err != nil || n != 1 {
		return 0, fmt.Errorf("invalid integer: %s", s)
	}
	return n, nil
}