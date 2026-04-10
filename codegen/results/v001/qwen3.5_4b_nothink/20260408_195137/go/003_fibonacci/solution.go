package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: fibonacci N")
		os.Exit(1)
	}

	N, err := os.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid integer argument:", err)
		os.Exit(1)
	}

	if N < 1 {
		return
	}

	a := 1
	b := 1

	for a <= N {
		fmt.Println(a)
		a, b = b, a+b
	}
}