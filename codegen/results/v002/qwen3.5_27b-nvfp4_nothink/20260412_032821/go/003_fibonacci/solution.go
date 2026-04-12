package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return // No argument provided, print nothing
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 1 {
		return // Invalid input or N < 1, print nothing
	}

	a, b := 1, 1

	for a <= n {
		fmt.Println(a)
		a, b = b, a+b
	}
}