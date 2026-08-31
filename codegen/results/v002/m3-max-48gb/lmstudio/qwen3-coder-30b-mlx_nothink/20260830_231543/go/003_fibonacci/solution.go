package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	if n < 1 {
		return
	}

	// Start the Fibonacci sequence
	a, b := 1, 1
	
	// Print the first number if it's within range
	if a <= n {
		fmt.Println(a)
	}
	
	// Print the second number if it's within range
	if b <= n {
		fmt.Println(b)
	}
	
	// Generate and print subsequent Fibonacci numbers
	for {
		next := a + b
		if next > n {
			break
		}
		fmt.Println(next)
		a, b = b, next
	}
}