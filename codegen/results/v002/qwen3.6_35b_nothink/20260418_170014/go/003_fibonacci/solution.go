package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	if n < 1 {
		return
	}

	// Generate Fibonacci numbers up to n
	// Start with 1, 1, 2, 3, 5, ...
	a := 1
	b := 1

	// Print the first Fibonacci number (1)
	fmt.Println(a)

	// If n >= 1, we already printed the first 1.
	// Now continue generating until we exceed n.
	// But wait: if n is exactly 1, we printed one 1.
	// The sequence is 1, 1, 2, 3, 5, ...
	// So we need to print the second 1 if n >= 1.
	if b <= n {
		fmt.Println(b)
	}

	// Generate subsequent Fibonacci numbers
	for {
		c := a + b
		if c > n {
			break
		}
		fmt.Println(c)
		a = b
		b = c
	}
}