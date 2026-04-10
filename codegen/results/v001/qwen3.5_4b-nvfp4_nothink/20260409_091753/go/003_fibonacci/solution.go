package main

import (
	"fmt"
	"os"
)

func main() {
	// Check argument count
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	// Parse N
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 1 {
		// If N is not a valid integer or less than 1, print nothing as per requirement
		os.Exit(0)
	}

	// Handle case where N >= 1
	fmt.Println(1)
	first := 1
	fibSecond := 1 // Represents the current Fibonacci number 'after' the first one

	for fibSecond <= n && fibSecond != 0 {
		// Generate next Fibonacci numbers if possible
		if fibSecond == 0 || fibSecond < n {
			fibNext := first + fibSecond
			fmt.Println(fibNext)
			first = fibSecond
			fibSecond = fibNext
		} else {
			break
		}
	}
}