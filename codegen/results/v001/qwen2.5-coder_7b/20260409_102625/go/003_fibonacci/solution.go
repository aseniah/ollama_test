package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run fibonacci.go <integer>")
		return
	}

	nStr := os.Args[1]
	n, err := strconv.Atoi(nStr)
	if err != nil {
		fmt.Printf("Invalid argument: %s\n", nStr)
		return
	}

	if n < 1 {
		return
	}

	var a, b = 0, 1
	for a <= n {
		fmt.Println(a)
		a, b = b, a+b
	}
}