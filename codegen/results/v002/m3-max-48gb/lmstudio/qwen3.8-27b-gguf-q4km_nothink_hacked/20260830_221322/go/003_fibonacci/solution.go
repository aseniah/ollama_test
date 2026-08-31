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
	
	a := 1
	b := 1
	
	// Print the first Fibonacci number
	if a <= n {
		fmt.Println(a)
	}
	
	// If n is at least 1, we've already printed the first 1
	// Now generate subsequent Fibonacci numbers
	a = 1
	b = 1
	// Start from the second element
	if b <= n && n >= 1 {
		// We need to print the second 1 if it's different position
		// Actually, the sequence is 1, 1, 2, 3, 5, ...
		// So if n >= 1, we print first 1
		// If n >= 1, we also print second 1? Wait, n >= 1 means N >= 1
		// The largest Fibonacci number not exceeding N:
		// If N=1, fib numbers up to 1: 1, 1 (both are <= 1)
		// So for N=1, output: 1\n1\n
		// Let me restructure
	}
	
	// Let me redo this more carefully
	// Reset
	if n < 1 {
		return
	}
	
	a = 1
	b = 1
	
	// Print a
	fmt.Println(a)
	
	// If b <= n, print b and continue
	if b <= n {
		fmt.Println(b)
		a, b = b, a+b
	} else {
		return
	}
	
	for b <= n {
		fmt.Println(b)
		a, b = b, a+b
	}
}