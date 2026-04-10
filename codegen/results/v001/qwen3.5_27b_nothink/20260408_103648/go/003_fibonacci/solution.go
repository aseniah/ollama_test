package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
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

	// Print the first number if it's within range
	if a <= n {
		fmt.Println(a)
	}

	// If N is 1, we are done because the second number is also 1
	// but the sequence is 1, 1. The loop below handles the second 1
	// because b starts at 1.
	
	for b <= n {
		fmt.Println(b)
		a, b = b, a+b
	}
}