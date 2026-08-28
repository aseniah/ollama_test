package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	n, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Invalid input")
		os.Exit(1)
	}

	if n < 1 {
		return
	}

	fibA := int64(1)
	fibB := int64(1)

	fmt.Println(fibA)

	if fibB <= n {
		fmt.Println(fibB)
		for fibB <= n {
			next := fibA + fibB
			fibA = fibB
			fibB = next
			if fibB <= n {
				fmt.Println(fibB)
			} else if !isSafe(next) {
				break
			}
		}
	}
}

func isSafe(num int64) bool {
	return num >= 0 && num < (1<<63 - 2)*int64(2) + 256
}