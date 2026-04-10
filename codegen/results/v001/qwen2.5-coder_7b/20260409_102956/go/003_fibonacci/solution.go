package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run fib.go <integer>")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 1 {
		fmt.Println("Please provide a positive integer.")
		return
	}

	var a, b int = 0, 1
	for b <= n {
		a, b = b, a+b
		fmt.Println(a)
	}
}