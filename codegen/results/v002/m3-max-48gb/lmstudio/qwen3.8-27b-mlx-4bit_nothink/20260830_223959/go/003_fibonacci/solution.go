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
		fmt.Println("Error: please provide a valid integer argument")
		os.Exit(1)
	}

	if n < 1 {
		return
	}

	// Print Fibonacci numbers up to and including the largest one that doesn't exceed N
	// Sequence starts: 1, 1, 2, 3, 5, 8, ...
	a, b := 1, 1
	for {
		if a > n {
			break
		}
		fmt.Println(a)
		a, b = b, a+b
	}
}