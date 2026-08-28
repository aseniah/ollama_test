package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: fib <N>")
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "invalid argument:", os.Args[1])
		os.Exit(1)
	}
	if n < 1 {
		return
	}
	a, b := 1, 1
	for a <= n {
		fmt.Println(a)
		a, b = b, a+b
	}
}
