package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run program.go <integer>")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 0 {
		fmt.Println("Please provide a non-negative integer.")
		os.Exit(1)
	}

	if n == 0 {
		return
	}

	fib1, fib2 := 1, 1
	fmt.Println(fib1)

	for {
		nextFib := fib1 + fib2
		if nextFib > n {
			break
		}
		fmt.Println(nextFib)
		fib1, fib2 = fib2, nextFib
	}
}