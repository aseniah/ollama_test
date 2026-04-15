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

	// Print first 1 if it's <= n
	if a <= n {
		fmt.Println(a)
	}

	// If n >= 1, we also need to handle the second 1 if it's distinct in the sequence
	// The problem asks for: 1, 1, 2, 3, 5...
	// So if n >= 1, we print the first 1.
	// Then we loop while b <= n, printing b each time.
	// If n >= 1, the second 1 (b) is also <= n.

	for b <= n {
		fmt.Println(b)
		a, b = b, a+b
	}
}