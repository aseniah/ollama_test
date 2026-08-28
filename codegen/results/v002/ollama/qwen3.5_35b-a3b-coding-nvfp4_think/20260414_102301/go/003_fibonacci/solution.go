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

	a, b := 1, 1

	// Print first Fibonacci number
	if a <= n {
		fmt.Println(a)
	}

	// Print second Fibonacci number
	if b <= n {
		fmt.Println(b)
	}

	// Generate and print remaining Fibonacci numbers
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