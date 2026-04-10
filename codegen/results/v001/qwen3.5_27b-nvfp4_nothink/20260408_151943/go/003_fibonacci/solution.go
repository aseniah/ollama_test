package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run main.go <integer N>")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: argument must be a valid integer")
		os.Exit(1)
	}

	if n < 1 {
		return
	}

	a := 1
	b := 1

	for a <= n {
		fmt.Println(a)
		a, b = b, a+b
	}
}