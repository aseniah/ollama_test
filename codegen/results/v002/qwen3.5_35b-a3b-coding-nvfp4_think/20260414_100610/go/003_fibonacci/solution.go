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

	a, b := 1, 1

	// Print first Fibonacci number
	fmt.Println(a)

	// Print remaining Fibonacci numbers up to N
	for b <= n {
		fmt.Println(b)
		a, b = b, a+b
	}
}