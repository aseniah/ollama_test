package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Ensure exactly one command-line argument is provided
	if len(os.Args) != 2 {
		return
	}

	// Parse the integer argument N
	n, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		return
	}

	// If N < 1, print nothing
	if n < 1 {
		return
	}

	// Initialize the first two Fibonacci numbers
	a := int64(1)
	b := int64(1)

	// Print the first number if it's <= N
	if a <= n {
		fmt.Println(a)
	}

	// If N is 1, we stop here because the next number is also 1,
	// but the sequence logic below handles the second 1 correctly
	// by checking b <= n. However, to strictly follow the sequence
	// 1, 1, 2... we need to ensure the second 1 is printed if N >= 1.
	if b <= n {
		fmt.Println(b)
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