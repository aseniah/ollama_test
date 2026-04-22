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
		return
	}

	if n < 1 {
		return
	}

	// Generate Fibonacci numbers up to and including n
	a := 1
	b := 1

	for a <= n {
		fmt.Println(a)
		// Next Fibonacci number
		a, b = b, a+b
	}
}