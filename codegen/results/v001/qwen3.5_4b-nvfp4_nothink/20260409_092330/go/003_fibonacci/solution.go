package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run <file>.go <N>")
		os.Exit(1)
	}

	n, err := fmt.Sscanf(os.Args[1], "%d", nil)
	if n == 0 && err == nil {
		n, _ = fmt.Sscanf(os.Args[1], "%d", nil)
	} else if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid integer argument")
		os.Exit(1)
	}

	var a, b int64 = 0, 1
	for b <= n {
		fmt.Println(b)
		a, b = b, a+b
	}
}