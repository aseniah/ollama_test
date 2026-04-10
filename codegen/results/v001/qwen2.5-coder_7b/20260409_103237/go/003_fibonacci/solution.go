package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run fibonacci.go <N>")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 1 {
		fmt.Println("Please provide a positive integer.")
		return
	}

	a, b := 1, 1
	for a <= n {
		fmt.Println(a)
		a, b = b, a+b
	}
}