package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: program <N>")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: invalid integer argument")
		os.Exit(1)
	}

	if n < 1 {
		return
	}

	a := 0
	b := 1
	for b <= n && b > 0 {
		fmt.Println(b)
		a, b = b, a+b
	}
}